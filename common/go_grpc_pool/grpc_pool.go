package go_grpc_pool

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"sync"
	"time"
)

var (
	// ErrClosed is the error when the client pool is closed
	ErrClosed = errors.New("grpc pool: client pool is closed")
	// ErrTimeout is the error when the client pool timed out
	ErrTimeout = errors.New("grpc pool: client pool timed out")
	// ErrAlreadyClosed is the error when the client conn was already closed
	ErrAlreadyClosed = errors.New("grpc pool: the connection was already closed")
	// ErrFullPool is the error when the pool is already full
	ErrFullPool = errors.New("grpc pool: closing a ClientConn into a full pool")
)

// Factory 链接函数的类型
type Factory func(addr string) (*grpc.ClientConn, error)

// FactoryWithContext  链接函数类型的实现
type FactoryWithContext func(ctx context.Context) (*grpc.ClientConn, error)

// ClientConn grpc Client 客户端
type ClientConn struct {
	*grpc.ClientConn
	pool          *Pool
	timeUsed      time.Time
	timeInitiated time.Time
	unhealthy     bool
}

// Pool 客户端连接池 chan
type Pool struct {
	clients         chan ClientConn
	factory         FactoryWithContext
	idleTimeout     time.Duration
	maxLifeDuration time.Duration
	mu              sync.RWMutex
}

// New 实例化连接池
func New(addr string, factory Factory, init, capacity int, idleTimeout time.Duration, maxLifeDuration ...time.Duration) (*Pool, error) {
	f := func(ctx context.Context) (*grpc.ClientConn, error) { return factory(addr) }
	return NewWithContext(context.Background(), f, init, capacity, idleTimeout, maxLifeDuration...)
}

// NewWithContext 实例化 新建客户端链接 添加到连接池
func NewWithContext(ctx context.Context, factory FactoryWithContext, init, capacity int, idleTimeout time.Duration, maxLifeDuration ...time.Duration) (*Pool, error) {
	if capacity <= 0 {
		capacity = 1
	}
	if init < 0 {
		init = 0
	}
	if init > capacity {
		init = capacity
	}
	p := &Pool{
		clients:     make(chan ClientConn, capacity),
		factory:     factory,
		idleTimeout: idleTimeout,
	}
	if len(maxLifeDuration) > 0 {
		p.maxLifeDuration = maxLifeDuration[0]
	}
	for i := 0; i < init; i++ {
		c, err := factory(ctx)
		if err != nil {
			return nil, err
		}
		p.clients <- ClientConn{
			ClientConn:    c,
			pool:          p,
			timeUsed:      time.Now(),
			timeInitiated: time.Now(),
		}
	}
	for i := 0; i < capacity-init; i++ {
		p.clients <- ClientConn{
			pool: p,
		}
	}
	return p, nil
}

// getClients 私有方法 获取连接池中的客户端管道
func (p *Pool) getClients() chan ClientConn {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.clients
}

// Close p.Close 关闭连接池
func (p *Pool) Close() {
	p.mu.Lock()
	clients := p.clients
	p.clients = nil
	p.mu.Unlock()
	if clients == nil {
		return
	}
	close(clients)
	for client := range clients {
		if client.ClientConn == nil {
			continue
		}
		client.ClientConn.Close()
	}
}

// IsClosed 是否关闭
func (p *Pool) IsClosed() bool {
	return p == nil || p.getClients() == nil
}

// Get 获取链接
func (p *Pool) Get(ctx context.Context) (*ClientConn, error) {
	clients := p.getClients()
	if clients == nil {
		return nil, ErrClosed
	}
	wrapper := ClientConn{
		pool: p,
	}
	select {
	case wrapper = <-clients:
	case <-ctx.Done():
		return nil, ErrTimeout
	}
	idleTimeout := p.idleTimeout
	if wrapper.ClientConn != nil && idleTimeout > 0 && wrapper.timeUsed.Add(idleTimeout).Before(time.Now()) {
		_ = wrapper.ClientConn.Close()
		wrapper.ClientConn = nil
	}
	var err error
	if wrapper.ClientConn == nil {
		wrapper.ClientConn, err = p.factory(ctx)
		if err != nil {
			clients <- ClientConn{pool: p}
		}
		wrapper.timeInitiated = time.Now()
	}
	return &wrapper, err
}

func (c *ClientConn) Unhealthy() {
	c.unhealthy = true
}

// Close ClientConn
func (c *ClientConn) Close() error {
	if c == nil {
		return nil
	}
	if c.ClientConn == nil {
		return ErrAlreadyClosed
	}
	if c.pool.IsClosed() {
		return ErrClosed
	}
	maxDuration := c.pool.maxLifeDuration
	if maxDuration > 0 && c.timeInitiated.Add(maxDuration).Before(time.Now()) {
		c.Unhealthy()
	}
	wrapper := ClientConn{
		ClientConn: c.ClientConn,
		pool:       c.pool,
		timeUsed:   time.Now(),
	}
	if c.unhealthy {
		_ = wrapper.ClientConn.Close()
		wrapper.ClientConn = nil
	} else {
		wrapper.timeInitiated = c.timeInitiated
	}
	select {
	case c.pool.clients <- wrapper:
	default:
		return ErrFullPool
	}
	c.ClientConn = nil
	return nil
}

func (p *Pool) Capacity() int {
	if p.IsClosed() {
		return 0
	}
	return cap(p.clients)
}

// Available 可用数量
func (p *Pool) Available() int {
	if p.IsClosed() {
		return 0
	}
	return len(p.clients)
}

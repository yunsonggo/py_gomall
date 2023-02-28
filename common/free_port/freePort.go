package free_port

import "net"

func FreePort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, nil
	}
	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, nil
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port, nil
}

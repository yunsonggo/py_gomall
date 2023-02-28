package user_config

var AppConf = new(Conf)

type Conf struct {
	Name       string   `mapstructure:"name" json:"name" yaml:"name"`
	Mode       string   `mapstructure:"mode" json:"mode" yaml:"mode"`
	Listen     string   `mapstructure:"listen" json:"listen" yaml:"listen"`
	Port       int      `mapstructure:"port" json:"port" yaml:"port"`
	FreePort   bool     `mapstructure:"free_port" json:"free_port" yaml:"free_port"`
	Tage       []string `mapstructure:"tage" json:"tage" yaml:"tage"`
	Services   []string `mapstructure:"services" json:"services" yaml:"services"`
	*LogConf   `mapstructure:"log" json:"log"`
	*Jwt       `mapstructure:"jwt" json:"jwt"`
	*RedisConf `mapstructure:"redis" json:"redis" yaml:"redis"`
	*NacosConf `mapstructure:"user_web_nacos" json:"user_web_nacos" yaml:"user_web_nacos"`
	*Consul    `mapstructure:"consul" json:"consul" yaml:"consul"`
}

// LogConf 日志配置模型
type LogConf struct {
	Level        string `mapstructure:"level" json:"level" yaml:"level"`
	Path         string `mapstructure:"path" json:"path" yaml:"path"`
	MaxSizeMB    int    `mapstructure:"max_size_mb" json:"max_size_mb" yaml:"max_size_mb"`
	MaxAgeDay    int    `mapstructure:"max_age_day" json:"max_age_day" yaml:"max_age_day"`
	MaxBackupDay int    `mapstructure:"max_backup_day" json:"max_backup_day" yaml:"max_backup_day"`
	Compress     bool   `mapstructure:"compress" json:"compress" yaml:"compress"`
}

type Jwt struct {
	Key          string `mapstructure:"key" json:"key" yaml:"key"`
	Issuer       string `mapstructure:"issuer" json:"issuer" yaml:"issuer"`
	Subject      string `mapstructure:"subject" json:"subject" yaml:"subject"`
	ExpireMinute int    `mapstructure:"expire_minute" json:"expire_minute" yaml:"expire_minute"`
	ExpireDay    int    `mapstructure:"expire_day" json:"expire_day" yaml:"expire_day"`
}

// NacosConf nacos配置模型
type NacosConf struct {
	NacosAddr   string `mapstructure:"nacos_addr" json:"nacos_addr" yaml:"nacos_addr"`
	NacosPort   uint64 `mapstructure:"nacos_port" json:"nacos_port" yaml:"nacos_port"`
	NamespaceId string `mapstructure:"namespace_id" json:"namespace_id" yaml:"namespace_id"`
	TimeoutMs   uint64 `mapstructure:"timeout_ms" json:"timeout_ms" yaml:"timeout_ms"`
	LoadCache   bool   `mapstructure:"load_cache" json:"load_cache" yaml:"load_cache"`
	LogDir      string `mapstructure:"log_dir" json:"log_dir" yaml:"log_dir"`
	CacheDir    string `mapstructure:"cache_dir" json:"cache_dir" yaml:"cache_dir"`
	RotateTime  string `mapstructure:"rotate_time" json:"rotate_time" yaml:"rotate_time"`
	MaxAge      int64  `mapstructure:"max_age" json:"max_age" yaml:"max_age"`
	LogLevel    string `mapstructure:"log_level" json:"log_level" yaml:"log_level"`
	DataId      string `mapstructure:"data_id" json:"data_id" yaml:"data_id"`
	Group       string `mapstructure:"group" json:"group" yaml:"group"`
}

type RedisConf struct {
	RedisAddr          string `mapstructure:"redis_addr" json:"red/Users/yunsong/go/src/2021/mallstore/webs/dial/registerConsole.gois_addr" yaml:"redis_addr"`
	RedisUser          string `mapstructure:"redis_user" json:"redis_user" yaml:"redis_user"`
	RedisPort          int    `mapstructure:"redis_port" json:"redis_port" yaml:"redis_port"`
	RedisPass          string `mapstructure:"redis_pass" json:"redis_pass" yaml:"redis_pass"`
	RedisDB            int    `mapstructure:"redis_db" json:"redis_db" yaml:"redis_db"`
	WithExpirySecond   int    `mapstructure:"with_expiry_second" json:"with_expiry_second" yaml:"with_expiry_second"`
	WithTriesTimes     int    `mapstructure:"with_tries_times" json:"with_tries_times" yaml:"with_tries_times"`
	ReadTimeoutSecond  int    `mapstructure:"read_timeout_second" json:"read_timeout_second" yaml:"read_timeout_second"`
	WriteTimeoutSecond int    `mapstructure:"write_timeout_second" json:"write_timeout_second" yaml:"write_timeout_second"`
	PoolSize           int    `mapstructure:"pool_size" json:"pool_size" yaml:"pool_size"`
	MinIdleConns       int    `mapstructure:"min_idle_conns" json:"min_idle_conns" yaml:"min_idle_conns"`
	MaxConnAgeMinute   int    `mapstructure:"max_conn_age_minute" json:"max_conn_age_minute" yaml:"max_conn_age_minute"`
	PoolTimeoutMinute  int    `mapstructure:"pool_timeout_minute" json:"pool_timeout_minute" yaml:"pool_timeout_minute"`
	IdleTimeoutMinute  int    `mapstructure:"idle_timeout_minute" json:"idle_timeout_minute" yaml:"idle_timeout_minute"`
}

type Consul struct {
	ConsulHost     string   `mapstructure:"consul_host" json:"consul_host"`
	ConsulPort     int      `mapstructure:"consul_port" json:"consul_port"`
	TimeoutSec     string   `json:"timeout_sec"`
	IntervalSec    string   `json:"interval_sec"`
	RemoveAfterSec string   `json:"remove_after_sec"`
	Tages          []string `json:"tages"`
}

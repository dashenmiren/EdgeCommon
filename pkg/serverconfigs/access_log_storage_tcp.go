package serverconfigs

// AccessLogTCPStorageConfig TCP存储策略
type AccessLogTCPStorageConfig struct {
	Network string `yaml:"network" json:"network"` // tcp, unix
	Addr    string `yaml:"addr" json:"addr"`
}

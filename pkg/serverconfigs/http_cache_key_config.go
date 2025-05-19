package serverconfigs

// HTTPCacheKeyConfig 缓存Key配置
type HTTPCacheKeyConfig struct {
	IsOn   bool   `yaml:"isOn" json:"isOn"`
	Scheme string `yaml:"scheme" json:"scheme"`
	Host   string `yaml:"host" json:"host"`
}

func (this *HTTPCacheKeyConfig) Init() error {
	return nil
}

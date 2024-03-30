package serverconfigs

const DefaultOpenFileCacheMax = 1024

// OpenFileCacheConfig open file cache配置
type OpenFileCacheConfig struct {
	IsOn bool `yaml:"isOn" json:"isOn"`
	Max  int  `yaml:"max" json:"max"`
}

func (this *OpenFileCacheConfig) Init() error {
	return nil
}

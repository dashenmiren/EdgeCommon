//go:build !plus

package ossconfigs

type OSSConfig struct {
}

func NewOSSConfig() *OSSConfig {
	return &OSSConfig{}
}

func (this *OSSConfig) Init() error {
	return nil
}

func (this *OSSConfig) Summary() string {
	return ""
}

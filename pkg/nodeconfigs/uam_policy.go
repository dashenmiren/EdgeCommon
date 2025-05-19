//go:build !plus

package nodeconfigs

func NewUAMPolicy() *UAMPolicy {
	return &UAMPolicy{}
}

type UAMPolicy struct {
	IsOn bool `yaml:"isOn" json:"isOn"`
}

func (this *UAMPolicy) Init() error {
	return nil
}

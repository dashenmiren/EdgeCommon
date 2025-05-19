package reporterconfigs

type NodeConfig struct {
	Id int64 `json:"id"`
}

func (this *NodeConfig) Init() error {
	return nil
}

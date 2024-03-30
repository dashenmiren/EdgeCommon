package firewallconfigs

type HTTPFirewallPost307Action struct {
	Life  int32         `yaml:"life" json:"life"`
	Scope FirewallScope `yaml:"scope" json:"scope"`
}

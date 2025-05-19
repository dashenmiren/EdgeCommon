package firewallconfigs

type HTTPFirewallGet302Action struct {
	Life  int32         `yaml:"life" json:"life"`
	Scope FirewallScope `yaml:"scope" json:"scope"`
}

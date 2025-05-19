package firewallconfigs

type FirewallScope = string

const (
	FirewallScopeGlobal  FirewallScope = "global"
	FirewallScopeService FirewallScope = "service"
)

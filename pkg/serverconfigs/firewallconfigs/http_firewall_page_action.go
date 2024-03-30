package firewallconfigs

import "net/http"

// HTTPFirewallPageAction default page action
type HTTPFirewallPageAction struct {
	IsPrior bool `yaml:"isPrior" json:"isPrior"`

	Status int    `yaml:"status" json:"status"`
	Body   string `yaml:"body" json:"body"`
}

func DefaultHTTPFirewallPageAction() *HTTPFirewallPageAction {
	return &HTTPFirewallPageAction{
		Status: http.StatusForbidden,
		Body: `<!DOCTYPE html>
<html lang="en">
<head>
	<title>403 Forbidden</title>
	<style>
		address { line-height: 1.8; }
	</style>
</head>
<body>
<h1>403 Forbidden By WAF</h1>
<address>Connection: ${remoteAddr} (Client) -&gt; ${serverAddr} (Server)</address>
<address>Request ID: ${requestId}</address>
</body>
</html>`,
	}
}

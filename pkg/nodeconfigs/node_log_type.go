package nodeconfigs

type NodeLogType = string

const (
	NodeLogTypeListenAddressFailed    NodeLogType = "listenAddressFailed"
	NodeLogTypeServerConfigInitFailed NodeLogType = "serverConfigInitFailed"
	NodeLogTypeRunScriptFailed        NodeLogType = "runScriptFailed"
	NodeLogTypeScriptConsoleLog       NodeLogType = "scriptConsoleLog"
)

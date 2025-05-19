package userconfigs

type UserModule = string

const (
	UserModuleCDN      UserModule = "cdn"
	UserModuleAntiDDoS UserModule = "antiDDoS"
	UserModuleNS       UserModule = "ns"
)

var DefaultUserModules = []UserModule{UserModuleCDN}

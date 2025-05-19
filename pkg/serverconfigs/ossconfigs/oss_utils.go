package ossconfigs

import "strings"

func IsOSSProtocol(protocol string) bool {
	return strings.HasPrefix(protocol, "oss:")
}

package shared_test

import (
	"testing"

	"github.com/dashenmiren/EdgeCommon/pkg/serverconfigs/shared"
)

func TestNewDataMap(t *testing.T) {
	var m = shared.NewDataMap()
	t.Log("data:", m.Read([]byte("e10adc3949ba59abbe56e057f20f883e")))
	var key = m.Put([]byte("123456"))
	t.Log("keyData:", key)
	t.Log("keyString:", string(key))
	t.Log("data:", string(m.Read(key)))
}

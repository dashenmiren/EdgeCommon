package nodeconfigs

import (
	"testing"
	"time"

	"github.com/iwind/TeaGo/logs"
)

func TestFindAllTimeZoneLocations(t *testing.T) {
	var before = time.Now()
	var locations = FindAllTimeZoneLocations()
	t.Log(len(locations), "locations")
	t.Logf("%.2f %s", time.Since(before).Seconds()*1000, "ms")
	logs.PrintAsJSON(locations, t)
}

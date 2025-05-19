package schedulingconfigs

import "testing"

func TestFindSchedulingType(t *testing.T) {
	t.Logf("%p", FindSchedulingType("roundRobin"))
	t.Logf("%p", FindSchedulingType("roundRobin"))
	t.Logf("%p", FindSchedulingType("roundRobin"))
}

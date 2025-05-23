package schedulingconfigs

import (
	"github.com/dashenmiren/EdgeCommon/pkg/serverconfigs/shared"
	"github.com/iwind/TeaGo/maps"
	"net/http"
	"testing"
)

func TestStickyScheduling_NextArgument(t *testing.T) {
	s := &StickyScheduling{}
	s.Add(&TestCandidate{
		Name:   "a",
		Weight: 1,
	})
	s.Add(&TestCandidate{
		Name:   "b",
		Weight: 2,
	})
	s.Add(&TestCandidate{
		Name:   "c",
		Weight: 3,
	})
	s.Add(&TestCandidate{
		Name:   "d",
		Weight: 6,
	})
	s.Start()

	t.Log(s.mapping)

	req, err := http.NewRequest(http.MethodGet, "http://www.example.com/?origin=c", nil)
	if err != nil {
		t.Fatal(err)
	}

	options := maps.Map{
		"type":  "argument",
		"param": "origin",
	}
	call := shared.NewRequestCall()
	call.Request = req
	call.Options = options
	t.Log(s.Next(call))
	t.Log(options)
}

func TestStickyScheduling_NextCookie(t *testing.T) {
	s := &StickyScheduling{}
	s.Add(&TestCandidate{
		Name:   "a",
		Weight: 1,
	})
	s.Add(&TestCandidate{
		Name:   "b",
		Weight: 2,
	})
	s.Add(&TestCandidate{
		Name:   "c",
		Weight: 3,
	})
	s.Add(&TestCandidate{
		Name:   "d",
		Weight: 6,
	})
	s.Start()

	t.Log(s.mapping)

	req, err := http.NewRequest(http.MethodGet, "http://www.example.com/?origin=c", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.AddCookie(&http.Cookie{
		Name:  "origin",
		Value: "c",
	})

	options := maps.Map{
		"type":  "cookie",
		"param": "origin",
	}
	call := shared.NewRequestCall()
	call.Request = req
	call.Options = options
	t.Log(s.Next(call))
	t.Log(options)
}

func TestStickyScheduling_NextHeader(t *testing.T) {
	s := &StickyScheduling{}
	s.Add(&TestCandidate{
		Name:   "a",
		Weight: 1,
	})
	s.Add(&TestCandidate{
		Name:   "b",
		Weight: 2,
	})
	s.Add(&TestCandidate{
		Name:   "c",
		Weight: 3,
	})
	s.Add(&TestCandidate{
		Name:   "d",
		Weight: 6,
	})
	s.Start()

	t.Log(s.mapping)

	req, err := http.NewRequest(http.MethodGet, "http://www.example.com/?origin=c", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("origin", "c")

	options := maps.Map{
		"type":  "header",
		"param": "origin",
	}
	call := shared.NewRequestCall()
	call.Request = req
	call.Options = options
	t.Log(s.Next(call))
	t.Log(options)
}

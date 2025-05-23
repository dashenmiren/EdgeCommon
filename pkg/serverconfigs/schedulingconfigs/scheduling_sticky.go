package schedulingconfigs

import (
	"github.com/dashenmiren/EdgeCommon/pkg/serverconfigs/shared"
	"github.com/iwind/TeaGo/maps"
	"github.com/iwind/TeaGo/rands"
	"math/rand"
	"net/http"
	"time"
)

// StickyScheduling Sticky调度算法
type StickyScheduling struct {
	Scheduling

	count   int
	mapping map[string]CandidateInterface // code => candidate
}

// Start 启动
func (this *StickyScheduling) Start() {
	this.mapping = map[string]CandidateInterface{}
	for _, c := range this.Candidates {
		for _, code := range c.CandidateCodes() {
			this.mapping[code] = c
		}
	}

	this.count = len(this.Candidates)
	rand.Seed(time.Now().UnixNano())
}

// Next 获取下一个候选对象
func (this *StickyScheduling) Next(call *shared.RequestCall) CandidateInterface {
	if this.count == 0 {
		return nil
	}

	if call == nil || call.Options == nil {
		return this.Candidates[rands.Int(0, this.count-1)]
	}

	typeCode := call.Options.GetString("type")
	param := call.Options.GetString("param")

	if call.Request == nil {
		return this.Candidates[rand.Int()%this.count]
	}

	code := ""
	if typeCode == "cookie" {
		cookie, err := call.Request.Cookie(param)
		if err == nil {
			code = cookie.Value
		}
	} else if typeCode == "header" {
		code = call.Request.Header.Get(param)
	} else if typeCode == "argument" {
		code = call.Request.URL.Query().Get(param)
	}

	matched := false
	var c CandidateInterface = nil

	defer func() {
		if !matched && c != nil {
			codes := c.CandidateCodes()
			if len(codes) == 0 {
				return
			}
			if typeCode == "cookie" {
				call.AddResponseCall(func(resp http.ResponseWriter) {
					http.SetCookie(resp, &http.Cookie{
						Name:    param,
						Value:   codes[0],
						Path:    "/",
						Expires: time.Now().AddDate(0, 1, 0),
					})
				})
			} else {
				call.AddResponseCall(func(resp http.ResponseWriter) {
					resp.Header().Set(param, codes[0])
				})
			}
		}
	}()

	if len(code) == 0 {
		c = this.Candidates[rand.Int()%this.count]
		return c
	}

	found := false
	c, found = this.mapping[code]
	if !found {
		c = this.Candidates[rand.Int()%this.count]
		return c
	}

	matched = true
	return c
}

// Summary 获取简要信息
func (this *StickyScheduling) Summary() maps.Map {
	return maps.Map{
		"code":        "sticky",
		"name":        "Sticky算法",
		"description": "利用Cookie、URL参数或者HTTP Header来指定源站",
		"networks":    []string{"http"},
	}
}

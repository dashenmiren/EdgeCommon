// Copyright 2022 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://cdn.foyeseo.com .

package serverconfigs

import (
	"github.com/dashenmiren/EdgeCommon/pkg/serverconfigs/shared"
	"net/http"
	"regexp"
	"strings"
)

type UserAgentAction = string

const (
	UserAgentActionAllow UserAgentAction = "allow"
	UserAgentActionDeny  UserAgentAction = "deny"
)

type UserAgentKeyword struct {
	keyword string
	reg     *regexp.Regexp
}

func NewUserAgentKeyword(keyword string) *UserAgentKeyword {
	var reg *regexp.Regexp
	if strings.Contains(keyword, "*") {
		var pieces = strings.Split(keyword, "*")
		for index, piece := range pieces {
			pieces[index] = regexp.QuoteMeta(piece)
		}
		var newKeyword = strings.Join(pieces, ".*")
		r, _ := regexp.Compile("(?i)" + newKeyword)
		reg = r
	} else {
		r, _ := regexp.Compile("(?i)" + regexp.QuoteMeta(keyword))
		reg = r
	}

	return &UserAgentKeyword{
		keyword: keyword,
		reg:     reg,
	}
}

type UserAgentFilter struct {
	Keywords []string        `yaml:"keywords" json:"keywords"` // 关键词
	Action   UserAgentAction `yaml:"action" json:"action"`     // 动作

	keywords []*UserAgentKeyword
}

func (this *UserAgentFilter) Init() error {
	this.keywords = []*UserAgentKeyword{}
	for _, keyword := range this.Keywords {
		this.keywords = append(this.keywords, NewUserAgentKeyword(keyword))
	}

	return nil
}

func (this *UserAgentFilter) Match(userAgent string) bool {
	if len(this.Keywords) == 0 {
		return false
	}

	for _, keyword := range this.keywords {
		if len(keyword.keyword) == 0 {
			// 空白
			if len(userAgent) == 0 {
				return true
			}
		} else if keyword.reg != nil {
			if keyword.reg.MatchString(userAgent) {
				return true
			}
		} else if strings.Contains(userAgent, keyword.keyword) {
			return true
		}
	}

	return false
}

type UserAgentConfig struct {
	IsPrior bool               `yaml:"isPrior" json:"isPrior"`
	IsOn    bool               `yaml:"isOn" json:"isOn"`
	Filters []*UserAgentFilter `yaml:"filters" json:"filters"`

	OnlyURLPatterns   []*shared.URLPattern `yaml:"onlyURLPatterns" json:"onlyURLPatterns"`     // 仅限的URL
	ExceptURLPatterns []*shared.URLPattern `yaml:"exceptURLPatterns" json:"exceptURLPatterns"` // 排除的URL
}

func NewUserAgentConfig() *UserAgentConfig {
	return &UserAgentConfig{}
}

func (this *UserAgentConfig) Init() error {
	// filters
	for _, filter := range this.Filters {
		err := filter.Init()
		if err != nil {
			return err
		}
	}

	// url patterns
	for _, pattern := range this.ExceptURLPatterns {
		err := pattern.Init()
		if err != nil {
			return err
		}
	}

	for _, pattern := range this.OnlyURLPatterns {
		err := pattern.Init()
		if err != nil {
			return err
		}
	}

	return nil
}

func (this *UserAgentConfig) AllowRequest(req *http.Request) bool {
	if len(this.Filters) == 0 {
		return true
	}

	var userAgent = req.UserAgent()

	for _, filter := range this.Filters {
		if filter.Match(userAgent) {
			return filter.Action == UserAgentActionAllow
		}
	}

	return true
}

func (this *UserAgentConfig) MatchURL(url string) bool {
	// except
	if len(this.ExceptURLPatterns) > 0 {
		for _, pattern := range this.ExceptURLPatterns {
			if pattern.Match(url) {
				return false
			}
		}
	}

	// only
	if len(this.OnlyURLPatterns) > 0 {
		for _, pattern := range this.OnlyURLPatterns {
			if pattern.Match(url) {
				return true
			}
		}
		return false
	}

	return true
}

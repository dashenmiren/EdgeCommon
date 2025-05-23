// Copyright 2023 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://cdn.foyeseo.com .

package serverconfigs

type HTTPHTMLOptimizationConfig struct {
	HTTPBaseOptimizationConfig

	IsOn bool `yaml:"isOn" json:"isOn"`

	KeepComments            bool `yaml:"keepComments" json:"keepComments"`
	KeepConditionalComments bool `yaml:"keepConditionalComments" json:"keepConditionalComments"`
	KeepDefaultAttrVals     bool `yaml:"keepDefaultAttrVals" json:"keepDefaultAttrVals"`
	KeepDocumentTags        bool `yaml:"keepDocumentTags" json:"keepDocumentTags"`
	KeepEndTags             bool `yaml:"keepEndTags" json:"keepEndTags"`
	KeepQuotes              bool `yaml:"keepQuotes" json:"keepQuotes"`
	KeepWhitespace          bool `yaml:"keepWhitespace" json:"keepWhitespace"`
}

func NewHTTPHTMLOptimizationConfig() *HTTPHTMLOptimizationConfig {
	return &HTTPHTMLOptimizationConfig{
		KeepDefaultAttrVals: true,
		KeepDocumentTags:    true,
		KeepEndTags:         true,
		KeepQuotes:          true,
	}
}

func (this *HTTPHTMLOptimizationConfig) Init() error {
	err := this.HTTPBaseOptimizationConfig.Init()
	if err != nil {
		return err
	}
	return nil
}

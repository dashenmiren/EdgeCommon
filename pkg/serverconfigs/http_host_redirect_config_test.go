package serverconfigs

import "testing"

func TestHTTPHostRedirectConfig_Init_Regexp(t *testing.T) {
	{
		config := &HTTPHostRedirectConfig{
			BeforeURL:   "http://teaos.cn",
			MatchRegexp: true,
		}
		err := config.Init()
		if err != nil {
			t.Fatal(err)
		}
	}

	{
		config := &HTTPHostRedirectConfig{
			BeforeURL:   `http://(\w+).cn`,
			MatchRegexp: true,
		}
		err := config.Init()
		if err != nil {
			t.Fatal(err)
		}
	}
}

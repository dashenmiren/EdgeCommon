package langs

import "testing"

func TestMessageCode_For(t *testing.T) {
	defaultManager.AddLang("en-us").
		Set("name", "Lily")

	var messageCode MessageCode = "name"
	t.Log(messageCode.String(), string(messageCode))
	t.Log(messageCode.For("en-us"))
}

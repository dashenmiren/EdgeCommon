package nodeutils

import (
	"testing"

	"github.com/iwind/TeaGo/maps"
)

func TestBase64EncodeMap(t *testing.T) {
	{
		var m maps.Map
		encodedString, err := Base64EncodeMap(m)
		if err != nil {
			t.Fatal(err)
		}
		t.Log("encoded string:", encodedString)

		m, err = Base64DecodeMap(encodedString)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%#v", m)
	}

	{
		var m = maps.Map{}
		encodedString, err := Base64EncodeMap(m)
		if err != nil {
			t.Fatal(err)
		}
		t.Log("encoded string:", encodedString)

		m, err = Base64DecodeMap(encodedString)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%#v", m)
	}

	{
		var m = maps.Map{"userId": 1, "name": "李白"}
		encodedString, err := Base64EncodeMap(m)
		if err != nil {
			t.Fatal(err)
		}
		t.Log("encoded string:", encodedString)

		m, err = Base64DecodeMap(encodedString)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%#v", m)
	}
}

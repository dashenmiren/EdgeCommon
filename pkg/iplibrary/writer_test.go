package iplibrary_test

import (
	"bytes"
	"testing"

	"github.com/dashenmiren/EdgeCommon/pkg/iplibrary"
)

func TestNewWriter(t *testing.T) {
	var buf = &bytes.Buffer{}
	var writer = iplibrary.NewWriter(buf, &iplibrary.Meta{
		Author: "GoCDN <https://google.com>",
	})

	err := writer.WriteMeta()
	if err != nil {
		t.Fatal(err)
	}

	err = writer.Write("192.168.1.100", "192.168.1.100", 100, 200, 300, 400, 500)
	if err != nil {
		t.Fatal(err)
	}

	err = writer.Write("192.168.2.100", "192.168.3.100", 101, 201, 301, 401, 501)
	if err != nil {
		t.Fatal(err)
	}

	err = writer.Write("192.168.3.101", "192.168.3.101", 101, 201, 301, 401, 501)
	if err != nil {
		t.Fatal(err)
	}

	err = writer.Write("::1", "::2", 101, 201, 301, 401, 501)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(buf.String())
	t.Log("sum:", writer.Sum())
}

// Copyright 2022 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://cdn.foyeseo.com .

package iplibrary_test

import (
	"bytes"
	"github.com/dashenmiren/EdgeCommon/pkg/iplibrary"
	"github.com/iwind/TeaGo/logs"
	"testing"
)

func TestNewWriter(t *testing.T) {
	//write
	var buf = &bytes.Buffer{}
	var writer = iplibrary.NewWriterV1(buf, &iplibrary.Meta{
		Author: "GoEdge <https://cdn.foyeseo.com>",
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

	err = writer.Write("10.0.0.1", "10.0.0.2", 101, 201, 301, 401, 501)
	if err != nil {
		t.Fatal(err)
	}

	err = writer.Write("10.0.0.3", "10.0.0.4", 101, 201, 301, 401, 501)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(buf.String())
	t.Log("sum:", writer.Sum())

	// read
	reader, err := iplibrary.NewReaderV2(buf)
	if err != nil {
		t.Fatal(err)
	}

	logs.PrintAsJSON(reader.IPv4Items(), t)
	logs.PrintAsJSON(reader.IPv6Items(), t)

	_ = reader
}

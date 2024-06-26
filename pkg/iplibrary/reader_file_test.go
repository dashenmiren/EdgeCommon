package iplibrary_test

import (
	"encoding/json"
	"net"
	"testing"

	"github.com/dashenmiren/EdgeCommon/pkg/iplibrary"
	"github.com/iwind/TeaGo/maps"
)

func TestNewFileReader(t *testing.T) {
	reader, err := iplibrary.NewFileReader("./ip-20c1461c.db", "123456")
	if err != nil {
		t.Fatal(err)
	}

	for _, ip := range []string{
		"127.0.0.1",
		"192.168.0.1",
		"192.168.0.150",
		"8.8.8.8",
		"111.197.165.199",
		"175.178.206.125",
	} {
		var result = reader.Lookup(net.ParseIP(ip))
		if result.IsOk() {
			var data = maps.Map{
				"countryId":    result.CountryId(),
				"countryName":  result.CountryName(),
				"provinceId":   result.ProvinceId(),
				"provinceName": result.ProvinceName(),
				"cityId":       result.CityId(),
				"cityName":     result.CityName(),
				"townId":       result.TownId(),
				"townName":     result.TownName(),
				"providerId":   result.ProviderId(),
				"providerName": result.ProviderName(),
				"summary":      result.Summary(),
			}
			dataJSON, err := json.MarshalIndent(data, "", "  ")
			if err != nil {
				t.Fatal(err)
			}
			t.Log(ip, "=>", string(dataJSON))
		} else {
			t.Log(ip+":", "not found")
		}
	}
}

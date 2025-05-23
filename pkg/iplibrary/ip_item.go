// Copyright 2022 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://cdn.foyeseo.com .

package iplibrary

import (
	"github.com/iwind/TeaGo/types"
)

type ipv4ItemV1 struct {
	IPFrom uint32
	IPTo   uint32

	Region *ipRegion
}

type ipv6ItemV1 struct {
	IPFrom uint64
	IPTo   uint64

	Region *ipRegion
}

type ipv4ItemV2 struct {
	IPFrom [4]byte
	IPTo   [4]byte

	Region *ipRegion
}

type ipv6ItemV2 struct {
	IPFrom [16]byte
	IPTo   [16]byte

	Region *ipRegion
}

type ipRegion struct {
	CountryId  uint16
	ProvinceId uint16
	CityId     uint32
	TownId     uint32
	ProviderId uint16
}

func HashRegion(countryId uint16, provinceId uint16, cityId uint32, townId uint32, providerId uint16) string {
	var providerHash = ""
	if providerId > 0 {
		providerHash = "_" + types.String(providerId)
	}

	if townId > 0 {
		return "t" + types.String(townId) + providerHash
	}
	if cityId > 0 {
		return "c" + types.String(cityId) + providerHash
	}
	if provinceId > 0 {
		return "p" + types.String(provinceId) + providerHash
	}
	return "a" + types.String(countryId) + providerHash
}

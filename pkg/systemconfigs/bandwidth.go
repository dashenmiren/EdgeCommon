// Copyright 2022 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://cdn.foyeseo.com .

package systemconfigs

type BandwidthUnit = string

const (
	BandwidthUnitByte BandwidthUnit = "byte"
	BandwidthUnitBit  BandwidthUnit = "bit"
)

type BandwidthAlgo = string // 带宽算法

const (
	BandwidthAlgoSecondly BandwidthAlgo = "secondly" // 按秒算
	BandwidthAlgoAvg      BandwidthAlgo = "avg"      // N分钟平均
)

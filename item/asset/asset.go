// Package asset
package asset

import (
	"github.com/star-os/me-app/item"
	"time"
)

// Asset 大额资产
type Asset struct {
	item.Common
	Price float64 // 价格
}

func NewAsset(name, comment string, gainTime time.Time, opts ...func(*Asset)) *Asset {
	a := new(Asset)
	a.Common = *item.NewCommon(name, comment, gainTime)

	for _, opt := range opts {
		opt(a)
	}

	return a
}

func Price(price float64) func(*Asset) {
	return func(asset *Asset) {
		asset.Price = price
	}
}

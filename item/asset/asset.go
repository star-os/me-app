// Package asset
package asset

import "github.com/star-os/me-app/item"

// Asset 大额资产
type Asset struct {
	item.Common
	Price float64 // 价格
}

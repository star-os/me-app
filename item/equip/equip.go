package equip

import (
	"github.com/star-os/me-app/item"
	"time"
)

// Equip 装备定义
type Equip struct {
	item.Common
	Price      float64 // 价格
	ExpireTime float64 // 过期时间
	// 图片还不知道怎么添加
}

func NewEquip(name, comment string, gainTime time.Time, opts ...func(equip *Equip)) *Equip {
	e := new(Equip)
	e.Common = *item.NewCommon(name, comment, gainTime)
	for _, opt := range opts {
		opt(e)
	}
	return e
}

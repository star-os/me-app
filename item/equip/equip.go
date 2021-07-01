package equip

import (
	"github.com/star-os/me-app/item"
)

// Equip 装备定义
type Equip struct {
	item.Common
	// 图片还不知道怎么添加
}

type Opt func(*Equip)

func Name(name string) Opt {
	return func(e *Equip) {
		e.Name = name
	}
}

func Comment(comment string) Opt {
	return func(e *Equip) {
		e.Comment = comment
	}
}

func Price(price float64) Opt {
	return func(e *Equip) {
		e.Price = price
	}
}

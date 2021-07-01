// Package item 是单个物品的定义
package item

type Item interface {
}

type Common struct {
	Name    string  // 名称
	Comment string  // 评价
	Price   float64 // 价格
}

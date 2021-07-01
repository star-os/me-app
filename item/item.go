// Package item 是单个物品的定义
package item

import "time"

type Item interface {
}

// Common 是创建一个条目所必须的内容
type Common struct {
	Name     string    // 名称
	Comment  string    // 评价
	GainTime time.Time // 得到的时间
}

func NewCommon(name, comment string, gainTime time.Time) *Common {
	return &Common{Name: name, Comment: comment, GainTime: gainTime}
}

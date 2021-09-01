// Package item 是单个物品的定义
package item

import "time"

type Item interface {
	// 目前还没看到有什么意义
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

func (c *Common) SetName(name string) {
	c.Name = name
}

func (c *Common) SetComment(comment string) {
	c.Comment = comment
}

func (c *Common) SetGainTime(t time.Time) {
	// 考虑做一个默认时间
	c.GainTime = t
}

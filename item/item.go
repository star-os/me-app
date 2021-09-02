// Package item 是单个物品的定义
package item

import "time"

// Item 物品结构体
type Item struct {
	Name     string    // 名称，必需
	Comment  string    // 简介，非必需
	GainTime time.Time // 获得的时间，自动生成/手动输入，必需
	Price    float64   // 价格
	Image    float64   // 图片，我也没想好怎么整
	Status   int8      // 状态，如已售出、已用完等
}

// NewItem 是Item的构造函数，用于初始化其中的字段，乖梨来完成
func NewItem() *Item {
	return &Item{}
}



/*以下都是弃用内容，它的构造函数使用了一个比较高级的方法，
我的一粒里面有介绍，但是不用按照这个来实现，自己写就行了

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
*/

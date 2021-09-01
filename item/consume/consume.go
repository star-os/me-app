// Package consume 消费品
package consume

import (
	"github.com/star-os/me-app/item"
	"time"
)

type Consume struct {
	item.Common
	Price      float64       // 价格
	ExpireTime time.Duration // 使用时间
}

func Price(price float64) func(*Consume) {
	return func(c *Consume) {
		c.Price = price
	}
}

func ExpireTime(d time.Duration) func(*Consume) {
	return func(c *Consume) {
		c.ExpireTime = d
	}
}

func NewConsume(name, comment string, gainTime time.Time, opts ...func(*Consume)) *Consume {
	c := new(Consume)
	c.Common = *item.NewCommon(name, comment, gainTime)

	for _, opt := range opts {
		opt(c)
	}
	return c
}

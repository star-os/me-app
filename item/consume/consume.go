// Package consume 消费品
package consume

import (
	"github.com/star-os/me-app/item"
	"time"
)

type Consume struct {
	item.Common
	Price      float64   // 价格
	ExpireTime time.Time // 使用时间
}

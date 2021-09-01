package achievement

import (
	"github.com/star-os/me-app/item"
	"time"
)

// Achievement 成就内容
type Achievement struct {
	item.Common
}

func NewAchievement(name, comment string, gainTime time.Time) *Achievement {
	return &Achievement{
		Common: *item.NewCommon(name, comment, gainTime),
	}
}

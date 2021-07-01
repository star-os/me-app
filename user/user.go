package user

import (
	"github.com/star-os/me-app/item/achievement"
	"github.com/star-os/me-app/item/asset"
	"github.com/star-os/me-app/item/consume"
	"github.com/star-os/me-app/item/equip"
)

type User struct {
	Equips       []equip.Equip
	Assets       []asset.Asset
	Achievements []achievement.Achievement
	Consume      []consume.Consume
}

func InitUser() *User {
	u := new(User)
	// 检查一下文件，再考虑

	return u
}

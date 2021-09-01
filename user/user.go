package user

import (
	"fmt"
	"github.com/star-os/me-app/file"
	"github.com/star-os/me-app/item/achievement"
	"github.com/star-os/me-app/item/asset"
	"github.com/star-os/me-app/item/consume"
	"github.com/star-os/me-app/item/equip"
)

const (
	FileName = `userfile`
)

type User struct {
	Gen          int64 // 用于记录当前的代际，避免新文件覆盖老文件
	Equips       []equip.Equip
	Assets       []asset.Asset
	Achievements []achievement.Achievement
	Consumes     []consume.Consume
}

// Init 用于初始化User信息
func (u *User) Init() error {
	return file.ReadFromJson(FileName, u)
}

// Save 保存User信息到文件中
// 会对User.Generation与文件中的内容进行比较，确认是否需要更新
func (u *User) Save() error {
	u.Gen++
	var err error
	defer func() {
		if err != nil {
			u.Gen--
		}
	}()

	// 保存文件
	err = file.SaveAsJson(FileName, u)
	if err != nil {
		return fmt.Errorf("Save: %w", err)
	}
	return nil
}

// GetField 输入字段名，获取相应字段的内容
func (u *User) GetField(name string) interface{} {
	switch name {
	case `assets`:
		return u.Assets
	case `equips`:
		return u.Equips
	case `achievements`:
		return u.Achievements
	case `consumes`:
		return u.Consumes
	default:
		return nil
	}
}

// AddItem 添加某个内容
func (u *User) AddItem(item interface{}) {
	switch item.(type) {
	case asset.Asset:
		u.Assets = append(u.Assets, item.(asset.Asset))
	case equip.Equip:
		u.Equips = append(u.Equips, item.(equip.Equip))
	case achievement.Achievement:
		u.Achievements = append(u.Achievements, item.(achievement.Achievement))
	case consume.Consume:
		u.Consumes = append(u.Consumes, item.(consume.Consume))
	default:
		return
	}
}

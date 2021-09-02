package user

import "github.com/star-os/me-app/item"

/*
const (
	FileName = `userfile`
)
*/

// Group 分类，持有一组物品，以及这组物品的分组名
type Group struct {
	Name  string
	Items []item.Item
}

func NewGroup() *Group {
	return &Group{}
}

// SetName 设置分组名称
func (g Group) SetName() {

}

// New 向一组物品中添加一个新的物品
func (g Group) New() {

}

// Delete 删除一组物品中的一个物品
func (g Group) Delete() {
	// 想想，怎么删除一个指定的物品
}

// Sell 出售一个物品
func (g Group) Sell() {

}

// Rent 出租一个物品
func (g Group) Rent() {

}

// User 用户
type User struct {
	UserName string  // 用户名
	Groups   []Group // 用户持有物品的所有分类
}

func NewUser() *User {
	return &User{}
}

// NewGroup 添加一个新的分类
func (u User) NewGroup() {

}

// DeleteGroup 删除一个分类
func (u User) DeleteGroup() {

}

// SetUserName 设置用户名
func (u User) SetUserName() {

}

/* 下面的内容都是老黄历了
type User struct {
	// Gen          int64 // 用于记录当前的代际，避免新文件覆盖老文件
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
*/

package navigator

import (
	"github.com/bianxl-yy/GoBlog/internal/app/system/setting/base"
	"github.com/knadh/koanf/providers/confmap"
)

type Navigator struct {
	Id         string `json:"id"`
	AuthorId   string `json:"author_id"`
	CreateTime int64  `json:"create_time"`
	EditTime   int64  `json:"edit_time"`
	UpdateTime int64  `json:"update_time"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Link        string `json:"link"`
	Sort        uint   `json:"-"`
}

// TODO: 待添加
func (n *Navigator) SortItem() {

}

// Add 添加导航
func (n *Navigator) Add() error {
	navigators, err := n.Get()
	if err != nil {
		return err
	}

	navigators = append(navigators, n)

	return base.Config.Load(confmap.Provider(map[string]interface{}{
		"navigators": navigators,
	}, "."), nil)
}

// Get 读取导航
func (n *Navigator) Get() ([]*Navigator, error) {
	navigators := make([]*Navigator, 0)
	return navigators, base.Config.Unmarshal("navigators", navigators)
}

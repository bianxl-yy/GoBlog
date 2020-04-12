package model

import (
	"github.com/knadh/koanf/providers/confmap"
)

type Navigator struct {
	*Model
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

	return Config.Load(confmap.Provider(map[string]interface{}{
		"navigators": navigators,
	}, "."), nil)
}

// Get 读取导航
func (n *Navigator) Get() ([]*Navigator, error) {
	navigators := make([]*Navigator, 0)
	return navigators, Config.Unmarshal("navigators", navigators)
}

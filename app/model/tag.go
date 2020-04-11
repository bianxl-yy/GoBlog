package model

import (
	"net/url"
	"strings"
)

// Content Tag struct. It convert tag string to proper struct or link.
type Tag struct {
	Model
	Name string
	Slug string
	Cid  []string
}

var (
	tagArticleIndex map[string][]string // 标签内容导航
)

// Link returns tag name url-encoded link.
func (t *Tag) Link() string {
	return "/tag/" + url.QueryEscape(strings.Replace(t.Name, ".", "-", -1))
}

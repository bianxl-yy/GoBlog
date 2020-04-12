package model

import (
	"encoding/json"
	"errors"
)

// Content Tag struct. It convert tag string to proper struct or link.
type Tag struct {
	Model
	Name       string
	Slug       string
	ArticleNum uint
	Articles   []*Article
}

func (c *Tag) GetKey() (string, error) {
	return c.Id, nil
}

func (c *Tag) GetType() (string, error) {
	return "tag", nil
}

func (c *Tag) Content() ([]byte, error) {
	bytes, e := json.Marshal(c)
	if e != nil {
		contentType, _ := c.GetType()
		return nil, errors.New("json encode '" + contentType + "' error")
	}
	return bytes, nil
}

func (c *Tag) Unmarshal(data []byte) error {
	return json.Unmarshal(data, c)
}

// Get 读取标签信息
// isArticle 可指定是否加载标签相关文章
func (c *Tag) Get(isArticles ...bool) (err error) {

	if len(isArticles) <= 0 || len(isArticles) > 0 && isArticles[0] {
		return storageManage.Get(c)
	}

	// 加载标签信息
	if err := storageManage.Get(c); err != nil {
		return err
	}
	tagArticles := &TagArticles{ByTagId: c.Id}
	// 加载相关文章
	c.Articles, err = tagArticles.Get()
	if err != nil {
		return err
	}

	return nil
}

// Remove 删除标签相关记录
func (c *Tag) Remove() (err error) {
	if err = storageManage.Remove(c); err != nil {
		return nil
	}

	// 删除标签相关文章记录
	tagArticles := &TagArticles{ByTagId: c.Id}
	return tagArticles.Remove()
}

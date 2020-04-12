package model

import (
	"encoding/json"
	"errors"
)

// Content Tag struct. It convert tag string to proper struct or link.
type TagArticles struct {
	Model
	ByTagId    string
	ArticlesId []string
}

func (c *TagArticles) GetKey() (string, error) {
	return c.ByTagId, nil
}

func (c *TagArticles) GetType() (string, error) {
	return "tag_articles", nil
}

func (c *TagArticles) Content() ([]byte, error) {
	bytes, e := json.Marshal(c)
	if e != nil {
		contentType, _ := c.GetType()
		return nil, errors.New("json encode '" + contentType + "' error")
	}
	return bytes, nil
}

func (c *TagArticles) Unmarshal(data []byte) error {
	return json.Unmarshal(data, c)
}

// Get 读取标签相关文章
func (c *TagArticles) Get() (result []*Article, err error) {
	if err := storageManage.Get(c); err != nil {
		return nil, err
	}

	for _, v := range c.ArticlesId {
		article := &Article{
			Model: &Model{Id: v},
		}
		if err := storageManage.Get(article); err != nil {
			return nil, err
		}
		result = append(result, article)
	}

	return result, nil
}

// Get 删除标签相关文章记录
func (c *TagArticles) Remove() (err error) {
	return storageManage.Remove(c)
}

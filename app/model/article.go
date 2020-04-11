package model

import (
	"errors"
	"github.com/bianxl-yy/GoBlog/app/utils"
	"github.com/rs/xid"
	"os"
	"path/filepath"
	"strings"
)

// Content instance, defines content data items.
type Content struct {
	Model
	Title      string     `json:"title"`
	Slug       string     `json:"slug"`
	Summary    string     `json:"summary"`
	Text       string     `json:"text"`
	Type       string     `json:"type"`
	Status     string     `json:"status"`
	Format     string     `json:"format"`
	Hits       uint       `json:"hits"`
	CommentNum uint       `json:"comment_num"`
	IsTop      bool       `json:"is_top"`
	IsComment  bool       `json:"is_comment"`
	Comments   []*Comment `json:"comments"`
	Tags       []*Tag     `json:"tags"`
}

func (c *Content) GetKey() string {
	return c.Id
}

func (c *Content) GetType() string {
	return c.Type
}

// CreateContent creates new content.
// t means content type, article or page.
func (c *Content) Write() (*Content, error) {
	// 验证Slug是否已存在
	if c.GetBySlug() != nil {
		return nil, errors.New("slug-repeat")
	}
	// 验证是否设置作者ID
	if c.AuthorId == "" {
		return nil, errors.New("AuthorId cannot be null")
	}
	// 验证内容类型
	if c.Type != "article" || c.Type != "page" {
		return nil, errors.New("content type error")
	}

	// 判断是否为更新
	if c.Id == "" {
		guid := xid.New()
		c.Id = guid.String()
		c.CreateTime = utils.Now()
		c.EditTime = c.CreateTime
		c.UpdateTime = c.CreateTime
	} else {
		c.UpdateTime = utils.Now()
	}

	return c, storageManage.Write(c)
}

// GetContentById gets a content by given id.
func (c *Content) Get() *Content {
	return contents[c.Id]
}

// GetBySlug gets a content by given slug.
func (c *Content) GetBySlug() *Content {
	for _, content := range contents {
		if content.Slug == c.Slug {
			return c
		}
	}
	return nil
}

// RemoveContent 删除内容
func (c *Content) Delete() error {
	c.Status = "DELETE"
	return storageManage.Write(c)
}

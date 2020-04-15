package page

import (
	"encoding/json"
	"errors"
	"github.com/bianxl-yy/GoBlog/app/utils"
	"github.com/bianxl-yy/GoBlog/internal/pkg/storage"
	"github.com/rs/xid"
)

// Content instance, defines content data items.
type Page struct {
	Id         string `json:"id"`
	AuthorId   string `json:"author_id"`
	CreateTime int64  `json:"create_time"`
	EditTime   int64  `json:"edit_time"`
	UpdateTime int64  `json:"update_time"`
	Title string `json:"title"`
	Slug  string `json:"slug"`
	Text string `json:"text"`
	Status     string `json:"status"`
	Format     string `json:"format"`
	Hits       uint   `json:"hits"`
	CommentNum uint   `json:"comment_num"`
	IsComment bool       `json:"is_comment"`
}

func (c *Page) GetKey() (string, error) {
	return c.Id, nil
}

func (c *Page) GetType() (string, error) {
	return "page", nil
}

func (c *Page) Content() ([]byte, error) {
	bytes, e := json.Marshal(c)
	if e != nil {
		contentType, _ := c.GetType()
		return nil, errors.New("json encode '" + contentType + "' error")
	}
	return bytes, nil
}

func (c *Page) Unmarshal(data []byte) error {
	return json.Unmarshal(data, c)
}

// CreateContent creates new content.
// t means content type, article or page.
func (c *Page) Write() (string, error) {
	// 验证Slug是否已存在
	// TODO: 待处理
	/*if c.GetBySlug() != nil {
		return nil, errors.New("slug-repeat")
	}
	*/
	// 验证是否设置作者ID
	if c.AuthorId == "" {
		return "", errors.New("AuthorId cannot be null")
	}

	// 判断是否为更新
	if c.Id == "" {
		guid := xid.New()
		c.Id = guid.String()
		//c.Type = "page"
		c.CreateTime = utils.Now()
		c.EditTime = c.CreateTime
		c.UpdateTime = c.CreateTime
	} else {
		c.UpdateTime = utils.Now()
	}

	return c.Id, storage.Write(c)
}

// GetContentById gets a content by given id.
func (c *Page) Get() error {
	return storageManage.Get(c)
}

// RemoveContent 删除内容
func (c *Page) Remove() error {
	c.Status = "DELETE"
	return storageManage.Write(c)
}

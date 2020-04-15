package content

import (
	"encoding/json"
	"errors"
	"github.com/bianxl-yy/GoBlog/app/utils"
	"github.com/rs/xid"
)

// Content instance, defines content data items.
type Article struct {
	Id         string `json:"id"`
	AuthorId   string `json:"author_id"`
	CreateTime int64  `json:"create_time"`
	EditTime   int64  `json:"edit_time"`
	UpdateTime int64  `json:"update_time"`
	Title   string `json:"title"`
	Slug    string `json:"slug"`
	Summary string `json:"summary"`
	Text    string `json:"text"`
	//Type       string     `json:"type"`
	Status     string     `json:"status"`
	Format     string     `json:"format"`
	Hits       uint       `json:"hits"`
	CommentNum uint       `json:"comment_num"`
	IsTop      bool       `json:"is_top"`
	IsComment  bool       `json:"is_comment"`
	Comments   []*Comment `json:"comments"`
	Tags       []*Tag     `json:"tags"`
}

func (c *Article) GetKey() (string, error) {
	return c.Id, nil
}

func (c *Article) GetType() (string, error) {
	return "article", nil
}

func (c *Article) Content() ([]byte, error) {
	bytes, err := json.Marshal(c)
	if err != nil {
		contentType, _ := c.GetType()
		return nil, errors.New("json encode '" + contentType + "' error")
	}
	return bytes, nil
}

func (c *Article) Unmarshal(data []byte) error {
	return json.Unmarshal(data, c)
}

// CreateContent creates new content.
// t means content type, article or page.
func (c *Article) Write() (string, error) {
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
		//c.Type = "article"
		c.CreateTime = utils.Now()
		c.EditTime = c.CreateTime
		c.UpdateTime = c.CreateTime
	} else {
		c.UpdateTime = utils.Now()
	}

	return c.Id, storageManage.Write(c)
}

// GetContentById gets a content by given id.
func (c *Article) Get() error {
	return storageManage.Get(c)
}

// RemoveContent 删除内容
func (c *Article) Remove() error {
	c.Status = "DELETE"
	return storageManage.Write(c)
}

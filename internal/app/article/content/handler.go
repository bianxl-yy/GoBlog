package content

import (
	"github.com/bianxl-yy/GoBlog/app/model"
	"github.com/labstack/echo"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Article 文章内容
func Get(c echo.Context) (err error) {
	slug := c.Param("slug")

	// 读取文章内容
	article := model.GetContentBySlug(slug)
	if article == nil || article.Type != "article" || article.Status != "publish" {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status": -1,
			"msg":    "页面不存在",
		})
	}

	// TODO: 待修改
	defer func() {
		if err == nil {
			article.Hits++
		}
	}()

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": 1,
		"msg":    "success",
		"data":   article,
	})
}

func Articles(c echo.Context) error {
	pageNum, _ := strconv.Atoi(c.Param("page"))
	size, _ := strconv.Atoi(model.GetSetting("article_size"))

	articles, pager := model.GetPublishArticleList(pageNum, size)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": 1,
		"msg":    "success",
		"data": map[string]interface{}{
			"articles": articles, // 文章列表
			"pager":    pager,    // 页码
		},
	})
}

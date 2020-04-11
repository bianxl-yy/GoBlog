package handler

import (
	"github.com/bianxl-yy/GoBlog/app/model"
	"github.com/labstack/echo"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Article 文章内容
func Article(c echo.Context) (err error) {
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

// ArticlesByTag 指定标签下的文章列表
func ArticlesByTag(c echo.Context) error {
	page, _ := strconv.Atoi(c.Param("page"))      // 页码
	name, _ := url.QueryUnescape(c.Param("name")) // 标签名称
	size, _ := strconv.Atoi(model.GetSetting("article_size"))

	articles, pager := model.GetTaggedArticleList(name, page, size)
	// fix dotted tag
	if len(articles) < 1 && strings.Contains(name, "-") {
		articles, pager = model.GetTaggedArticleList(strings.Replace(name, "-", ".", -1), page, size)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": 1,
		"msg":    "success",
		"data": map[string]interface{}{
			"Articles": articles,
			"Pager":    pager,
			"Tag":      name,
		},
	})
}

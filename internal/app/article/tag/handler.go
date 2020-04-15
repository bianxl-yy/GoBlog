package tag

import (
	"github.com/bianxl-yy/GoBlog/app/model"
	"github.com/labstack/echo"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// ArticlesByTag 指定标签下的文章列表
func Articles(c echo.Context) error {
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

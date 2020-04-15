package page

import (
	"github.com/bianxl-yy/GoBlog/app/model"
	"github.com/labstack/echo"
	"net/http"
)

func Get(c echo.Context) (err error) {
	slug := c.Param("slug")

	page := model.GetContentBySlug(slug)
	if page == nil || page.Status != "publish" || page.Type != "page" {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status": -1,
			"msg":    "页面不存在",
		})
	}

	// TODO: 待修改
	defer func() {
		if err == nil {
			page.Hits++
		}
	}()

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": 1,
		"msg":    "success",
		"data":   page,
	})
}

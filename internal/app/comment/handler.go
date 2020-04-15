package comment

import (
	"github.com/bianxl-yy/GoBlog/app/model"
	"github.com/bianxl-yy/GoBlog/app/utils"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func Get(c echo.Context) error {
	cid, _ := strconv.Atoi(c.Param("id"))
	if cid < 1 || model.GetContentById(cid) == nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status": -1,
			"msg":    "评论信息有误，请联系管理员",
		})
	}

	// 校验评论信息
	msg := validateComment(c)
	if msg != "" {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status": -1,
			"msg":    msg,
		})
	}

	pid, _ := strconv.Atoi(c.FormValue("pid"))
	email := c.FormValue("email")

	co := &model.Comment{
		Author:    c.FormValue("user"),
		Email:     email,
		Url:       c.FormValue("url"),
		Content:   c.FormValue("content"),
		Avatar:    utils.Gravatar(email, "50"),
		Pid:       pid,
		Ip:        c.RealIP(),
		UserAgent: c.Request().UserAgent(),
	}

	// 创建评论
	model.CreateComment(cid, co)
	// 创建后台通知
	model.CreateMessage("comment", co)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": 1,
		"msg":    "success",
		"data":   co,
	})
}

func validateComment(c echo.Context) string {
	user := c.FormValue("user")
	content := c.FormValue("content")
	email := c.FormValue("email")
	url := c.FormValue("url")
	if utils.IsEmptyString(user) || utils.IsEmptyString(content) {
		return "称呼，邮箱，内容必填"
	}
	if !utils.IsEmail(email) {
		return "邮箱格式错误"
	}
	if !utils.IsEmptyString(url) && !utils.IsURL(url) {
		return "网址格式错误"
	}
	return ""
}

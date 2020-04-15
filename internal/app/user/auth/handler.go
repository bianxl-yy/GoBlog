package auth

import (
	"github.com/bianxl-yy/GoBlog/app/model"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

// Login 用户登录
func Login(c echo.Context) error {
	user := model.GetUserByName(c.FormValue("user"))
	if user == nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status": -1,
			"msg":    "用户名不能为空",
		})
	}
	if !user.CheckPassword(c.FormValue("password")) {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status": -1,
			"msg":    "密码为空或错误",
		})
	}

	exp := 3600 * 24 * 3
	expStr := strconv.Itoa(exp)
	s := model.CreateToken(user, c, int64(exp))

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": 1,
		"msg":    "success",
		"data": map[string]interface{}{
			"token-user":  strconv.Itoa(s.UserId), // 用户名
			"token-value": s.Value,                // token
			"exp-time":    expStr,                 // 过期时间
		},
	})
}

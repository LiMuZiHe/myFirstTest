package controller

import (
	"net/http"
	"wxproject1/model"

	"github.com/gin-gonic/gin"
)

type Resp struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
}

func Login(c *gin.Context) {
	var user model.User
	if c.ShouldBind(&user) == nil {
		if user2, err := user.SelectSelf(); err == nil {
			if user2.Password == user.Password {
				c.JSON(http.StatusOK, Resp{Success: true, Msg: user.Role})
			} else {
				c.JSON(http.StatusOK, Resp{Success: false, Msg: "密码错误"})
			}

		} else {
			c.JSON(http.StatusOK, Resp{Success: false, Msg: err.Error()})
		}
	}
}

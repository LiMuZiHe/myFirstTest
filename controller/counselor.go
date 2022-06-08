package controller

import (
	"net/http"
	"wxproject1/dao"
	"wxproject1/model"

	"github.com/gin-gonic/gin"
)

type Counselor dao.Counselor

func GetCounselorInfo(c *gin.Context) {
	user := model.User{ID: "2233", Role: "counselor"}
	if counselorInfo, err := user.SelectCounselor(); err != nil {
		c.JSON(http.StatusOK, Resp{Success: false, Msg: err.Error()})
	} else {
		c.JSON(http.StatusOK, counselorInfo)
	}
}

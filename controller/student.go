package controller

import (
	"log"
	"net/http"
	"wxproject1/dao"
	"wxproject1/model"

	"github.com/gin-gonic/gin"
)

func GetStudentMin(c *gin.Context) {
	//var user model.User
	var studentMin dao.Person
	//id, _ := strconv.Atoi(user.ID)
	id := "123" //测试
	studentMin = model.SelectPersonMin(id)
	c.JSON(http.StatusOK, studentMin)
}

func StudentInfo(c *gin.Context) {
	id := "123"
	studentInfo := model.GetStudentInfo(id)
	c.JSON(http.StatusOK, studentInfo)
}
func UpdataStudentInfo(c *gin.Context) {
	var studentInfo model.Student
	if err := c.ShouldBind(&studentInfo); err != nil {
		c.JSON(http.StatusOK, Resp{Success: false, Msg: err.Error()})
	}
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusOK, Resp{Success: false, Msg: err.Error()})
	} else {
		log.Println(file.Filename)
	}

	studentInfo.UpdataInfo()
	c.JSON(http.StatusOK, Resp{Success: true})
}

func SubmitJourney(c *gin.Context) {
	var journey model.Journey
	if err := c.ShouldBind(&journey); err != nil {
		c.JSON(http.StatusOK, Resp{Success: false, Msg: err.Error()})
	} else {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusOK, Resp{Success: false, Msg: err.Error()})
		} else {
			journey.JourneyCard = file.Filename
			journey.Create()
			c.JSON(http.StatusOK, Resp{Success: true})
		}

	}
}
func SubmitReport(c *gin.Context) {
	var report model.Report
	if err := c.ShouldBind(&report); err != nil {
		c.JSON(http.StatusOK, Resp{Success: false, Msg: err.Error()})
	} else {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusOK, Resp{Success: false, Msg: err.Error()})
		} else {
			report.ReportScreenshot = file.Filename
			report.Create()
			c.JSON(http.StatusOK, Resp{Success: true})
		}

	}
}

package routers

import (
	"wxproject1/controller"
	"wxproject1/setting"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	if setting.Conf.Release {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	r.Static("/static", "static")

	// login
	r.POST("/login", controller.Login)

	// register
	r.POST("/register")
	student := r.Group("/student")
	{
		student.GET("", controller.GetStudentMin)
		student.GET("/information", controller.StudentInfo)
		student.POST("/information", controller.UpdataStudentInfo)
		student.POST("/journey", controller.SubmitJourney)
		student.POST("/report", controller.SubmitReport)
	}

	// teacher := r.Group("/teacher")
	// {

	// }

	counselor := r.Group("/counselor")
	{
		counselor.GET("", controller.GetCounselorInfo)
		//counselor.GET("/student-list",controller.GetStudent)
	}

	return r
}

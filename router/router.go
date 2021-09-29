package router

import (
	"IceBreaking/controller"
	"github.com/gin-gonic/gin"
)

type Router struct {
}

func (r *Router) Run(port int) {
	initRouter(port)
}

func initRouter(port int) {
	// [端口范围](https://blog.csdn.net/yyj108317/article/details/81134241)
	if port < 5000 || port > 65535 {
		panic("非法端口")
	}
	r := gin.Default()
	//student
	groupStudent := r.Group("/student")
	{
		groupStudent.GET("/all", requestEntry(controller.GetStudents))

		groupStudent.GET("/id", requestEntry(controller.GetStudentByUuid))

		groupStudent.GET("/rand", requestEntry(controller.GetRandStudentsWithPicture))

		groupStudent.POST("/add", requestEntry(controller.AddStudent))

		groupStudent.GET("/count", requestEntry(controller.CountStudents))

	}

	//picture
	groupPicture := r.Group("/picture")
	{
		groupPicture.GET("/verify", requestEntry(controller.VerifyPictureBelongToStudent))

		groupPicture.POST("/upload", requestEntry(controller.UploadPicture))
	}

	err := r.Run(":" + string(rune(port)))
	panic(err)
}

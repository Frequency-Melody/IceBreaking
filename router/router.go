package router

import (
	"IceBreaking/controller"
	"github.com/gin-gonic/gin"
)

type Router struct {
}

func (r *Router) Init() {
	initRouter()
}

func initRouter() {
	r := gin.Default()
	//student
	groupStudent := r.Group("/student")
	{
		groupStudent.GET("/all", requestEntry(controller.GetStudents))

		groupStudent.GET("/id", requestEntry(controller.GetStudentByUuid))

		groupStudent.GET("/rand", requestEntry(controller.GetRandStudentWithPicture))

		groupStudent.POST("/add", requestEntry(controller.AddStudent))

		groupStudent.GET("/count", requestEntry(controller.CountStudents))
	}

	//picture
	groupPicture := r.Group("/picture")
	{
		groupPicture.GET("/verify", requestEntry(controller.VerifyPictureBelongToStudent))
	}

	err := r.Run()
	panic(err)
}

func requestEntry(handler func(c *gin.Context) (int, interface{})) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(handler(c))
		c.Abort()
	}
}

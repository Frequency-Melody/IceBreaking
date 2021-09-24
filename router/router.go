// Package router 路由
package router

import (
	"IceBreaking/controller"
	"IceBreaking/response"
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
		groupStudent.GET("/all", response.RequestEntry(controller.GetStudents))

		groupStudent.GET("/id", response.RequestEntry(controller.GetStudentById))

		groupStudent.GET("/rand", response.RequestEntry(controller.GetRandStudentWithPicture))

		groupStudent.POST("/add", response.RequestEntry(controller.AddStudent))

		groupStudent.GET("/count", response.RequestEntry(controller.CountStudents)
	}

	//picture
	groupPicture := r.Group("/picture")
	{
		groupPicture.GET("/verify", controller.VerifyPictureBelongToStudent)
	}

	err := r.Run()
	panic(err)
}




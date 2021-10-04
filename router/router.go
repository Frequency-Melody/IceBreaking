package router

import (
	"IceBreaking/controller"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Router struct {
}

func (r *Router) Run(port int) {
	initRouter(port)
}

func initRouter(port int) {
	// [端口范围](https://blog.csdn.net/yyj108317/article/details/81134241)
	fmt.Println(port)
	if port < 5000 || port > 65535 {
		panic("非法端口")
	}
	r := gin.Default()
	r.Use(Cors())
	//r.Use(Auth())

	//auth
	groupRedirect := r.Group("/auth")
	{
		groupRedirect.GET("/auth", controller.Auth)
		groupRedirect.GET("/login", controller.Login)
		groupRedirect.GET("/validate", requestEntry(controller.Validate))
	}

	//student
	groupStudent := r.Group("/student")
	groupStudent.Use(Auth())
	{
		groupStudent.GET("/all", requestEntry(controller.GetStudents))

		// 获取自己的信息
		groupStudent.GET("/id", requestEntry(controller.GetStudentByUuid))

		groupStudent.GET("/rand", requestEntry(controller.GetRandStudentsWithPicture))

		groupStudent.POST("/add", requestEntry(controller.AddStudent))

		groupStudent.GET("/count", requestEntry(controller.CountStudents))

		// 获取某用户是否隐藏图片这一信息
		groupStudent.GET("/status", requestEntry(controller.GetPictureStatus))

		// 更新某用户是否隐藏图片这一信息
		groupStudent.POST("/status", requestEntry(controller.UpdatePictureStatus))

	}

	//picture
	groupPicture := r.Group("/picture")
	groupPicture.Use(Auth())
	{
		groupPicture.GET("/verify", requestEntry(controller.VerifyPictureBelongToStudent))

		groupPicture.POST("/upload", requestEntry(controller.UploadPicture))
	}

	err := r.Run(":" + strconv.Itoa(port))
	panic(err)
}

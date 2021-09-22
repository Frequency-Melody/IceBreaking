package router

import (
	"IceBreaking/crud"
	"IceBreaking/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Router struct {
}

func (r *Router) Init() {
	initRouter()
}

const (
	MIN_RAND_NUM = 2 // 每次最少随机的人数
)

func initRouter() {
	r := gin.Default()
	//student
	{
		r.GET("/stu/all", func(c *gin.Context) {
			c.JSON(http.StatusOK, service.GetStudents())
		})

		r.GET("/stu/id", func(c *gin.Context) {
			idString := c.DefaultQuery("id", "")
			if idString == "" {
				c.JSON(http.StatusBadRequest, service.MakeErrJson(service.ParamError()))
				return
			}
			id, err := strconv.Atoi(idString)
			if err != nil {
				c.JSON(http.StatusBadRequest, service.MakeErrJson(service.ParamError()))
				return
			}
			c.JSON(http.StatusOK, service.GetStudentById(id))
		})

		r.GET("/stu/rand", func(c *gin.Context) {
			// num 是每次返回的学生的数量，且不得小于 MIN_RAND_NUM
			numString := c.DefaultQuery("num", "")
			if numString == "" {
				c.JSON(http.StatusBadRequest, service.MakeErrJson(service.ParamError()))
				return
			}
			num, err := strconv.Atoi(numString)
			if err != nil {
				c.JSON(http.StatusBadRequest, service.MakeErrJson(service.ParamError()))
				return
			}
			if num < MIN_RAND_NUM {
				c.JSON(http.StatusBadRequest, service.MakeErrJson(service.RandNumTooSmallError()))
				return
			}
			c.JSON(http.StatusOK, service.GetRandStudentWithPicture(num))
			return
		})

		r.POST("/stu/add", func(c *gin.Context) {
			//name := c.DefaultPostForm("name", "")
			//department := c.DefaultPostForm("department", "")
			//hidePicStr := c.PostForm("hidePic")
			stu := crud.Student{}
			err := c.ShouldBindJSON(&stu)
			if err != nil {
				c.JSON(http.StatusBadRequest, service.MakeErrJson(service.ParamError()))
				return
			}

			c.JSON(http.StatusOK, service.AddStudent(&stu))
		})
	}

	err := r.Run()
	panic(err)
}

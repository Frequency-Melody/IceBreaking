package router

import (
	"IceBreaking/model"
	"IceBreaking/response"
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
		r.GET("/student/all", func(c *gin.Context) {
			c.JSON(http.StatusOK, service.GetStudents())
		})

		r.GET("/student/id", func(c *gin.Context) {
			idString := c.DefaultQuery("id", "")
			if idString == "" {
				c.JSON(http.StatusBadRequest, response.MakeErrJson(response.ParamError()))
				return
			}
			id, err := strconv.Atoi(idString)
			if err != nil {
				c.JSON(http.StatusBadRequest, response.MakeErrJson(response.ParamError()))
				return
			}
			c.JSON(http.StatusOK, service.GetStudentById(id))
		})

		r.GET("/student/rand", func(c *gin.Context) {
			// num 是每次返回的学生的数量，且不得小于 MIN_RAND_NUM
			numString := c.DefaultQuery("num", "")
			if numString == "" {
				c.JSON(http.StatusBadRequest, response.MakeErrJson(response.ParamError()))
				return
			}
			num, err := strconv.Atoi(numString)
			if err != nil {
				c.JSON(http.StatusBadRequest, response.MakeErrJson(response.ParamError()))
				return
			}
			if num < MIN_RAND_NUM {
				c.JSON(http.StatusBadRequest, response.MakeErrJson(response.RandNumTooSmallError()))
				return
			}
			c.JSON(http.StatusOK, service.GetRandStudentWithPicture(num))
			return
		})

		r.POST("/student/add", func(c *gin.Context) {
			//name := c.DefaultPostForm("name", "")
			//department := c.DefaultPostForm("department", "")
			//hidePicStr := c.PostForm("hidePic")
			stu := model.Student{}
			err := c.ShouldBindJSON(&stu)
			if err != nil {
				c.JSON(http.StatusBadRequest, response.MakeErrJson(response.ParamError()))
				return
			}

			c.JSON(http.StatusOK, service.AddStudent(&stu))
		})

		r.GET("/student/count", func(c *gin.Context) {
			c.JSON(http.StatusOK, service.CountStudents())
		})
	}

	//picture
	{
		r.GET("/picture/verify", func(c *gin.Context) {
			studentIdString := c.DefaultQuery("studentId", "")
			if studentIdString == "" {
				c.JSON(http.StatusBadRequest, response.MakeErrJson(response.ParamError()))
				return
			}
			studentId, err := strconv.Atoi(studentIdString)
			if err != nil {
				c.JSON(http.StatusBadRequest, response.MakeErrJson(response.ParamError()))
				return
			}

			pictureIdString := c.DefaultQuery("pictureId", "")
			if pictureIdString == "" {
				c.JSON(http.StatusBadRequest, response.MakeErrJson(response.ParamError()))
				return
			}
			pictureId, err := strconv.Atoi(pictureIdString)
			if err != nil {
				c.JSON(http.StatusBadRequest, response.MakeErrJson(response.ParamError()))
				return
			}

			c.JSON(http.StatusOK, service.VerifyPictureBelongToStudent(pictureId, studentId))
		})
	}

	err := r.Run()
	panic(err)
}

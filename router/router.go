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
	groupStudent := r.Group("/student")
	{
		groupStudent.GET("/all", func(c *gin.Context) {
			c.JSON(http.StatusOK, service.GetStudents())
		})

		groupStudent.GET("/id", func(c *gin.Context) {
			if id, err := strconv.Atoi(c.Query("id")); err != nil {
				c.JSON(http.StatusBadRequest, response.MakeErrJson(response.ParamError("id 不能为空")))
			} else {
				c.JSON(http.StatusOK, service.GetStudentById(id))
			}
		})

		groupStudent.GET("/rand", func(c *gin.Context) {
			// num 是每次返回的学生的数量，且不得小于 MIN_RAND_NUM
			numString := c.DefaultQuery("num", "")
			if numString == "" {
				c.JSON(http.StatusBadRequest, response.MakeErrJson(response.ParamError("num 不能为空")))
				return
			}
			num, err := strconv.Atoi(numString)
			if err != nil {
				c.JSON(http.StatusBadRequest, response.MakeErrJson(response.ParamError("num 必须为数字")))
				return
			}
			if num < MIN_RAND_NUM {
				c.JSON(http.StatusBadRequest, response.MakeErrJson(response.RandNumTooSmallError()))
				return
			}
			c.JSON(http.StatusOK, service.GetRandStudentWithPicture(num))
			return
		})

		groupStudent.POST("/add", func(c *gin.Context) {
			stu := model.Student{}
			err := c.ShouldBindJSON(&stu)
			if err != nil {
				c.JSON(http.StatusBadRequest, response.MakeErrJson(response.ParamError(err.Error())))
				return
			}

			c.JSON(http.StatusOK, service.AddStudent(&stu))
		})

		groupStudent.GET("/count", func(c *gin.Context) {
			c.JSON(http.StatusOK, service.CountStudents())
		})
	}

	//picture
	groupPicture := r.Group("/picture")
	{
		groupPicture.GET("/verify", func(c *gin.Context) {
			var studentId, pictureId int
			var err error
			if studentId, err = strconv.Atoi(c.Query("studentId")); err != nil {
				c.JSON(http.StatusBadRequest, response.MakeErrJson(response.ParamError(err.Error())))
				return
			}
			if pictureId, err = strconv.Atoi(c.Query("pictureId")); err != nil {
				c.JSON(http.StatusBadRequest, response.MakeErrJson(response.ParamError(err.Error())))
			}
			if err != nil {
				c.JSON(http.StatusBadRequest, response.MakeErrJson(response.ParamError(err.Error())))
				return
			}
			c.JSON(http.StatusOK, service.VerifyPictureBelongToStudent(studentId, pictureId))
		})
	}

	err := r.Run()
	panic(err)
}

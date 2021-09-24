package controller

import (
	"IceBreaking/model"
	"IceBreaking/response"
	"IceBreaking/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

const (
	MIN_RAND_NUM = 2 // 每次最少随机的人数
)

func GetStudents(_ *gin.Context) response.Response {
	return service.GetStudents()
}

func GetStudentById(c *gin.Context) response.Response {
	if id, err := strconv.Atoi(c.Query("id")); err != nil {
		return response.ParamError{}
	} else {
		return service.GetStudentById(id)
	}
}

func GetRandStudentWithPicture(c *gin.Context) (interface{}){
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
}

func AddStudent(c *gin.Context) (int, interface{}){
	student := model.Student{}
	err := c.ShouldBindJSON(&student)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.MakeErrJson(response.ParamError("")))
		return
	}

	c.JSON(http.StatusOK, service.AddStudent(&student))
}

func CountStudents(c *gin.Context) (int, interface{}){
	c.JSON(http.StatusOK, service.CountStudents())
}

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

func GetStudents(c *gin.Context) (int, interface{}) {
	return http.StatusOK, service.GetStudents()
}

func GetStudentByUuid(c *gin.Context) (int, interface{}) {
	uuid := c.Query("uuid")
	if uuid == ""{
		return http.StatusBadRequest, response.MakeErrJson(response.ParamError("uuid 不能为空"))
	} else {
		return http.StatusOK, service.GetStudentByUuid(uuid)
	}
}

func GetRandStudentWithPicture(c *gin.Context) (int, interface{}) {
	// num 是每次返回的学生的数量，且不得小于 MIN_RAND_NUM
	numString := c.DefaultQuery("num", "")
	if numString == "" {
		return http.StatusBadRequest, response.MakeErrJson(response.ParamError("num 不能为空"))
	}
	num, err := strconv.Atoi(numString)
	if err != nil {
		return http.StatusBadRequest, response.MakeErrJson(response.ParamError("num 必须为数字"))
	}
	if num < MIN_RAND_NUM {
		return http.StatusBadRequest, response.MakeErrJson(response.RandNumTooSmallError())
	}
	return http.StatusOK, service.GetRandStudentWithPicture(num)
}

func AddStudent(c *gin.Context) (int, interface{}) {
	stu := model.Student{}
	err := c.ShouldBindJSON(&stu)
	if err != nil {
		return http.StatusBadRequest, response.MakeErrJson(response.ParamError(err.Error()))
	}
	return http.StatusOK, service.AddStudent(&stu)
}

func CountStudents(c *gin.Context) (int, interface{}) {
	return http.StatusOK, service.CountStudents()
}

package controller

import (
	"IceBreaking/model"
	"IceBreaking/response"
	"IceBreaking/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

const (
	MIN_RAND_NUM = 2 // 每次最少随机的人数
)

// @Summary 获取所有学生
// @Description 获取所有学生
// @Accept json
// @Produce  json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /student/all [get]
func GetStudents(c *gin.Context) response.Response {
	return service.GetStudents()
}

func GetStudentByUuid(c *gin.Context) response.Response {
	uuid := c.Query("uuid")
	if uuid == "" {
		return response.LackUuidParamError
	} else {
		return service.GetStudentByUuid(uuid)
	}
}

func GetRandStudentsWithPicture(c *gin.Context) response.Response {
	// num 是每次返回的学生的数量，且不得小于 MIN_RAND_NUM
	numString := c.DefaultQuery("num", "")
	if numString == "" {
		return response.LackRandNumParamError
	}
	num, err := strconv.Atoi(numString)
	if err != nil {
		return response.ParamError
	}
	if num < MIN_RAND_NUM {
		return response.RandNumTooSmallError
	}
	return service.GetRandStudentsWithPicture(num)
}

func AddStudent(c *gin.Context) response.Response {
	stu := model.Student{}
	err := c.ShouldBindJSON(&stu)
	if err != nil {
		return response.ParamError
	}
	return service.AddStudent(&stu)
}

func CountStudents(c *gin.Context) response.Response {
	return service.CountStudents()
}

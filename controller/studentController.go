package controller

import (
	"IceBreaking/model"
	"IceBreaking/response"
	"IceBreaking/response/dto"
	"IceBreaking/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

const (
	MIN_RAND_NUM = 2 // 每次最少随机的人数
)

func GetStudents(c *gin.Context) response.Response {
	return service.GetStudents()
}

func GetStudentByUuid(c *gin.Context) response.Response {
	studentUuidInterface, _ := c.Get("uuid")
	studentUuid := studentUuidInterface.(string)
	//uuid := c.Query("uuid")
	return service.GetStudentByUuid(studentUuid)
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

// GetPictureStatus 获取图片显隐状态
func GetPictureStatus(c *gin.Context) response.Response  {
	studentUuidInterface, _ := c.Get("uuid")
	studentUuid := studentUuidInterface.(string)
	return service.GetPictureStatus(studentUuid)
}

// UpdatePictureStatus 更新学生是否隐藏图片
func UpdatePictureStatus(c *gin.Context) response.Response {
	studentUuidInterface, _ := c.Get("uuid")
	studentUuid := studentUuidInterface.(string)
	hidePicDto := dto.HidePicDto{}
	err := c.ShouldBindJSON(&hidePicDto)
	if err != nil {
		fmt.Println(err)
		return response.ParamError
	}
	return service.UpdatePictureStatus(studentUuid, hidePicDto.HidePic)
}

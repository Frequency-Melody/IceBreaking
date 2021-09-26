package controller

import (
	"IceBreaking/response"
	"IceBreaking/service"
	"github.com/gin-gonic/gin"
)

func VerifyPictureBelongToStudent(c *gin.Context) response.Response {
	studentUuid := c.Query("studentUuid")
	pictureUuid := c.Query("pictureUuid")
	if studentUuid == "" {
		return response.LackStudentUuidParamError
	}
	if pictureUuid == "" {
		return response.LackPictureUuidParamError
	}
	return service.VerifyPictureBelongToStudent(pictureUuid, studentUuid)

}

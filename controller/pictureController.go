package controller

import (
	"IceBreaking/log"
	"IceBreaking/response"
	"IceBreaking/service"
	"IceBreaking/util"
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

func UploadPicture(c *gin.Context) response.Response {
	//studentUuid := c.PostForm("studentUuid")
	studentUuidInterface, _ := c.Get("uuid")
	studentUuid := studentUuidInterface.(string)
	_, picture, err := c.Request.FormFile("picture")
	if err != nil {
		log.Sugar().Error("获取上传文件失败: %v", err)
		return response.FileUploadFailedError
	}
	if picture.Size > (2<<20)*32 {
		return response.FileTooLargeError
	}
	if !util.IsUploadPicture(picture.Header.Get("Content-Type")) {
		return response.NotImageError
	}

	return service.UploadPictureOfStudent(picture, studentUuid)
}

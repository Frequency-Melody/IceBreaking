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
	_, headers, err := c.Request.FormFile("picture")
	if err != nil {
		log.Sugar().Error("获取上传文件失败: %v", err)
	}
	if headers.Size > (2<<20) * 32 {
		return response.FileTooLargeError
	}
	if !util.IsUploadPicture(headers.Header.Get("Content-Type")) {
		return response.NotImageError
	}
	if err := c.SaveUploadedFile(headers, "./static/"+headers.Filename); err != nil {
		log.Sugar().Error(err)
		return response.FileUploadFailedError
	}
	return response.Success
}

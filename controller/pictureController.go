package controller

import (
	"IceBreaking/log"
	"IceBreaking/response"
	"IceBreaking/service"
	"IceBreaking/util"
	"bytes"
	"github.com/gin-gonic/gin"
	"io/ioutil"
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
	_, picture, err := c.Request.FormFile("picture")
	if err != nil {
		log.Sugar().Error("获取上传文件失败: %v", err)
	}
	if picture.Size > (2<<20)*32 {
		return response.FileTooLargeError
	}
	if !util.IsUploadPicture(picture.Header.Get("Content-Type")) {
		return response.NotImageError
	}
	fileHandle, err := picture.Open()
	if err != nil {
		log.Sugar().Error("流文件打开错误")
		return response.FileUploadFailedError
	}
	defer fileHandle.Close()
	fileByte, _ := ioutil.ReadAll(fileHandle)
	//上传到oss上
	return service.UploadFileToOss(picture.Filename, bytes.NewReader(fileByte))
}

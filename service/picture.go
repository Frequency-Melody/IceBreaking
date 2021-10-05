package service

import (
	"IceBreaking/crud"
	"IceBreaking/db"
	"IceBreaking/log"
	"IceBreaking/model"
	"IceBreaking/response"
	"IceBreaking/response/dto"
	"bytes"
	"io/ioutil"
	"mime/multipart"
)

// VerifyPictureBelongToStudent 验证某个图片是否属于某个学生，并返回学生信息
func VerifyPictureBelongToStudent(pictureUuid, studentUuid string) response.Response {
	student := &model.Student{}
	relationStudentPic := &model.RelationStudentPic{}
	// 通过关联表获取这张图片的正确的学生的信息的id
	db.Get().Where(&model.RelationStudentPic{PictureUuid: pictureUuid}).First(relationStudentPic)
	studentWhere := &model.Student{}
	// 通过 id 查询完整的学生信息
	studentWhere.Uuid = relationStudentPic.StudentUuid
	db.Get().Where(studentWhere).First(student)
	studentUuidNameDto := &dto.StudentUuidNameDto{Uuid: student.Uuid, Name: student.Name}
	if student.Uuid == studentUuid {
		return &dto.PictureVerifyDto{Verify: true, StudentInfo: studentUuidNameDto}
	} else {
		return &dto.PictureVerifyDto{Verify: false, StudentInfo: studentUuidNameDto}
	}
}

func UploadPictureOfStudent(picture *multipart.FileHeader, studentUuid string) response.Response {
	// 判断是否存在该学生
	student := crud.GetStudentByUuid(studentUuid)
	if student == nil || student.Uuid == "" {
		return response.NoStudentError
	}
	fileHandle, err := picture.Open()
	if err != nil {
		log.Sugar().Error("流文件打开错误")
		return response.FileUploadFailedError
	}
	defer fileHandle.Close()
	fileByte, _ := ioutil.ReadAll(fileHandle)
	res := UploadFileToOss(picture.Filename, bytes.NewReader(fileByte))
	// 如果报错就直接返回上传图片至阿里云时的错误
	if res.Error() != nil {
		return res
	}
	// 无果不报错，那么 res.Data() 返回的一定是 {"url": "图片url"}
	pictureUrlMap, _ := res.Data().(map[string]string)
	pictureUrl := pictureUrlMap["url"]
	pictureUuid := crud.UploadPicture(pictureUrl)
	if pictureUuid == "" {
		return response.MysqlInsertError
	}
	if !crud.CreateOrUpdateRelationOfPictureAndStudent(studentUuid, pictureUuid) {
		return response.MysqlInsertError
	}
	crud.UpdateStudentHasPic(student, true)
	return &dto.PictureUrlDto{Url: pictureUrl}
}

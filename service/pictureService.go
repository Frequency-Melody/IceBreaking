package service

import (
	"IceBreaking/crud"
	"IceBreaking/db"
	"IceBreaking/log"
	"IceBreaking/model"
	"IceBreaking/response"
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
	crud.SelectStudentInsensitiveFiled().Where(studentWhere).First(student)
	if student.Uuid == studentUuid {
		return response.PictureVerifyDto{Verify: true, StudentInfo: student}
	} else {
		return response.PictureVerifyDto{Verify: false, StudentInfo: student}
	}
}

func UploadPictureOfStudent(picture *multipart.FileHeader, studentUuid string) response.Response {
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
	//上传到oss上
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
	if !crud.RelatePictureAndStudent(studentUuid, pictureUuid) {
		return response.MysqlInsertError
	}
	return &response.PictureUrlDto{Url: pictureUrl}
}

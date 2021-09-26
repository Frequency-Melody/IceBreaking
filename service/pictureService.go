package service

import (
	"IceBreaking/crud"
	"IceBreaking/db"
	"IceBreaking/model"
	"IceBreaking/response"
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

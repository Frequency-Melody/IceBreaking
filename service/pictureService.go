package service

import (
	"IceBreaking/db"
	"IceBreaking/model"
	"IceBreaking/response"
	"github.com/gin-gonic/gin"
)

// VerifyPictureBelongToStudent 验证某个图片是否属于某个学生，并返回学生信息
func VerifyPictureBelongToStudent(pictureId, studentId int) response.JsonResponse {
	student := &model.Student{}
	relationStudentPic := &model.RelationStudentPic{}
	// 通过关联表获取这张图片的正确的学生的信息的id
	db.Get().Where(&model.RelationStudentPic{PictureId: pictureId}).First(relationStudentPic)
	studentWhere := &model.Student{}
	// 通过 id 查询完整的学生信息
	studentWhere.ID = relationStudentPic.StudentId
	db.Get().Where(studentWhere).First(student)
	if student.ID == studentId {
		return response.MakeSuccessJson(gin.H{"verify": "true", "studentInfo": student})
	} else {
		return response.MakeSuccessJson(gin.H{"verify": "false", "studentInfo": student})
	}
}
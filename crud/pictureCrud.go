// Package crud ：CRUD 操作
package crud

import (
	"IceBreaking/db"
	"IceBreaking/model"
)

func GetPictureByStudentId(studentId int) (pic model.Picture, err error) {
	rsp := &model.RelationStudentPic{}
	db.Get().Where(&model.RelationStudentPic{StudentId: studentId}).First(&rsp)
	pictureId := rsp.PictureId
	return GetPictureByPictureId(pictureId)
}

func GetPictureByPictureId(id int) (pic model.Picture, err error) {
	picWhere := &model.Picture{}
	picWhere.ID = id
	db.Get().Where(picWhere).First(&pic)
	return
}

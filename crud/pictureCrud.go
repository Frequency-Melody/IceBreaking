package crud

import (
	"IceBreaking/db"
	"IceBreaking/model"
	"gorm.io/gorm"
)

func GetPictureByStudentId(studentId int) (pic model.Picture) {
	rsp := &model.RelationStudentPic{}
	db.Get().Where(&model.RelationStudentPic{StudentId: studentId}).First(&rsp)
	pictureId := rsp.PictureId
	return GetPictureByPictureId(pictureId)
}

func GetPictureByPictureId(id int) (pic model.Picture) {
	db.Get().Where(&model.Picture{Model: gorm.Model{ID: uint(id)}}).First(&pic)
	return
}

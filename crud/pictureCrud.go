package crud

import (
	"IceBreaking/db"
	"gorm.io/gorm"
)

func GetPictureByStudentId(studentId int) (pic Picture) {
	rsp := &RelationStudentPic{}
	db.Get().Where(&RelationStudentPic{StudentId: studentId}).First(&rsp)
	pictureId := rsp.PictureId
	return GetPictureByPictureId(pictureId)
}

func GetPictureByPictureId(id int) (pic Picture) {
	db.Get().Where(&Picture{Model: gorm.Model{ID: uint(id)}}).First(&pic)
	return
}

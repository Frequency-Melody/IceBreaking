package crud

import (
	"IceBreaking/db"
	"IceBreaking/model"
)

func GetPictureByStudentId(studentId int) (pic model.Picture) {
	rsp := &model.RelationStudentPic{}
	db.Get().Where(&model.RelationStudentPic{StudentId: studentId}).First(&rsp)
	pictureId := rsp.PictureId
	return GetPictureByPictureId(pictureId)
}

func GetPictureByPictureId(id int) (pic model.Picture) {
	picWhere := &model.Picture{}
	picWhere.ID = id
	db.Get().Where(picWhere).First(&pic)
	return
}

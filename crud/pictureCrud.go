package crud

import (
	"IceBreaking/db"
	"IceBreaking/model"
)

func GetPictureByStudentUuid(studentUuid string) (picture model.Picture) {
	rsp := &model.RelationStudentPic{}
	db.Get().Where(&model.RelationStudentPic{StudentUuid: studentUuid}).First(&rsp)
	pictureUuid := rsp.PictureUuid
	return GetPictureByPictureUuid(pictureUuid)
}

func GetPictureByPictureUuid(uuid string) (picture model.Picture) {
	pictureWhere := &model.Picture{}
	pictureWhere.Uuid = uuid
	db.Get().Where(pictureWhere).First(&picture)
	return
}

package crud

import (
	"IceBreaking/db"
	"IceBreaking/log"
	"IceBreaking/model"
	"github.com/go-basic/uuid"
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

func UploadPicture(pictureUrl string) (pictureUuid string) {
	pictureUuid = uuid.New()
	picture := &model.Picture{}
	picture.Uuid = pictureUuid
	picture.Url = pictureUrl
	if err := db.Get().Create(picture).Error; err != nil {
		log.Sugar().Error("数据库插入错误：", err)
		return ""
	} else {
		return pictureUuid
	}
}

func RelatePictureAndStudent(studentUuid string, pictureUuid string) (ok bool) {
	err := db.Get().Create(&model.RelationStudentPic{StudentUuid: studentUuid, PictureUuid: pictureUuid}).Error
	if err != nil {
		log.Sugar().Error("创建学生-图片关联失败：", err)
		return false
	}
	return true
}

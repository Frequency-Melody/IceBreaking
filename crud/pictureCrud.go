package crud

import (
	"IceBreaking/db"
	"IceBreaking/log"
	"IceBreaking/model"
	"github.com/go-basic/uuid"
	"gorm.io/gorm"
)


func SelectPictureUuidAndUrlFiled() *gorm.DB {
	return db.Get().Select("url", "uuid")
}
func GetPictureByStudentUuid(studentUuid string) (picture model.Picture) {
	rsp := &model.RelationStudentPic{}
	db.Get().Where(&model.RelationStudentPic{StudentUuid: studentUuid}).First(&rsp)
	pictureUuid := rsp.PictureUuid
	return GetPictureByPictureUuid(pictureUuid)
}

func GetPictureByPictureUuid(uuid string) (picture model.Picture) {
	pictureWhere := &model.Picture{}
	pictureWhere.Uuid = uuid
	SelectPictureUuidAndUrlFiled().Where(pictureWhere).First(&picture)
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

// CreateRelationOfPictureAndStudent 在学生-图片 关联表中创建记录
func CreateRelationOfPictureAndStudent(studentUuid string, pictureUuid string) (ok bool) {
	relation := model.RelationStudentPic{}
	relation.Uuid = uuid.New()
	relation.StudentUuid = studentUuid
	relation.PictureUuid = pictureUuid
	err := db.Get().Create(&relation).Error
	if err != nil {
		log.Sugar().Error("创建学生-图片关联失败：", err)
		return false
	}
	log.Sugar().Info("在学生-图片关联表中创建记录")
	return true
}

// UpdateRelationOfPictureAndStudent 更新已有图片的学生的图片信息
func UpdateRelationOfPictureAndStudent(studentUuid string, pictureUuid string) (ok bool) {
	// 其实这个更新的话，要把旧照片删了更合理些（我没做
	relation := model.RelationStudentPic{}
	relation.StudentUuid = studentUuid
	err := db.Get().Model(&relation).Update("picture_uuid", pictureUuid).Error
	if err != nil {
		log.Sugar().Error("更新 relation_student_pics 表失败：", err)
		return false
	}
	return true
}

// CreateOrUpdateRelationOfPictureAndStudent 若学生无图片，在 学生-图片 关联表中添加记录；否则更新记录
func CreateOrUpdateRelationOfPictureAndStudent(studentUuid string, pictureUuid string) (ok bool) {
	relation := model.RelationStudentPic{}
	relation.StudentUuid = studentUuid
	db.Get().Where(&relation).Find(&relation)
	// relation 存在，更新记录
	if relation.Uuid != "" {
		err := db.Get().Model(&relation).Update("picture_uuid", pictureUuid).Error
		if err != nil {
			log.Sugar().Error("更新 relation_student_pics 表失败：", err)
			return false
		}
		log.Sugar().Info("在学生-图片关联表中更新记录")
		return true
	} else {
		// 否则创建记录
		return CreateRelationOfPictureAndStudent(studentUuid, pictureUuid)
	}
}

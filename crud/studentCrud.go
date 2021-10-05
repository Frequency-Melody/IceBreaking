package crud

import (
	"IceBreaking/db"
	"IceBreaking/log"
	"IceBreaking/model"

	"gorm.io/gorm"
)

// SelectStudentInsensitiveFiled 获取学生表的非敏感字段（已被 Trim 接口部分取代）
func SelectStudentInsensitiveFiled() *gorm.DB {
	return db.Get().Select("name", "Uuid")
}

func GetStudents() (students []*model.Student) {
	db.Get().Where(&model.Student{}).Find(&students)
	return
}

func CountStudents() (count int64) {
	db.Get().Model(&model.Student{}).Count(&count)
	return
}

func AddStudent(student *model.Student) (err error) {
	//student.Uuid = uuid.New()
	if err = db.Get().Create(student).Error; err != nil {
		log.Sugar().Info("数据库重复插入记录, StaffId:", student.StaffId, ", Name:", student.Name)
		return err
	}
	return nil
}

func GetStudentByUuid(studentUuid string) (stu *model.Student) {
	studentWhere := &model.Student{}
	studentWhere.Uuid = studentUuid
	db.Get().Where(studentWhere).First(&stu)
	return
}

func GetStudentByStaffId(staffId string) (stu *model.Student) {
	db.Get().Where(&model.Student{StaffId: staffId}).First(&stu)
	return
}

// GetStudentsCanBeShown 获取能展示图片的学生列表，即既有图片又不隐藏图片的人
func GetStudentsCanBeShown() (students []*model.Student) {
	whereQuery := map[string]interface{}{"hide_pic": false, "has_pic": true}
	db.Get().Where(whereQuery).Find(&students)
	return
}

func UpdateStudentHasPic(student *model.Student, hasPic bool) {
	db.Get().Model(&student).Update("has_pic", hasPic)
}

func UpdatePictureStatus(studentUuid string, hidePic bool) {
	student := GetStudentByUuid(studentUuid)
	db.Get().Model(&student).Update("hide_pic", hidePic)
}

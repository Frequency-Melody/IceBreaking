package crud

import (
	"IceBreaking/db"
	"IceBreaking/log"
	"IceBreaking/model"
	"github.com/go-basic/uuid"
	"gorm.io/gorm"
)

// SelectStudentInsensitiveFiled 获取学生表的非敏感字段
func SelectStudentInsensitiveFiled() *gorm.DB {
	return db.Get().Select("name", "Uuid")
}

func GetStudents() (students []*model.Student) {
	SelectStudentInsensitiveFiled().Where(&model.Student{}).Find(&students)
	return
}

func CountStudents() (count int64) {
	SelectStudentInsensitiveFiled().Model(&model.Student{}).Count(&count)
	return
}

func AddStudent(student *model.Student) (studentUuid string, err error) {
	if err = db.Get().Create(student).Error; err != nil {
		log.Sugar().Info("数据库重复插入记录, StaffId:", student.StaffId, ", Name:", student.Name)
		return uuid.New(), err
	}
	return student.Uuid, nil
}

func GetStudentByUuid(studentUuid string) (stu *model.Student) {
	studentWhere := &model.Student{}
	studentWhere.Uuid = studentUuid
	SelectStudentInsensitiveFiled().Where(studentWhere).First(&stu)
	return
}

// GetStudentsCanBeShown 获取能展示图片的学生列表，即既有图片又不隐藏图片的人
func GetStudentsCanBeShown() (students []*model.Student) {
	whereQuery := map[string]interface{}{"hide_pic": false, "has_pic": true}
	SelectStudentInsensitiveFiled().Where(whereQuery).Find(&students)
	return
}

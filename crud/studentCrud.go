package crud

import (
	"IceBreaking/db"
	"IceBreaking/model"
	"gorm.io/gorm"
	"log"
)

// SelectStudentInsensitiveFiled 获取学生表的非敏感字段
func SelectStudentInsensitiveFiled() *gorm.DB {
	return db.Get().Select("name", "ID")
}

func GetStudents() (students []*model.Student) {
	err := SelectStudentInsensitiveFiled().Where(&model.Student{}).Find(&students).Error
	if err != nil {
		log.Println(err)
	}
	return
}

func CountStudents() (count int64) {
	err := SelectStudentInsensitiveFiled().Model(&model.Student{}).Count(&count).Error
	if err != nil {
		log.Println(err)
	}
	return
}

func GetStudentIds() (studentIds []*model.StudentId) {
	err := db.Get().Where(map[string]interface{}{"hide_pic": false}).Find(&studentIds).Error
	if err != nil {
		log.Println(err)
	}
	return
}

func AddStudent(student *model.Student) (studentId int) {
	if err := db.Get().Create(student).Error; err != nil {
		log.Println(err)
		return
	}
	return student.ID
}

func GetStudentById(studentId int) (stu *model.Student) {
	studentWhere := &model.Student{}
	studentWhere.ID = studentId
	err := SelectStudentInsensitiveFiled().Where(studentWhere).First(&stu).Error
	if err != nil {
		log.Println(err)
	}
	return
}

func AddStudentId(studentId int, hidePic bool){
	err := db.Get().Create(&model.StudentId{StudentId: studentId, HidePic: hidePic}).Error
	if err != nil {
		log.Println(err)
	}
	return
}

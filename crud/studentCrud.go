package crud

import (
	"IceBreaking/db"
	"IceBreaking/model"
	"gorm.io/gorm"
)

func GetStudents() (students []*model.Student) {
	db.Get().Where(&model.Student{}).Find(&students)
	return
}

func CountStudents() (count int64) {
	db.Get().Model(&model.Student{}).Count(&count)
	return
}

func GetStudentIds() (studentIds []*model.StudentId) {
	db.Get().Where(map[string]interface{}{"hide_pic": false}).Find(&studentIds)
	return
}

func AddStudent(stu *model.Student) (studentId uint, err error) {
	if err = db.Get().Create(stu).Error; err != nil {
		return 0, err
	}
	return stu.ID, nil
}

func GetStudentById(studentId uint) (stu *model.Student) {
	db.Get().Where(&model.Student{Model: gorm.Model{ID: studentId}}).First(&stu)
	return
}

func AddStudentId(studentId int, hidePic bool) {
	db.Get().Create(&model.StudentId{StudentId: studentId, HidePic: hidePic})
	return
}

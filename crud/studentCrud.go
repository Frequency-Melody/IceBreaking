package crud

import (
	"IceBreaking/db"
	"gorm.io/gorm"
)

func GetStudents() (students []*Student) {
	db.Get().Where(&Student{}).Find(&students)
	return
}

func GetStudentIds() (studentIds []*StudentId) {
	db.Get().Where(map[string]interface{}{"hide_pic": false}).Find(&studentIds)
	return
}

func AddStudent(stu *Student) (studentId uint, err error) {
	if err = db.Get().Create(stu).Error; err != nil {
		return 0, err
	}
	return stu.ID, nil
}

func GetStudentById(studentId uint) (stu *Student) {
	db.Get().Where(&Student{Model: gorm.Model{ID: studentId}}).First(&stu)
	return
}

func AddStudentId(studentId int, hidePic bool) {
	db.Get().Create(&StudentId{StudentId: studentId, HidePic: hidePic})
	return
}

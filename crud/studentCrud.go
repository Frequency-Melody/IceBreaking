package crud

import (
	"IceBreaking/db"
	"IceBreaking/model"
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

func AddStudent(student *model.Student) (studentId int, err error) {
	if err = db.Get().Create(student).Error; err != nil {
		return 0, err
	}
	return student.ID, nil
}

func GetStudentById(studentId int) (stu *model.Student) {
	studentWhere := &model.Student{}
	studentWhere.ID = studentId
	db.Get().Where(studentWhere).First(&stu)
	return
}

func AddStudentId(studentId int, hidePic bool) {
	db.Get().Create(&model.StudentId{StudentId: studentId, HidePic: hidePic})
	return
}

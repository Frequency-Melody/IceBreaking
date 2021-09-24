package crud

import (
	"IceBreaking/db"
	"IceBreaking/model"
	"gorm.io/gorm"
)

// SelectStudentInsensitiveFiled 获取学生表的非敏感字段
func SelectStudentInsensitiveFiled() *gorm.DB {
	return db.Get().Select("name", "ID")
}

func GetStudents() (students []*model.Student) {
	SelectStudentInsensitiveFiled().Where(&model.Student{}).Find(&students)
	return
}

func CountStudents() (count int64) {
	SelectStudentInsensitiveFiled().Model(&model.Student{}).Count(&count)
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
	SelectStudentInsensitiveFiled().Where(studentWhere).First(&stu)
	return
}

func AddStudentId(studentId int, hidePic bool) {
	db.Get().Create(&model.StudentId{StudentId: studentId, HidePic: hidePic})
	return
}

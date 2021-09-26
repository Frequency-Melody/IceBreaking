package crud

import (
	"IceBreaking/db"
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

// GetStudentVos 获取学生表视图，存了学生 uuid 和 是否隐藏图片信息，用来快速随机学生
func GetStudentVos() (studentVos []*model.StudentVo) {
	db.Get().Where(map[string]interface{}{"hide_pic": false}).Find(&studentVos)
	return
}

func AddStudent(student *model.Student) (studentUuid string, err error) {
	if err = db.Get().Create(student).Error; err != nil {
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

func AddStudentVo(studentVo model.StudentVo) {
	db.Get().Create(&studentVo)
	return
}

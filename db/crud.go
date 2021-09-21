package db

import (
	"gorm.io/gorm"
)

func GetStudents() (students []*Student) {
	db.Where(&Student{}).Find(&students)
	return
}

func GetStudentIds() (studentIds []*StudentId) {
	db.Where(map[string]interface{}{"hide_pic": false}).Find(&studentIds)
	return
}

func AddStudent(stu *Student) (studentId uint, err error) {
	if err = db.Create(stu).Error; err != nil {
		return 0, err
	}
	return stu.ID, nil
}

func GetStudentById(studentId uint) (stu *Student) {
	db.Where(&Student{Model: gorm.Model{ID: studentId}}).First(&stu)
	return
}

func AddStudentId(studentId int, hidePic bool) {
	db.Create(&StudentId{StudentId: studentId, HidePic: hidePic})
	return
}

func GetPictureByStudentId(studentId int) (pic Picture) {
	rsp := &RelationStudentPic{}
	db.Where(&RelationStudentPic{StudentId: studentId}).First(&rsp)
	pictureId := rsp.PictureId
	return GetPictureByPictureId(pictureId)
}

func GetPictureByPictureId(id int) (pic Picture) {
	db.Where(&Picture{Model: gorm.Model{ID: uint(id)}}).First(&pic)
	return
}

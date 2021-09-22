package crud

import (
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name       string `gorm:"unique" binding:"required"`
	Department string `gorm:"comment:部门" binding:"required"`
	HidePic    bool   `gorm:"comment:是否隐藏照片"`
}

type Picture struct {
	gorm.Model
	Url string `gorm:"comment:图片在阿里云 OSS 中的地址"`
	//Student   Student
	//StudentId int
}

// 学生与照片的关联表
// 其实一对一关系，把 StuId 放在 Picture 字段里也行
// 但是为了防止传给前端的时候，被前端知道 StuId 不太好
// 再次横跳，其实返回的时候把 StuId 修剪一下也行
type RelationStudentPic struct {
	gorm.Model
	StudentId int
	PictureId int
}

// 一张表只存学生 id，这样能快速检索到有哪些学生
// 同时必须复制 HidePic 字段，否则可能查出的内容无效
type StudentId struct {
	gorm.Model
	StudentId int
	HidePic   bool
}

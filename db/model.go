package db

import (
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name string
	Department string `gorm:"comment:部门"`
	HidePic bool	`gorm:"comment:是否隐藏照片"`
}

type Picture struct {
	gorm.Model
	Url string `gorm:"comment:图片在阿里云 OSS 中的地址"`
	Student Student
	StudentId int
}

// 学生与照片的关联表
// 其实一对一关系，把 StuId 放在 Picture 字段里也行
// 但是为了防止传给前端的时候，被前端知道 StuId 不太好
type AssStuPic struct {
	gorm.Model
	StudentId int
	PictureId int
}
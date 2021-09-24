package model

import (
	"time"
)

type ModelWithoutDelete struct {
	ID        int       `gorm:"primarykey"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type Student struct {
	ModelWithoutDelete
	Name       string `binding:"required"`
	StaffId    int    `gorm:"unique" binding:"required" json:"staffId,omitempty"`
	Department string `gorm:"comment:部门" binding:"required" json:"department,omitempty"`
	HidePic    bool   `gorm:"comment:是否隐藏照片" json:"hidePic,omitempty"`
}

type Picture struct {
	ModelWithoutDelete
	Url string `gorm:"comment:图片在阿里云 OSS 中的地址"`
}

// RelationStudentPic 学生与照片的关联表
// 其实一对一关系，把 StuId 放在 Picture 字段里也行
// 但是为了防止传给前端的时候，被前端知道 StuId 不太好
// 再次横跳，其实返回的时候把 StuId 修剪一下也行
type RelationStudentPic struct {
	ModelWithoutDelete
	StudentId int `json:"studentId" binding:"required"`
	PictureId int `json:"pictureId" binding:"required"`
}

// StudentId 一张表只存学生 id，这样能快速检索到有哪些学生
// 同时必须复制 HidePic 字段，否则可能查出的内容无效
type StudentId struct {
	ModelWithoutDelete
	StudentId int
	HidePic   bool `json:"-"`
}
package model

import (
	"IceBreaking/err"
	"time"
)

type ModelWithoutDelete struct {
	ID        int       `gorm:"primaryKey" json:"id,omitempty"`
	Uuid      string    `gorm:"index"`
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

func (s Student) Error() error {
	if s.Uuid == "" {
		return err.DataEmptyError()
	}
	return nil
}

func (s Student) Code() int {
	if s.Uuid == "" {
		return 40400
	} else {
		return 20000
	}
}

func (s Student) Data() interface{} {
	return s
}

func (s Student) Redirect() string {
	return ""
}

type Picture struct {
	ModelWithoutDelete
	Url string `gorm:"comment:图片在阿里云 OSS 中的地址"`
}

func (p Picture) Error() error {
	if p.Uuid == "" {
		return err.DataEmptyError()
	}
	return nil
}

func (p Picture) Code() int {
	if p.Uuid == "" {
		return 40400
	} else {
		return 20000
	}
}

func (p Picture) Data() interface{} {
	return p
}

func (p Picture) Redirect() string {
	return ""
}

// RelationStudentPic 学生与照片的关联表
// 其实一对一关系，把 StuId 放在 Picture 字段里也行
// 但是为了防止传给前端的时候，被前端知道 StuId 不太好
// 再次横跳，其实返回的时候把 StuId 修剪一下也行
type RelationStudentPic struct {
	ModelWithoutDelete
	StudentUuid string `json:"studentUuid" binding:"required"`
	PictureUuid string `json:"pictureUuid" binding:"required"`
}

// StudentVo 一张表只存学生 id，这样能快速检索到有哪些学生
// 同时必须复制 HidePic 字段，否则可能查出的内容无效
type StudentVo struct {
	ModelWithoutDelete
	StudentUuid string
	HidePic     bool `json:"-"`
}

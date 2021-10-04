package model

import (
	"IceBreaking/errs"
	"github.com/go-basic/uuid"
	"gorm.io/gorm"
	"time"
)

type ModelWithoutDelete struct {
	ID        int       `gorm:"primaryKey" json:"id,omitempty"`
	Uuid      string    `gorm:"index,not null"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type Student struct {
	ModelWithoutDelete
	Name       string `binding:"required"`
	StaffId    string `gorm:"unique" binding:"required" json:"staffId"`
	Department string `gorm:"comment:部门" json:"department"`
	HidePic    bool   `gorm:"comment:是否隐藏照片" json:"hidePic"`
	HasPic     bool   `gorm:"comment:是否上传了照片" json:"hasPic"`
}

func (s *Student) BeforeCreate(db *gorm.DB) error {
	s.Uuid = uuid.New()
	return nil
}

func (s Student) Error() error {
	if s.Uuid == "" {
		return errs.DataEmptyError()
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

func (p *Picture) BeforeCreate(db *gorm.DB) error {
	p.Uuid = uuid.New()
	return nil
}

func (p Picture) Error() error {
	if p.Uuid == "" {
		return errs.DataEmptyError()
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

func (r *RelationStudentPic) BeforeCreate(db *gorm.DB) error {
	r.Uuid = uuid.New()
	return nil
}


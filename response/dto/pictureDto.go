package dto

import (
	"IceBreaking/err"
	"IceBreaking/response"
)

// PictureWithStudents 随机函数最终返回给前端的数据格式，包含一张照片和四个学生信息，实现 Response 接口
type PictureWithStudents struct {
	Picture  PictureUuidUrlDto
	Students []*StudentUuidNameDto
}

func (p PictureWithStudents) Error() error {
	if p.Picture.Uuid == "" {
		return err.DataEmptyError()
	}
	return nil
}

func (p PictureWithStudents) Code() int {
	if p.Picture.Uuid == "" {
		return 400000
	}
	return 20000
}

func (p PictureWithStudents) Data() interface{} {
	return p
}

func (p PictureWithStudents) Redirect() string {
	return ""
}

// PictureVerifyDto 图片验证 dto，实现 Response 接口
type PictureVerifyDto struct {
	response.BaseResponse `json:"-"`
	Verify                bool
	StudentInfo           *StudentUuidNameDto
}

func (d PictureVerifyDto) Data() interface{} {
	return map[string]interface{}{"verify": d.Verify, "studentInfo": d.StudentInfo}
}

// PictureUrlDto 图片在公网链接的 DTO，实现 Response 接口
type PictureUrlDto struct {
	response.BaseResponse `json:"-"`
	Url                   string
}

func (d *PictureUrlDto) Data() interface{} {
	return map[string]string{"url": d.Url}
}

type PictureUuidUrlDto struct {
	response.BaseResponse `json:"-"`
	Uuid                  string
	Url                   string
}

func (d PictureUuidUrlDto) Data() interface{} {
	return d
}

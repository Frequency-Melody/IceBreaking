package dto

import (
	"IceBreaking/response"
)

// StudentsDto 使 model.Student 切片实现 Response 接口
type StudentsDto struct {
	response.BaseResponse `json:"-"`
	Students              []*StudentUuidNameDto
}

func (s *StudentsDto) Data() interface{} {
	return s.Students
}

type StudentUuidNameDto struct {
	response.BaseResponse `json:"-"`
	Uuid                  string
	Name                  string
}

func (d *StudentUuidNameDto) Data() interface{} {
	return d
}

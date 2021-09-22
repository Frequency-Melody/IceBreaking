package service

import (
	"IceBreaking/crud"
)

type PictureWithStudents struct {
	Picture  *crud.Picture
	Students []*crud.Student
}

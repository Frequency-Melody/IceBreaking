package service

import "IceBreaking/db"

type PictureWithStudents struct {
	Picture  *db.Picture
	Students []*db.Student
}

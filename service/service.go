package service

import (
	"IceBreaking/db"
	"math/rand"
)

type Service struct {
	//DB     *gorm.DB
	//Router *gin.Engine
}

func (s *Service) GetStudents() JsonResponse {
	return s.MakeSuccessJson(db.GetStudents())
}

func (s *Service) GetStudentById(studentId int) JsonResponse {
	return s.MakeSuccessJson(db.GetStudentById(uint(studentId)))
}

func (s *Service) GetRandStudent() JsonResponse {
	stuIds := db.GetStudentIds()
	if len(stuIds) < 1 {
		return s.MakeErrJson(NoStudentError())
	}
	index := rand.Int() % len(stuIds)
	stu := db.GetStudentById(uint(stuIds[index].StudentId))
	return s.MakeSuccessJson(stu)
}

func (s *Service) GetRandStudentWithPicture(num int) JsonResponse {
	studentIds := db.GetStudentIds()
	studentNum := len(studentIds)
	if studentNum < 1 {
		return s.MakeErrJson(NoStudentError())
	}
	if studentNum < num {
		return s.MakeErrJson(NoEnoughStudentError())
	}
	// 抽取 n 个 studentId （的下标），即所有需要返回的学生的信息
	indexs := GetSomeRandNumber(num, 0, studentNum)
	students := make([]*db.Student, 0, num)
	for _, value := range indexs {
		student := db.GetStudentById(uint(studentIds[value].StudentId))
		students = append(students, student)
	}
	// 将第一个学生作为天选之子，返回图片
	// 也可以再次随机，util 里面有个 getOneRandNum
	selectedIndex := 0
	selectedStudentId := students[selectedIndex].ID
	picture := db.GetPictureByStudentId(int(selectedStudentId))
	return s.MakeSuccessJson(PictureWithStudents{Picture: &picture, Students: students})
}

func (s *Service) AddStudent(stu *db.Student) JsonResponse {
	id, err := db.AddStudent(stu)
	if err != nil {
		return s.MakeErrJson(StudentAlreadyExistError())
	}
	// 同步更新 StudentId 表，为了随机取学生时能快速查询
	db.AddStudentId(int(id), stu.HidePic)
	return s.MakeSuccessJson(map[string]int{"id": int(id)})
}

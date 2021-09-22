package service

import (
	"IceBreaking/crud"
	"IceBreaking/model"
	"IceBreaking/util"
	"math/rand"
)

func GetStudents() JsonResponse {
	return MakeSuccessJson(crud.GetStudents())
}

func GetStudentById(studentId int) JsonResponse {
	return MakeSuccessJson(crud.GetStudentById(studentId))
}

// GetRandStudent 获取一个随机学生，unused
func GetRandStudent() JsonResponse {
	stuIds := crud.GetStudentIds()
	if len(stuIds) < 1 {
		return MakeErrJson(NoStudentError())
	}
	index := rand.Int() % len(stuIds)
	stu := crud.GetStudentById(stuIds[index].StudentId)
	return MakeSuccessJson(stu)
}

// GetRandStudentWithPicture 随机 num 个学生，并且抽出一个人返回照片
func GetRandStudentWithPicture(num int) JsonResponse {
	studentIds := crud.GetStudentIds()
	studentNum := len(studentIds)
	if studentNum < 1 {
		return MakeErrJson(NoStudentError())
	}
	if studentNum < num {
		return MakeErrJson(NoEnoughStudentError())
	}
	// 抽取 n 个 studentId （的下标），即所有需要返回的学生的信息
	indexs := util.GetSomeRandNumber(num, 0, studentNum)
	students := make([]*model.Student, 0, num)
	for _, value := range indexs {
		student := crud.GetStudentById(studentIds[value].StudentId)
		students = append(students, student)
	}
	// 将第一个学生作为天选之子，返回图片
	// 也可以再次随机，util 里面有个 getOneRandNum
	selectedIndex := 0
	selectedStudentId := students[selectedIndex].ID
	picture := crud.GetPictureByStudentId(selectedStudentId)
	return MakeSuccessJson(model.PictureWithStudents{Picture: &picture, Students: students})
}

func AddStudent(stu *model.Student) JsonResponse {
	id, err := crud.AddStudent(stu)
	if err != nil {
		return MakeErrJson(StudentAlreadyExistError())
	}
	// 同步更新 StudentId 表，为了随机取学生时能快速查询
	crud.AddStudentId(id, stu.HidePic)
	return MakeSuccessJson(map[string]int{"id": id})
}

func CountStudents() JsonResponse {
	return MakeSuccessJson(map[string]interface{}{"count": crud.CountStudents()})
}

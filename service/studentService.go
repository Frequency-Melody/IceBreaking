package service

import (
	"IceBreaking/crud"
	"IceBreaking/model"
	"IceBreaking/response"
	"IceBreaking/util"
	"github.com/gin-gonic/gin"
	"github.com/go-basic/uuid"
	"math/rand"
)

func GetStudents() response.JsonResponse {
	return response.MakeSuccessJson(crud.GetStudents())
}

func GetStudentByUuid(studentUuid string) response.JsonResponse {
	return response.MakeSuccessJson(crud.GetStudentByUuid(studentUuid))
}

// GetRandStudent 获取一个随机学生，unused
func GetRandStudent() response.JsonResponse {
	studentVos := crud.GetStudentVos()
	if len(studentVos) < 1 {
		return response.MakeErrJson(response.NoStudentError())
	}
	index := rand.Int() % len(studentVos)
	stu := crud.GetStudentByUuid(studentVos[index].StudentUuid)
	return response.MakeSuccessJson(stu)
}

// GetRandStudentWithPicture 随机 num 个学生，并且抽出一个人返回照片
func GetRandStudentWithPicture(num int) response.JsonResponse {
	studentIds := crud.GetStudentVos()
	studentNum := len(studentIds)
	if studentNum < 1 {
		return response.MakeErrJson(response.NoStudentError())
	}
	if studentNum < num {
		return response.MakeErrJson(response.NoEnoughStudentError())
	}
	// 抽取 n 个 studentId （的下标），即所有需要返回的学生的信息
	indexs := util.GetSomeRandNumber(num, 0, studentNum)
	students := make([]*model.Student, 0, num)
	for _, value := range indexs {
		student := crud.GetStudentByUuid(studentIds[value].StudentUuid)
		students = append(students, student)
	}
	// 将第一个学生作为天选之子，返回图片
	// 也可以再次随机，util 里面有个 getOneRandNum
	selectedIndex := 0
	selectedStudentUuid := students[selectedIndex].Uuid
	picture := crud.GetPictureByStudentUuid(selectedStudentUuid)
	return response.MakeSuccessJson(model.PictureWithStudents{Picture: &picture, Students: students})
}

func AddStudent(student *model.Student) response.JsonResponse {
	student.Uuid = uuid.New()
	_, err := crud.AddStudent(student)
	if err != nil {
		return response.MakeErrJson(response.StudentAlreadyExistError())
	}
	// 同步更新 StudentId 表，为了随机取学生时能快速查询
	studentVo := model.StudentVo{StudentUuid: student.Uuid, HidePic: student.HidePic}
	crud.AddStudentVo(studentVo)
	return response.MakeSuccessJson(map[string]string{"uuid": student.Uuid})
}

func CountStudents() response.JsonResponse {
	return response.MakeSuccessJson(gin.H{"count": crud.CountStudents()})
}

// Package service 完成实际业务代码
package service

import (
	"IceBreaking/crud"
	"IceBreaking/model"
	"IceBreaking/response"
	"IceBreaking/util"
	"errors"
	"github.com/gin-gonic/gin"
	"math/rand"
)

func GetStudents() (data interface{}) {
	return crud.GetStudents()
}

func GetStudentById(studentId int) (response.Response) {
	return crud.GetStudentById(studentId)
}

// GetRandStudent 获取一个随机学生，unused
func GetRandStudent() (data interface{}) {
	stuIds, err := crud.GetStudentIds()
	if err != nil {
		return nil, err
	}
	if len(stuIds) < 1 {
		return response.NoStudentError(), errors.New("学生的长度至少为1")
	}
	index := rand.Int() % len(stuIds)
	return crud.GetStudentById(stuIds[index].StudentId)
}

// GetRandStudentWithPicture 随机 num 个学生，并且抽出一个人返回照片
func GetRandStudentWithPicture(num int) (data interface{}, err error) {
	studentIds, err := crud.GetStudentIds()
	studentNum := len(studentIds)
	if studentNum < 1 {
		return response.NoStudentError()), errors.New("没有该学生")
	}
	if studentNum < num {
		return response.NoEnoughStudentError(), error{"愿意公布照片的人数少于需要的随机人数"}
	}
	// 抽取 n 个 studentId （的下标），即所有需要返回的学生的信息
	indexs := util.GetSomeRandNumber(num, 0, studentNum)
	students := make([]*model.Student, 0, num)
	for _, value := range indexs {
		student, err := crud.GetStudentById(studentIds[value].StudentId)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}
	// 将第一个学生作为天选之子，返回图片
	// 也可以再次随机，util 里面有个 getOneRandNum
	selectedIndex := 0
	selectedStudentId := students[selectedIndex].ID
	picture, err := crud.GetPictureByStudentId(selectedStudentId)
	if err != nil {
		return nil, err
	}
	return model.PictureWithStudents{Picture: &picture, Students: students}, nil
}

func AddStudent(stu *model.Student) (data interface{}, err error) {
	id, err := crud.AddStudent(stu)
	if err != nil {
		return response.StudentAlreadyExistError(), err
	}
	// 同步更新 StudentId 表，为了随机取学生时能快速查询
	if err = crud.AddStudentId(id, stu.HidePic); err != nil {
		return nil, err
	} else {
		return map[string]int{"id": id}, nil
	}

}

func CountStudents() (data interface{}, err error) {
	if count, err := crud.CountStudents(); err != nil {
		return nil, err
	} else {
		return gin.H{"count": count}, nil
	}

}

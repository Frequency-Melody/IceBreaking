package service

import (
	"IceBreaking/crud"
	"IceBreaking/model"
	"IceBreaking/response"
	"IceBreaking/util"
	"fmt"
	"github.com/go-basic/uuid"
	"math/rand"
	"strconv"
)

func GetStudents() response.Response {
	return response.StudentsDto{Students: crud.GetStudents()}
}

func GetStudentByUuid(studentUuid string) response.Response {
	return crud.GetStudentByUuid(studentUuid)
}

// GetRandStudent 获取一个随机学生，unused
func GetRandStudent() response.Response {
	students := crud.GetStudents()
	index := rand.Int() % len(students)
	return students[index]
}

// GetRandStudentWithPicture 随机 num 个学生，并且抽出一个人返回照片
func GetRandStudentWithPicture(num int) response.Response {
	// 其实这里有个细节问题，理论上，只要有一个人上传图片，剩下的很多人都没有图片也能玩
	// 所以更合理的是计算学生总数，判断学生总数是否大于 rand num
	// 但是这样又会导致一个问题，随机 n 个人的时候，会很麻烦
	// 如果全抽有图片的，那会和之前利用学生总数判断的逻辑矛盾
	// 如果抽取时不加限定，那么，不能保证抽出来的 num 个人中有人有图片
	// 所以只能一直抽，直到 num 个人中有人有图片为止
	// 感觉没有必要想那么多，就默认不展示图片的不参与该游戏，名字也不出现在选项中
	// 即每次抽取的所有人，都是有图片的
	studentsCanBeShown := crud.GetStudentsCanBeShown()
	studentNum := len(studentsCanBeShown)
	if studentNum < 1 {
		return response.NoStudentError
	}
	if studentNum < num {
		return response.NoEnoughStudentError
	}
	// 抽取 n 个 studentId （的下标），即所有需要返回的学生的信息
	indexs := util.GetSomeRandNumber(num, 0, studentNum)
	// 随机出来的学生列表
	studentsRand := make([]*model.Student, 0, num)
	for _, value := range indexs {
		studentsRand = append(studentsRand, studentsCanBeShown[value])
	}
	// 再选个学生作为天选之子，返回图片
	selectedIndex := util.GetOneRandNum(0, num)
	fmt.Println(selectedIndex)
	selectedStudentUuid := studentsCanBeShown[selectedIndex].Uuid
	picture := crud.GetPictureByStudentUuid(selectedStudentUuid)
	return response.PictureWithStudents{Picture: &picture, Students: studentsCanBeShown}
}

func AddStudent(student *model.Student) response.Response {
	student.Uuid = uuid.New()
	_, err := crud.AddStudent(student)
	if err != nil {
		return response.StudentAlreadyExistError
	}
	return response.UuidDTO{Uuid: student.Uuid}
}

func CountStudents() response.Response {
	return &response.CountDto{Count: strconv.FormatInt(crud.CountStudents(), 10)}
}

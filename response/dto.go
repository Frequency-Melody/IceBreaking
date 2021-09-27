package response

import (
	"IceBreaking/err"
	"IceBreaking/model"
)

// JsonResponse Json 返回体结构
type JsonResponse struct {
	Error    int
	Msg      string
	Data     interface{}
	Redirect string
}

// PictureWithStudents 随机函数最终返回给前端的数据格式，包含一张照片和四个学生信息，实现 Response 接口
type PictureWithStudents struct {
	Picture  *model.Picture
	Students []*model.Student
}

func (p PictureWithStudents) Error() error {
	if p.Picture == nil || p.Picture.Uuid == "" {
		return err.DataEmptyError()
	}
	return nil
}

func (p PictureWithStudents) Code() int {
	if p.Picture == nil || p.Picture.Uuid == "" {
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

// UuidDTO 使 Uuid 单变量 实现 Response 接口
type UuidDTO struct {
	BaseResponse
	Uuid string
}

func (d UuidDTO) Data() interface{} {
	return map[string]string{"uuid": d.Uuid}
}

// StudentsDto 使 model.Student 切片实现 Response 接口
type StudentsDto struct {
	BaseResponse
	Students []*model.Student
}

func (s StudentsDto) Data() interface{} {
	return s.Students
}

// CountDto 数量 实现 Response 接口
type CountDto struct {
	BaseResponse
	Count string // 这里要用 string，因为 int 的默认值是 0，无法判断数据库查询是否成功
}

func (d *CountDto) Data() interface{} {
	return map[string]string{"count": d.Count}
}

func (d *CountDto) Error() error {
	if d.Count == "" {
		return err.MysqlQueryError()
	}
	return nil
}

// PictureVerifyDto 图片验证 dto，实现 Response 接口
type PictureVerifyDto struct {
	BaseResponse
	Verify      bool
	StudentInfo *model.Student
}

func (d PictureVerifyDto) Data() interface{} {
	return map[string]interface{}{"verify": d.Verify, "studentInfo": d.StudentInfo}
}

// PictureUrlDto 图片在公网链接的 DTO，实现 Response 接口
type PictureUrlDto struct {
	BaseResponse
	Url string
}

func (d *PictureUrlDto) Data() interface{} {
	return map[string]string{"url": d.Url}
}



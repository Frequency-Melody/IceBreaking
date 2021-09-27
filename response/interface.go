package response

import "github.com/gin-gonic/gin"

// Response 通用返回接口，内置错误，数据库 Model、DTO、常见错误都要实现这个接口
type Response interface {
	Error() error
	Code() int
	Data() interface{}
	Redirect() string
}

// Trimmer 修剪接口，用于默认剔除 Model 中的敏感字段
// 若实现了该接口，返回给前端时默认剔除敏感字段，防止 crud 忘记 select
type Trimmer interface {
	Trim() gin.H
}

// BaseResponse 一个实现了 Response 接口的基类，其他类可直接继承并重写部分接口
type BaseResponse int

func (b BaseResponse) Error() error {
	return nil
}

func (b BaseResponse) Code() int {
	return 20000
}

func (b BaseResponse) Data() interface{} {
	return nil
}

func (b BaseResponse) Redirect() string {
	return ""
}

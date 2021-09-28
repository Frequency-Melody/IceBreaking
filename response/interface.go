package response

// Response 通用返回接口，内置错误，数据库 Model、DTO、常见错误都要实现这个接口
type Response interface {
	Error() error
	Code() int
	Data() interface{}
	Redirect() string
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

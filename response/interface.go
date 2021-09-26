package response

type Response interface {
	Error() error
	Code() int
	Data() interface{}
	Redirect() string
}

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

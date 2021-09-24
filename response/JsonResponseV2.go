package response

import "errors"

type Response interface {
	Error() error
	Code()	int
	Data()	interface{}
	Redirect()	string
}

type ParamError struct {
}

func (p ParamError) Error() error {
	return errors.New("参数错误")
}

func (p ParamError) Code() int {
	return 40001
}

func (p ParamError) Data() interface{} {
	return nil
}

func (p ParamError) Redirect() string {
	return ""
}

//func MakeJsonFromErrMsg(errMsg ErrMsg, data interface{}, redirect string) JsonResponse {
//	return JsonResponse{errMsg.Error(), errMsg.Msg(), data, redirect}
//}
//
//func MakeErrJson(errMsg ErrMsg) (res JsonResponse) {
//	return MakeJsonFromErrMsg(errMsg, nil, "")
//}
//
//func MakeSuccessJson(data interface{}) (res JsonResponse) {
//	return MakeJsonFromErrMsg(&MsgSuccess{}, data, "")
//}

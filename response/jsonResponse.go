// Package response 返回体的 struct 以及封装的一些快捷函数
package response
//
//type JsonResponse struct {
//	Error string
//	Msg string
//	Data     interface{}
//	Redirect string
//}
//
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
//

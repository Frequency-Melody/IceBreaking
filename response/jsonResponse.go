package response

type JsonResponse struct {
	ErrMsg
	Data     interface{}
	Redirect string
}

func MakeErrJson(errMsg ErrMsg) (res JsonResponse) {
	return JsonResponse{errMsg, nil, ""}
}

func MakeSuccessJson(data interface{}) (res JsonResponse) {
	return JsonResponse{MsgSuccess(), data, ""}
}
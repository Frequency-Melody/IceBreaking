package response

type JsonResponse struct {
	Error    int
	Msg      string
	Data     interface{}
	Redirect string
}

func newJsonFromErrMsg(err ErrMsg, data interface{}) (res JsonResponse) {
	return JsonResponse{Error: err.Error, Msg: err.Msg, Data: data, Redirect: ""}
}

func MakeErrJson(err ErrMsg) (res JsonResponse) {
	return newJsonFromErrMsg(err, nil)
}

func MakeSuccessJson(data interface{}) (res JsonResponse) {
	return newJsonFromErrMsg(Success(), data)
}

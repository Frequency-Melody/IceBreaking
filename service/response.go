package service

type JsonResponse struct {
	Error    int
	Msg      string
	Data     interface{}
	Redirect string
}

func newJsonFromErrMsg(err ErrMsg, data interface{}) (res JsonResponse) {
	return JsonResponse{Error: err.Error, Msg: err.Msg, Data: data, Redirect: ""}
}

func (s *Service) MakeErrJson(err ErrMsg) (res JsonResponse) {
	return newJsonFromErrMsg(err, nil)
}

func (s *Service) MakeSuccessJson(data interface{}) (res JsonResponse) {
	return newJsonFromErrMsg(success(), data)
}

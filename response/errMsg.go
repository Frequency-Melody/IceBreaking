package response

type ErrMsg struct {
	Error int
	Msg   string
}

func Success() ErrMsg {
	return ErrMsg{Error: 20000, Msg: "ok"}
}

func ParamError() ErrMsg {
	return ErrMsg{Error: 40001, Msg: "参数错误"}
}

func NoStudentError() ErrMsg {
	return ErrMsg{Error: 40002, Msg: "学生为空"}
}

func StudentAlreadyExistError() ErrMsg {
	return ErrMsg{Error: 40003, Msg: "该学生已存在，请勿重复添加"}
}

func NoEnoughStudentError() ErrMsg {
	return ErrMsg{Error: 40004, Msg: "请求的学生数量大于数据库含有的学生总数"}
}

func RandNumTooSmallError() ErrMsg {
	return ErrMsg{Error: 40005, Msg: "每次随机的人数至少为 2"}
}

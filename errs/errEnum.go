package errs

type errEnum uint

func (e errEnum) Error() string {
	errMsg := map[errEnum]string{
		DataEmptyError:    "Data is empty",
		MysqlQueryError:   "数据库查询错误",
		InvalidTokenError: "无效的 token",
	}
	if msg, ok := errMsg[e]; ok {
		return msg
	} else {
		return "未知错误"
	}
}

const (
	DataEmptyError errEnum = iota
	MysqlQueryError
	InvalidTokenError
)

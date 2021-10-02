package errs

// 封装一些常见的错误，实现 error 接口（其实就是从 Go 源码复制的）
type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

func DataEmptyError() error {
	return &errorString{"Data is empty"}
}

func MysqlQueryError() error {
	return &errorString{"数据库查询错误"}
}

func InvalidTokenError() error {
	return &errorString{"无效的 token"}
}

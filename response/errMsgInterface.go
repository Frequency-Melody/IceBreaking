//package response
//
//type ErrMsg interface {
//	Error() string	//error code
//	Msg() string
//}
//
//type MsgSuccess struct {
//}
//
//func (*MsgSuccess)Error() string {
//	return "20000"
//}
//
//func (*MsgSuccess)Msg() string {
//	return "ok"
//}
//
//type ParamError struct {
//
//}
//
//func (p ParamError) Error() string {
//	return "40001"
//}
//
//func (p ParamError) Msg() string {
//	panic("参数错误")
//}
//
//
//

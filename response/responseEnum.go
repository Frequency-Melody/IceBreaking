package response

import (
	"IceBreaking/config"
	"errors"
	"github.com/go-basic/uuid"
	"net/url"
)

// ErrorType 实现 response.Response 接口
type ErrorType uint

func (e ErrorType) Error() error {
	errorMsg := map[ErrorType]string{
		// 成功
		Success: "OK",

		// 跳转
		RedirectToHduhelp: "跳转到杭电助手授权",

		// 参数错误
		ParamError:                "参数错误",
		RandNumTooSmallError:      "每次随机的人数过少",
		LackUuidParamError:        "uuid 不能为空",
		LackRandNumParamError:     "随机数 num 不能为空",
		LackStudentUuidParamError: "studentUuid 不能为空",
		LackPictureUuidParamError: "pictureUuid 不能为空",

		// 查询错误
		NoStudentError:           "学生为空",
		StudentAlreadyExistError: "该学生已存在，请勿重复添加",
		NoEnoughStudentError:     "分享照片的人数少于需要的随机人数",
		MysqlInsertError:         "数据库插入错误",

		// 文件相关错误
		FileTooLargeError:          "文件过大，仅支持最大 32M 文件",
		NotImageError:              "仅支持上传图片",
		FileUploadFailedError:      "文件上传错误",
		FileUploadToOssFailedError: "文件上传到阿里云 OSS 失败",

		// 鉴权错误
		LackAuthorizeTokenError: "AuthorizeToken 缺失",
		AuthorizeFailed:         "鉴权失败",
		InvalidTokenError:       "无效的 token",

		// 服务器错误
		ServerUnknownError: "服务器内部错误",
	}
	if m, ok := errorMsg[e]; ok {
		return errors.New(m)
	} else {
		return errors.New("undefined error")
	}
}

func (e ErrorType) Code() int {
	errorCode := map[ErrorType]int{
		// 成功
		Success: 20000,

		// 跳转
		RedirectToHduhelp: 30200,

		// 参数错误
		ParamError:                40001,
		RandNumTooSmallError:      40002,
		LackUuidParamError:        40003,
		LackRandNumParamError:     40004,
		LackStudentUuidParamError: 40005,
		LackPictureUuidParamError: 40006,

		// 查询错误
		NoStudentError:           40051,
		StudentAlreadyExistError: 40052,
		NoEnoughStudentError:     40053,
		MysqlInsertError:         40054,

		//文件相关错误
		FileTooLargeError:          40061,
		NotImageError:              40062,
		FileUploadFailedError:      40063,
		FileUploadToOssFailedError: 40064,

		//鉴权错误
		LackAuthorizeTokenError: 40100,
		AuthorizeFailed:         40101,
		InvalidTokenError:       40102,

		//其他错误
		ServerUnknownError: 50001,
	}
	if code, ok := errorCode[e]; ok {
		return code
	} else {
		return 50000
	}
}

func (e ErrorType) Data() interface{} {
	return nil
}

func (e ErrorType) Redirect() string {
	query := make(url.Values)
	query.Add("response_type", "code")
	query.Add("client_id", config.Get().Hduhelp.ClientId)
	query.Add("redirect_uri", "http://localhost:8091/auth/auth")
	query.Add("state", uuid.New())
	redirectUrl := url.URL{
		Scheme:   "https",
		Host:     "api.hduhelp.com",
		Path:     "/oauth/authorize",
		RawQuery: query.Encode(),
	}
	redirectMap := map[ErrorType]string{
		RedirectToHduhelp: redirectUrl.String(),
	}
	if reqUrl, ok := redirectMap[e]; ok {
		return reqUrl
	} else {
		return ""
	}
}

const (
	// 成功
	Success ErrorType = iota

	// 跳转
	RedirectToHduhelp

	// 参数错误
	ParamError
	RandNumTooSmallError
	LackUuidParamError
	LackRandNumParamError
	LackStudentUuidParamError
	LackPictureUuidParamError

	//查询错误
	NoStudentError
	StudentAlreadyExistError
	NoEnoughStudentError
	MysqlInsertError

	//文件相关错误
	FileTooLargeError
	NotImageError
	FileUploadFailedError
	FileUploadToOssFailedError

	//鉴权错误
	LackAuthorizeTokenError
	AuthorizeFailed
	InvalidTokenError

	//服务器错误
	ServerUnknownError
)

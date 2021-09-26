package response

import "errors"

// ErrorType 实现 response.Response 接口
type ErrorType uint

func (e ErrorType) Error() error {
	errorMsg := map[ErrorType]string{
		// 成功
		Success: "OK",

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
		NoEnoughStudentError:     "请求的学生数量大于数据库含有的学生总数",

		//文件相关错误
		FileTooLargeError:          "文件过大，仅支持最大 32M 文件",
		NotImageError:              "仅支持上传图片",
		FileUploadFailedError:      "文件上传错误",
		FileUploadToOssFailedError: "文件上传到阿里云 OSS 失败",
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

		//文件相关错误
		FileTooLargeError:          40061,
		NotImageError:              40062,
		FileUploadFailedError:      40063,
		FileUploadToOssFailedError: 40064,
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
	return ""
}

const (
	// 成功
	Success ErrorType = iota

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

	//文件相关错误
	FileTooLargeError
	NotImageError
	FileUploadFailedError
	FileUploadToOssFailedError
)

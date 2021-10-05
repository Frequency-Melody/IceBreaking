package dto

import (
	"IceBreaking/errs"
	"IceBreaking/response"
)

// JsonResponse Json 返回体结构
type JsonResponse struct {
	Error    int
	Msg      string
	Data     interface{}
	Redirect string
}

// UuidDTO 使 Uuid 单变量 实现 Response 接口
type UuidDTO struct {
	response.BaseResponse `json:"-"`
	Uuid                  string
}

func (d *UuidDTO) Data() interface{} {
	return map[string]string{"uuid": d.Uuid}
}

// CountDto 数量 实现 Response 接口
type CountDto struct {
	response.BaseResponse `json:"-"`
	Count                 string // 这里要用 string，因为 int 的默认值是 0，无法判断数据库查询是否成功
}

func (d *CountDto) Data() interface{} {
	return map[string]string{"count": d.Count}
}

func (d *CountDto) Error() error {
	if d.Count == "" {
		return errs.MysqlQueryError()
	}
	return nil
}

type HidePicDto struct {
	response.BaseResponse `json:"-"`
	HidePic               bool `json:"hide_pic"`
}

func (d *HidePicDto) Data() interface{} {
	return d
}

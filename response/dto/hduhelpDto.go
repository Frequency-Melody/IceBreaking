package dto

import "IceBreaking/response"

type AuthSuccessDto struct {
	response.BaseResponse `json:"-"`
	Token                 string `json:"token"`
	Name                  string `json:"name"`
}

func (d AuthSuccessDto) Data() interface{} {
	return d
}

type HduhelpBaseResponse struct {
	Cache bool   `json:"cache"`
	Error int    `json:"error"`
	Msg   string `json:"msg"`
}

type Code2TokenResponse struct {
	HduhelpBaseResponse
	Data GetTokenSuccessData `json:"data"`
}

type GetTokenSuccessData struct {
	AccessToken        string `json:"access_token"`
	AccessTokenExpire  int64  `json:"access_token_expire"`
	RefreshToken       string `json:"refresh_token"`
	RefreshTokenExpire int64  `json:"refresh_token_expire"`
	StaffId            string `json:"staff_id"`
}

type GetUserInfoResponse struct {
	HduhelpBaseResponse
	Data GetUserInfoData `json:"data"`
}

type GetUserInfoData struct {
	Grade      string `json:"GRADE"`
	StaffId    string `json:"STAFFID"`
	StaffName  string `json:"STAFFNAME"`
	StaffType  string `json:"STAFFTYPE"`
	StaffState string `json:"STAFFSTATE"`
	UnitCode   string `json:"UNITCODE"`
}

type ValidateResponse struct {
	HduhelpBaseResponse
	Data ValidateData `json:"data"`
}

type ValidateData struct {
	AccessToken       string `json:"access_token"`
	AccessTokenExpire int64  `json:"access_token_expire"`
	StaffId           string `json:"staff_id"`
	School            string `json:"school"`
}

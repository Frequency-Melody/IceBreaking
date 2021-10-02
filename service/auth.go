package service

import (
	"IceBreaking/config"
	"IceBreaking/crud"
	"IceBreaking/errs"
	"IceBreaking/log"
	"IceBreaking/model"
	"IceBreaking/response/dto"
	"github.com/parnurzeal/gorequest"
	"net/http"
	"net/url"
	"time"
)

func Code2Token(code, state string) (token string) {
	query := make(url.Values)
	query.Add("client_id", config.Get().Hduhelp.ClientId)
	query.Add("client_secret", config.Get().Hduhelp.ClientSecret)
	query.Add("grant_type", "authorization_code")
	query.Add("code", code)
	query.Add("state", state)
	reqUrl := url.URL{
		Scheme:   "https",
		Host:     "api.hduhelp.com",
		Path:     "/oauth/token",
		RawQuery: query.Encode(),
	}
	res := dto.Code2TokenResponse{}
	_, _, err := gorequest.New().Get(reqUrl.String()).
		Retry(3, time.Second, http.StatusBadRequest, http.StatusInternalServerError).
		EndStruct(&res)
	if err != nil {
		log.Sugar().Info(err)
		return ""
	}
	return res.Data.AccessToken
}

func GetPersonInfo(token string) (staffId, name string, err error) {
	reqUrl := url.URL{
		Scheme: "https",
		Host:   "api.hduhelp.com",
		Path:   "/base/person/info",
	}
	res := dto.GetUserInfoResponse{}
	_, _, err2 := gorequest.New().Get(reqUrl.String()).
		Retry(3, time.Second, http.StatusBadRequest, http.StatusInternalServerError).
		AppendHeader("Authorization", "token "+token).EndStruct(&res)
	if err2 != nil {
		return "", "", errs.InvalidTokenError()
	} else {
		return res.Data.StaffId, res.Data.StaffName, nil
	}
}

// LoginDatabase 将获取到的学生信息映射成自己的数据库的 uuid
// 若数据库无该学生，则新建记录
func LoginDatabase(staffId, name string) (studentUuid string, err error) {
	student := crud.GetStudentByStaffId(staffId)
	if student == nil || student.Uuid == "" {
		studentInsert := model.Student{StaffId: staffId, Name: name, Department: ""}
		if err := crud.AddStudent(&studentInsert); err != nil {
			return "", errs.MysqlQueryError()
		}
		return studentInsert.Uuid, nil
	} else {
		return student.Uuid, nil
	}
}

// Validate 校验 token 合法性，并返回学生在本库中的 uuid
func Validate(token string) (uuid string, err error) {
	reqUrl := url.URL{
		Scheme: "https",
		Host:   "api.hduhelp.com",
		Path:   "/oauth/token/validate",
	}
	res := dto.ValidateResponse{}
	_, _, err2 := gorequest.New().Get(reqUrl.String()).
		Retry(3, time.Second, http.StatusBadRequest, http.StatusInternalServerError).
		AppendHeader("Authorization", "token "+token).EndStruct(&res)
	if err2 != nil {
		log.Sugar().Error(err2)
		return "", errs.InvalidTokenError()
	}
	student := crud.GetStudentByStaffId(res.Data.StaffId)
	return student.Uuid, nil
}

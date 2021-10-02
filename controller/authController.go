package controller

import (
	"IceBreaking/config"
	"IceBreaking/log"
	"IceBreaking/response"
	"IceBreaking/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"strings"
)

// Auth 助手登录后回调此方法，此方法最后会将前端重定向至业务首页，将将 token 和 学生姓名附在 query 中
func Auth(c *gin.Context) {
	code := c.Query("code")
	state := c.Query("state")
	token := service.Code2Token(code, state)
	if token == "" {
		//return response.AuthorizeFailed
		log.Sugar().Error("获取 token 失败")
		c.Abort()
	}
	staffId, name, _ := service.GetPersonInfo(token)
	_, err := service.LoginDatabase(staffId, name)
	if err != nil {
		//return response.MysqlInsertError
		log.Sugar().Error("学号为 %s，名字为 %s 的用户建立失败", staffId, name)
		c.Abort()
	}
	//return dto.AuthSuccessDto{Token: token, Name: name}
	query := make(url.Values)
	query.Add("token", token)
	query.Add("name", name)
	redirectUrl := config.Get().FrontEnd.Home + "?" + query.Encode()
	c.Redirect(http.StatusFound, redirectUrl)
}

// Login 前端执行登录操作，跳转至助手授权页，再由助手跳转到 Auth（即上面那个函数）
func Login(c *gin.Context) {
	c.Redirect(http.StatusFound, response.RedirectToHduhelp.Redirect())
}

// Validate 校验前端的存储的 token 是否有效
func Validate(c *gin.Context) response.Response {
	tokenHeader := strings.TrimSpace(c.GetHeader("Authorization"))
	tokenSlice := strings.Split(tokenHeader, " ")
	if len(tokenSlice) < 2 || strings.TrimSpace(tokenSlice[1]) == "" {
		return response.LackStudentUuidParamError
	} else {
		if _, err := service.Validate(strings.TrimSpace(tokenSlice[1])); err != nil {
			return response.InvalidTokenError
		} else {
			return response.Success
		}
	}
}

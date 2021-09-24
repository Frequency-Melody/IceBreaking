// Package controller 用来检验请求的参数，实际业务由 service 完成
package controller

import (
	"IceBreaking/response"
	"IceBreaking/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func VerifyPictureBelongToStudent(c *gin.Context) {
	var studentId, pictureId int
	var err error
	if studentId, err = strconv.Atoi(c.Query("studentId")); err != nil {
		c.JSON(http.StatusBadRequest, response.MakeErrJson(response.ParamError(err.Error())))
		return
	}
	if pictureId, err = strconv.Atoi(c.Query("pictureId")); err != nil {
		c.JSON(http.StatusBadRequest, response.MakeErrJson(response.ParamError(err.Error())))
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, response.MakeErrJson(response.ParamError(err.Error())))
		return
	}
	c.JSON(http.StatusOK, service.VerifyPictureBelongToStudent(studentId, pictureId))
}

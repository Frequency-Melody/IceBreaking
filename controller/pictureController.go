package controller

import (
	"IceBreaking/response"
	"IceBreaking/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func VerifyPictureBelongToStudent(c *gin.Context) (int, interface{}) {
	var studentId, pictureId int
	var err error
	if studentId, err = strconv.Atoi(c.Query("studentId")); err != nil {
		return http.StatusBadRequest, response.MakeErrJson(response.ParamError(err.Error()))
	}
	if pictureId, err = strconv.Atoi(c.Query("pictureId")); err != nil {
		return http.StatusBadRequest, response.MakeErrJson(response.ParamError(err.Error()))
	}
	if err != nil {
		return http.StatusBadRequest, response.MakeErrJson(response.ParamError(err.Error()))
	}
	return http.StatusOK, service.VerifyPictureBelongToStudent(studentId, pictureId)
}

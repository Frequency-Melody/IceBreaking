package controller

import (
	"IceBreaking/response"
	"IceBreaking/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func VerifyPictureBelongToStudent(c *gin.Context) (int, interface{}) {
	studentUuid := c.Query("studentUuid")
	pictureUuid := c.Query("pictureUuid")
	if studentUuid == "" {
		return http.StatusBadRequest, response.MakeErrJson(response.ParamError("缺少 studentUuid 参数 "))
	}
	if pictureUuid == "" {
		return http.StatusBadRequest, response.MakeErrJson(response.ParamError("缺少 pictureUuid 参数"))
	}
	return http.StatusOK, service.VerifyPictureBelongToStudent(pictureUuid, studentUuid)

}

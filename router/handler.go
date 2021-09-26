package router

import (
	"IceBreaking/response"
	"github.com/gin-gonic/gin"
)

func requestEntry(handler func(c *gin.Context) (r response.Response)) gin.HandlerFunc {
	return func(c *gin.Context) {
		res := handler(c)
		if res.Error() != nil {
			c.JSON(res.Code()/100, response.JsonResponse{Error: res.Code(),
				Msg: res.Error().Error(), Data: res.Data(), Redirect: res.Redirect()})
		} else {
			c.JSON(res.Code()/100, response.JsonResponse{Error: res.Code(),
				Msg: "OK", Data: res.Data(), Redirect: res.Redirect()})
		}
		c.Abort()
	}
}

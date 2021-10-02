package router

import (
	"IceBreaking/response"
	"IceBreaking/response/dto"
	"IceBreaking/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := strings.TrimSpace(c.GetHeader("Authorization"))
		tokenSlice := strings.Split(tokenHeader, " ")
		if len(tokenSlice) < 2 || strings.TrimSpace(tokenSlice[1]) == "" {
			c.Redirect(http.StatusFound, response.RedirectToHduhelp.Redirect())
		} else {
			if uuid, err := service.Validate(strings.TrimSpace(tokenSlice[1])); err != nil {
				c.Abort()
				res := response.InvalidTokenError
				c.JSON(res.Code()/100, dto.JsonResponse{Error: res.Code(),
					Msg: res.Error().Error(), Data: res.Data(), Redirect: res.Redirect()})
			} else {
				c.Set("uuid", uuid)
			}
		}
		c.Next()
	}
}



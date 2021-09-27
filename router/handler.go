package router

import (
	"IceBreaking/response"
	"github.com/gin-gonic/gin"
)

// 每个返回体，无论是 Model、DTO 还是常见错误，都实现了 Response 接口
// 这个函数的作用是，从 Response 接口中提取内容，统一封装数据与错误
func requestEntry(handler func(c *gin.Context) (r response.Response)) gin.HandlerFunc {
	return func(c *gin.Context) {
		res := handler(c)
		if res.Error() != nil {
			if r, ok := res.(response.Trimmer); ok {
				c.JSON(res.Code()/100, response.JsonResponse{Error: res.Code(),
					Msg: res.Error().Error(), Data: r.Trim(), Redirect: res.Redirect()})
			} else {
				c.JSON(res.Code()/100, response.JsonResponse{Error: res.Code(),
					Msg: res.Error().Error(), Data: res.Data(), Redirect: res.Redirect()})
			}
		} else {
			c.JSON(res.Code()/100, response.JsonResponse{Error: res.Code(),
				Msg: "OK", Data: res.Data(), Redirect: res.Redirect()})
		}
		c.Abort()
	}
}

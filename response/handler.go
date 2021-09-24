package response

import "github.com/gin-gonic/gin"

func RequestEntry(handler func(c *gin.Context) (r Response)) gin.HandlerFunc {
	return func(c *gin.Context) {
		res := handler(c)
		if res.Error != nil {
			c.JSON(res.Code()/100, gin.H{"Error": res.Code(),
				"Msg": res.Error().Error(), "data":res.Data(), "redirect": res.Redirect()})
		}else {
			c.JSON(res.Code()/100, gin.H{"Error": res.Code(),
				"Msg": res.Error().Error(), "data":res.Data(), "redirect": res.Redirect()})
		}
		c.Abort()
	}
}

package middleware

import (
	"github.com/gin-gonic/gin"
)

//权限检查中间件
func AuthCheckRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		//data, _ := c.Get("JWT_PAYLOAD")
		//v := data.(jwtauth.MapClaims)
		//e, err := mycasbin.Casbin()
		//tools.HasError(err, "", 500)
		////检查权限
		//res, err := e.Enforce(v["rolekey"], c.Request.URL.Path, c.Request.Method)
		//log.Info(v["rolekey"], c.Request.URL.Path, c.Request.Method)
		//tools.HasError(err, "", 500)
		//
		//if res {
		//	c.Next()
		//} else {
		//	c.JSON(http.StatusOK, gin.H{
		//		"code": 403,
		//		"msg":  fmt.Sprintf("对不起，您没有 <%v-%v> 访问权限，请联系管理员", c.Request.URL.Path, c.Request.Method),
		//	})
		//	c.Abort()
		//	return
		//}
	}
}

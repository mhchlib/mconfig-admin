package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/mhchlib/logger"
)

func InitMiddleware(r *gin.Engine) {
	// 日志处理
	r.Use(LoggerToFile())
	// 自定义错误处理
	r.Use(CustomError)
	// NoCache is a middleware function that appends headers
	r.Use(NoCache)
	// 跨域处理
	r.Use(Options)
	// Secure is a middleware function that appends security
	r.Use(Secure)
	// Set X-Request-Id header
	r.Use(RequestId())

	corsMiddleWare := cors.Default()
	r.Use(corsMiddleWare)

	err := AuthInit()
	if err != nil {
		log.Fatal(err)
	}
	r.Use(Auth)
}

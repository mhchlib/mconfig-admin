package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mhchlib/mconfig-admin/pkg/service"
)

func AddRouters(engine *gin.Engine) {
	group := engine.Group("/api")
	addV1Routers(group)
}

func addV1Routers(group *gin.RouterGroup) {
	v1 := group.Group("v1")

	//app
	app := v1.Group("/app")
	app.POST("/", service.AddApp)
	app.GET("/list", service.ListApp)
	//app.DELETE("/:id", service.DeleteApp)

	//config
	config := v1.Group("/config")
	config.POST("/", service.AddConfig)
	config.GET("/:id", service.GetConfigById)
	config.DELETE("/:id", service.DeleteConfig)

	//filter

	//log
}

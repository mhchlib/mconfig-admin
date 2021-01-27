package router

import (
	"github.com/mhchlib/mconfig-admin/service"
	"github.com/gin-gonic/gin"
)

func RegedistRouters(engine *gin.Engine) {
	group := engine.Group("/api/v1")
	config := group.Group("/config")
	app := group.Group("/app")
	group.POST("/user/login", service.Login)
	config.POST("/", service.AddConfig)
	config.GET("/id/:id", service.GetConfigById)
	config.DELETE("/:id", service.DeleteConfig)
	config.POST("/publish", service.PublishConfig)
	config.POST("/save", service.SaveConfig)
	config.POST("/gray", service.PublishGrayConfig)
	app.POST("/", service.AddApp)
	app.GET("/list", service.GetAppList)
	app.DELETE("/:id", service.DeleteApp)
}

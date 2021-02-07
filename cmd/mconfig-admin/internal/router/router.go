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
	app.POST("", service.AddApp)
	app.GET("/list", service.ListApp)
	app.DELETE("/:id", service.DeleteAPP)
	app.PUT("/:id", service.UpdateApp)

	//env
	env := v1.Group("/env")
	env.POST("/", service.AddEnv)
	env.GET("/list", service.ListEnv)
	env.DELETE("/:id", service.DeleteEnv)
	env.PUT("/filter/:id", service.UpdateEnvFilter)
	env.PUT("/base/:id", service.UpdateEnv)

	//filter
	filter := v1.Group("/filter")
	filter.POST("/", service.AddFilter)
	filter.PUT("/", service.UpdateFilter)
	filter.GET("/:id", service.GetFilter)

	//config
	config := v1.Group("/config")
	config.POST("/", service.AddConfig)
	config.GET("/list", service.ListConfig)
	config.DELETE("/:id", service.DeleteConfig)
	config.PUT("/base/:id", service.UpdateConfig)
	config.PUT("/val/:id", service.UpdateConfigVal)
	config.PUT("/schema/:id", service.UpdateConfigSchema)
	config.PUT("/all/:id", service.UpdateConfigValAndConfig)

	//cluster
	cluster := v1.Group("/cluster")
	cluster.POST("", service.AddCluster)
	cluster.GET("/list", service.ListCluster)
	cluster.GET("/self/:id", service.GetCluster)
	cluster.DELETE("/:id", service.DeleteCluster)
	cluster.PUT("/:id", service.UpdateCluster)

	//deploy
	deploy := v1.Group("/deploy")
	deploy.POST("/config", service.DeployConfig)

	//service
	ser := v1.Group("/service")
	ser.POST("/detail", service.GetServiceDetail)
	ser.POST("/", service.UpdateServiceDetail)
	ser.POST("/delete", service.DeleteServiceItem)

}

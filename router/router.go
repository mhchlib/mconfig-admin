package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mhchlib/mconfig-admin/service"
)

func AddRouters(engine *gin.Engine) {
	api := engine.Group("/api")
	api.GET("/dashboard", service.Dashboard)
	addV1Routers(api)
}

func addV1Routers(group *gin.RouterGroup) {
	v1 := group.Group("v1")
	//dashboard
	dashboard := v1.Group("/dashboard")
	dashboard.GET("/", service.Dashboard)

	//user
	user := v1.Group("/user")
	user.POST("login", service.Login)

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
	filter.GET("/modes", service.GetFilterMode)
	filter.GET("/base/:id", service.GetFilter)

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
	deploy.POST("/config", service.DeployUpdateConfig)
	deploy.POST("/filter", service.DeployUpdateFilter)
	deploy.POST("/filter/delete", service.DeployDeleteFilter)
	deploy.POST("/config/delete", service.DeployDeleteConfig)

	//service
	ser := v1.Group("/service")
	ser.POST("/detail", service.GetServiceDetail)
	ser.POST("/", service.UpdateServiceDetail)
	ser.POST("/delete", service.DeleteServiceItem)

	//tag
	tag := v1.Group("/tag")
	tag.POST("", service.BuildTag)
	tag.GET("/list", service.ListTag)
	tag.GET("/self/:id", service.GetTag)
	tag.DELETE("/:id", service.DeleteTag)
}

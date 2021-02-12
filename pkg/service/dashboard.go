package service

import (
	"github.com/gin-gonic/gin"
	"github.com/mhchlib/mconfig-admin/pkg/model"
)

func Dashboard(c *gin.Context) {
	data := make(map[string]interface{})
	app, err := model.CountApp()
	if err != nil {
		responseDefaultFail(c, nil)
		return
	}
	data["app"] = app
	config, err := model.CountConfig()
	if err != nil {
		responseDefaultFail(c, nil)
		return
	}
	data["config"] = config
	cluster, err := model.CountCluster()
	if err != nil {
		responseDefaultFail(c, nil)
		return
	}
	data["cluster"] = cluster
	responseDefaultSuccess(c, data)
	return
}

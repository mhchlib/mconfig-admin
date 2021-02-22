package service

import (
	"github.com/gin-gonic/gin"
	"github.com/mhchlib/mconfig-admin/model"
	"github.com/mhchlib/mconfig-admin/pkg/tools"
)

func Dashboard(c *gin.Context) {
	data := make(map[string]interface{})
	app, err := model.CountApp()
	if err != nil {
		tools.ResponseDefaultFail(c, nil)
		return
	}
	data["app"] = app
	config, err := model.CountConfig()
	if err != nil {
		tools.ResponseDefaultFail(c, nil)
		return
	}
	data["config"] = config
	cluster, err := model.CountCluster()
	if err != nil {
		tools.ResponseDefaultFail(c, nil)
		return
	}
	data["cluster"] = cluster

	user, err := model.CountUser()
	if err != nil {
		tools.ResponseDefaultFail(c, nil)
		return
	}
	data["user"] = user

	tools.ResponseDefaultSuccess(c, data)
	return
}

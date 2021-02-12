package service

import (
	"context"
	"github.com/gin-gonic/gin"
	log "github.com/mhchlib/logger"
	"github.com/mhchlib/mconfig-admin/pkg/model"
	"github.com/mhchlib/mconfig-api/api/v1/server"
	"github.com/mhchlib/mconfig/pkg/store"
	"github.com/mhchlib/register"
	"github.com/mhchlib/register/reg"
	"google.golang.org/grpc"
	"time"
)

type DeployConfigRequest struct {
	Cluster int `form:"cluster" binding:"required"`
	Tag     int `form:"tag" binding:"required"`
}

func DeployConfig(c *gin.Context) {
	var param DeployConfigRequest
	err := c.Bind(&param)
	if err != nil {
		responseParamError(c)
		return
	}
	clusterId := param.Cluster
	tagId := param.Tag
	tag, err := model.GetTag(tagId)
	if err != nil {
		responseDefaultFail(c, err)
		return
	}
	config, err := model.GetConfig(tag.ConfigId)
	if err != nil {
		responseDefaultFail(c, err)
		return
	}
	if err != nil {
		responseDefaultFail(c, err)
		return
	}
	cluster, err := model.GetCluster(clusterId)
	if err != nil {
		responseDefaultFail(c, err)
		return
	}
	//获取services
	regClient, err := register.InitRegister(func(options *reg.Options) {
		options.RegisterStr = cluster.Register
		options.NameSpace = cluster.Namespace
	})
	if err != nil {
		responseDefaultFail(c, err)
		return
	}
	services, err := regClient.ListAllServices("mconfig-server")
	if services != nil && len(services) == 0 {
		responseDefaultFail(c, "该集群没有线上服务")
		return
	}
	if err != nil {
		responseDefaultFail(c, err)
		return
	}
	app, err := model.GetApp(config.App)
	if err != nil {
		responseDefaultFail(c, err)
		return
	}
	env, err := model.GetEnv(config.Env)
	if err != nil {
		responseDefaultFail(c, err)
		return
	}
	var filter string
	if env.Filter != -1 {
		f, err := model.GetFilter(env.Filter)
		if err != nil {
			responseDefaultFail(c, err)
			return
		}
		filter = f.Filter
	} else {
		filter = ""
	}

	configData := &server.UpdateConfigRequest{
		App:    app.Key,
		Env:    env.Key,
		Config: config.Key,
		Filter: filter,
		Val:    tag.Config,
	}
	//开始部署
	onceShare := false
	for _, service := range services {
		rpcAddress := service.Address
		metadata := service.Metadata
		mode := store.StoreMode(metadata["mode"].(string))
		//一次就好
		if (onceShare == false) && store.MODE_SHARE == mode {
			withTimeout, _ := context.WithTimeout(context.Background(), time.Second*5)
			dial, err := grpc.DialContext(withTimeout, rpcAddress, grpc.WithInsecure(), grpc.WithBlock())
			if err != nil {
				log.Info(err, " addr: ", service)
				continue
			}
			mconfigService := server.NewMConfigClient(dial)
			withTimeout, _ = context.WithTimeout(context.Background(), time.Second*20)
			_, err = mconfigService.UpdateConfig(withTimeout, configData)
			if err != nil {
				log.Error(err)
				responseDefaultFail(c, err)
				return
			}
			onceShare = true
		}

		//每次都要
		if store.MODE_LOCAL == mode {
			withTimeout, _ := context.WithTimeout(context.Background(), time.Second*5)
			dial, err := grpc.DialContext(withTimeout, rpcAddress, grpc.WithInsecure(), grpc.WithBlock())
			if err != nil {
				log.Info(err, " addr: ", service)
				continue
			}
			mconfigService := server.NewMConfigClient(dial)
			withTimeout, _ = context.WithTimeout(context.Background(), time.Second*20)
			_, err = mconfigService.UpdateConfig(withTimeout, configData)
			if err != nil {
				log.Error(err)
				responseDefaultFail(c, err)
				return
			}
		}
		responseDefaultSuccess(c, nil)
		return
	}
}

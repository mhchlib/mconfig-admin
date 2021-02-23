package service

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	log "github.com/mhchlib/logger"
	"github.com/mhchlib/mconfig-admin/model"
	"github.com/mhchlib/mconfig-admin/pkg/tools"
	"github.com/mhchlib/mconfig-api/api/v1/server"
	"github.com/mhchlib/mconfig/core/mconfig"
	"github.com/mhchlib/mconfig/core/store"
	"github.com/mhchlib/register"
	"github.com/mhchlib/register/reg"
	"google.golang.org/grpc"
	"time"
)

type DeployUpdateConfigRequest struct {
	Cluster int `form:"cluster" binding:"required"`
	Tag     int `form:"tag" binding:"required"`
}

func DeployUpdateConfig(c *gin.Context) {
	var param DeployUpdateConfigRequest
	err := c.Bind(&param)
	if err != nil {
		tools.ResponseParamError(c)
		return
	}
	clusterId := param.Cluster
	tagId := param.Tag
	tag, err := model.GetTag(tagId)
	if err != nil {
		tools.ResponseDefaultFail(c, err)
		return
	}
	config, err := model.GetConfig(tag.ConfigId)
	if err != nil {
		tools.ResponseDefaultFail(c, err)
		return
	}
	cluster, err := model.GetCluster(clusterId)
	if err != nil {
		tools.ResponseDefaultFail(c, err)
		return
	}
	//获取services
	regClient, err := register.InitRegister(func(options *reg.Options) {
		options.RegisterStr = cluster.Register
		options.NameSpace = cluster.Namespace
	})
	if err != nil {
		tools.ResponseDefaultFail(c, err)
		return
	}
	services, err := regClient.ListAllServices("mconfig-server")
	if services != nil && len(services) == 0 {
		tools.ResponseDefaultFail(c, "该集群没有线上服务")
		return
	}
	if err != nil {
		tools.ResponseDefaultFail(c, err)
		return
	}
	app, err := model.GetApp(config.App)
	if err != nil {
		tools.ResponseDefaultFail(c, err)
		return
	}
	env, err := model.GetEnv(config.Env)
	if err != nil {
		tools.ResponseDefaultFail(c, err)
		return
	}
	var filter string
	var filterObj *mconfig.StoreVal
	if env.Filter != -1 {
		f, err := model.GetFilter(env.Filter)
		if err != nil {
			tools.ResponseDefaultFail(c, err)
			return
		}
		mode, err := model.GetFilterMode(int(f.Mode))
		if err != nil {
			tools.ResponseDefaultFail(c, err)
			return
		}
		filterObj, err = mconfig.BuildFilterStoreVal(&mconfig.FilterStoreVal{
			Env:    mconfig.ConfigEnv(env.Key),
			Weight: env.Weight,
			Code:   mconfig.FilterVal(f.Filter),
			Mode:   mconfig.FilterMode(mode),
		})
		if err != nil {
			tools.ResponseDefaultFail(c, err)
			return
		}
	} else {
		filterObj, err = mconfig.BuildFilterStoreVal(&mconfig.FilterStoreVal{
			Env:    mconfig.ConfigEnv(env.Key),
			Weight: env.Weight,
			Code:   mconfig.FilterVal(""),
			Mode:   mconfig.FilterMode_lua,
		})
		if err != nil {
			tools.ResponseDefaultFail(c, err)
			return
		}
	}
	filterBytes, err := json.Marshal(filterObj)
	if err != nil {
		tools.ResponseDefaultFail(c, err)
		return
	}
	filter = string(filterBytes)

	configObj, err := mconfig.BuildConfigStoreVal(&mconfig.ConfigStoreVal{
		Key: mconfig.ConfigKey(config.Key),
		Val: mconfig.ConfigVal(tag.Config),
	})
	if err != nil {
		tools.ResponseDefaultFail(c, err)
		return
	}
	configBytes, err := json.Marshal(configObj)
	if err != nil {
		tools.ResponseDefaultFail(c, err)
		return
	}
	configData := &server.UpdateConfigRequest{
		App:    app.Key,
		Env:    env.Key,
		Config: config.Key,
		Filter: filter,
		Val:    string(configBytes),
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
				tools.ResponseDefaultFail(c, err)
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
				tools.ResponseDefaultFail(c, err)
				return
			}
		}
	}
	_ = model.UpdateConfigDeployData(config.Id, tagId)
	tools.ResponseDefaultSuccess(c, nil)
	return
}

type DeployUpdateFilterRequest struct {
	Cluster int `form:"cluster" binding:"required"`
	Env     int `form:"env" binding:"required"`
}

func DeployUpdateFilter(c *gin.Context) {
	var param DeployUpdateFilterRequest
	err := c.Bind(&param)
	if err != nil {
		tools.ResponseParamError(c)
		return
	}
	clusterId := param.Cluster
	envId := param.Env
	env, err := model.GetEnv(envId)
	if err != nil {
		tools.ResponseDefaultFail(c, err)
		return
	}
	cluster, err := model.GetCluster(clusterId)
	if err != nil {
		tools.ResponseDefaultFail(c, err)
		return
	}
	//获取services
	regClient, err := register.InitRegister(func(options *reg.Options) {
		options.RegisterStr = cluster.Register
		options.NameSpace = cluster.Namespace
	})
	if err != nil {
		tools.ResponseDefaultFail(c, err)
		return
	}
	services, err := regClient.ListAllServices("mconfig-server")
	if services != nil && len(services) == 0 {
		tools.ResponseDefaultFail(c, "该集群没有线上服务")
		return
	}
	if err != nil {
		tools.ResponseDefaultFail(c, err)
		return
	}
	var filter string
	if env.Filter != -1 {
		f, err := model.GetFilter(env.Filter)
		if err != nil {
			tools.ResponseDefaultFail(c, err)
			return
		}
		mode, err := model.GetFilterMode(int(f.Mode))
		if err != nil {
			tools.ResponseDefaultFail(c, err)
			return
		}
		filterObj, err := mconfig.BuildFilterStoreVal(&mconfig.FilterStoreVal{
			Env:    mconfig.ConfigEnv(env.Key),
			Weight: env.Weight,
			Code:   mconfig.FilterVal(f.Filter),
			Mode:   mconfig.FilterMode(mode),
		})
		if err != nil {
			tools.ResponseDefaultFail(c, err)
			return
		}
		filterBytes, err := json.Marshal(filterObj)
		if err != nil {
			tools.ResponseDefaultFail(c, err)
			return
		}
		filter = string(filterBytes)
	} else {
		filter = ""
	}
	app, err := model.GetApp(env.App)
	if err != nil {
		tools.ResponseDefaultFail(c, err)
		return
	}
	filterRequest := &server.UpdateFilterRequest{
		App:    app.Key,
		Env:    env.Key,
		Filter: filter,
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
			_, err = mconfigService.UpdateFilter(withTimeout, filterRequest)
			if err != nil {
				log.Error(err)
				tools.ResponseDefaultFail(c, err)
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
			_, err = mconfigService.UpdateFilter(withTimeout, filterRequest)
			if err != nil {
				log.Error(err)
				tools.ResponseDefaultFail(c, err)
				return
			}
		}
	}
	_ = model.UpdateEnvDeployData(env.Id)
	tools.ResponseDefaultSuccess(c, nil)
	return
}


type DeployDeleteFilterRequest struct {
	Cluster int `form:"cluster" binding:"required"`
	Env     int `form:"env" binding:"required"`
}

func DeployDeleteFilter(c *gin.Context) {
	var param DeployDeleteFilterRequest
	err := c.Bind(&param)
	if err != nil {
		tools.ResponseParamError(c)
		return
	}
	clusterId := param.Cluster
	envId := param.Env
	env, err := model.GetEnv(envId)
	if err != nil {
		tools.ResponseDefaultFail(c, err)
		return
	}
	cluster, err := model.GetCluster(clusterId)
	if err != nil {
		tools.ResponseDefaultFail(c, err)
		return
	}
	//获取services
	regClient, err := register.InitRegister(func(options *reg.Options) {
		options.RegisterStr = cluster.Register
		options.NameSpace = cluster.Namespace
	})
	if err != nil {
		tools.ResponseDefaultFail(c, err)
		return
	}
	services, err := regClient.ListAllServices("mconfig-server")
	if services != nil && len(services) == 0 {
		tools.ResponseDefaultFail(c, "该集群没有线上服务")
		return
	}
	if err != nil {
		tools.ResponseDefaultFail(c, err)
		return
	}
	app, err := model.GetApp(env.App)
	if err != nil {
		tools.ResponseDefaultFail(c, err)
		return
	}
	filterRequest := &server.DeletFilterRequest{
		App:    app.Key,
		Env:    env.Key,
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
			_, err = mconfigService.DeletFilter(withTimeout, filterRequest)
			if err != nil {
				log.Error(err)
				tools.ResponseDefaultFail(c, err)
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
			_, err = mconfigService.DeletFilter(withTimeout, filterRequest)
			if err != nil {
				log.Error(err)
				tools.ResponseDefaultFail(c, err)
				return
			}
		}
	}
	tools.ResponseDefaultSuccess(c, nil)
	return
}

type DeployDeleteConfigRequest struct {
	Cluster int `form:"cluster" binding:"required"`
	Config  int `form:"config" binding:"required"`
}

func DeployDeleteConfig(c *gin.Context) {
	var param DeployDeleteConfigRequest
	err := c.Bind(&param)
	if err != nil {
		tools.ResponseParamError(c)
		return
	}
	clusterId := param.Cluster

	cluster, err := model.GetCluster(clusterId)
	if err != nil {
		tools.ResponseDefaultFail(c, err)
		return
	}
	//获取services
	regClient, err := register.InitRegister(func(options *reg.Options) {
		options.RegisterStr = cluster.Register
		options.NameSpace = cluster.Namespace
	})
	if err != nil {
		tools.ResponseDefaultFail(c, err)
		return
	}
	services, err := regClient.ListAllServices("mconfig-server")
	if services != nil && len(services) == 0 {
		tools.ResponseDefaultFail(c, "该集群没有线上服务")
		return
	}
	if err != nil {
		tools.ResponseDefaultFail(c, err)
		return
	}
	config, err := model.GetConfig(param.Config)
	if err != nil {
		tools.ResponseDefaultFail(c, err)
		return
	}
	app, err := model.GetApp(config.App)
	if err != nil {
		tools.ResponseDefaultFail(c, err)
		return
	}
	envId := config.Env
	env, err := model.GetEnv(envId)
	if err != nil {
		tools.ResponseDefaultFail(c, err)
		return
	}
	configRequest := &server.DeletConfigRequest{
		App:    app.Key,
		Env:    env.Key,
		Config: config.Key,
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
			_, err = mconfigService.DeletConfig(withTimeout, configRequest)
			if err != nil {
				log.Error(err)
				tools.ResponseDefaultFail(c, err)
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
			_, err = mconfigService.DeletConfig(withTimeout, configRequest)
			if err != nil {
				log.Error(err)
				tools.ResponseDefaultFail(c, err)
				return
			}
		}
	}
	tools.ResponseDefaultSuccess(c, nil)
	return
}


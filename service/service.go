package service

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/ptypes/empty"
	log "github.com/mhchlib/logger"
	"github.com/mhchlib/mconfig-admin/pkg/tools"
	"github.com/mhchlib/mconfig-api/api/v1/server"
	"github.com/mhchlib/mconfig/core/mconfig"
	"google.golang.org/grpc"
	"time"
)

type GetServiceDetailRequest struct {
	Cluster int    `form:"cluster"`
	Service string `form:"service"  binding:"required"`
}

func GetServiceDetail(c *gin.Context) {
	var param GetServiceDetailRequest
	err := c.Bind(&param)
	if err != nil {
		tools.ResponseParamError(c)
		return
	}
	withTimeout, _ := context.WithTimeout(context.Background(), time.Second*5)
	dial, err := grpc.DialContext(withTimeout, param.Service, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Info(err, " addr: ", param.Service)
		tools.ResponseDefaultFail(c, nil)
		return
	}
	mconfigService := server.NewMConfigClient(dial)
	withTimeout, _ = context.WithTimeout(context.Background(), time.Second*20)
	data, err := mconfigService.GetNodeDetail(withTimeout, &empty.Empty{})
	if err != nil {
		log.Error(err)
		tools.ResponseDefaultFail(c, err)
		return
	}
	_ = dial.Close()
	v := &mconfig.NodeDetail{}
	err = json.Unmarshal(data.Data, v)
	if err != nil {
		log.Error(err)
		tools.ResponseDefaultFail(c, err)
		return
	}
	tools.ResponseDefaultSuccess(c, v)
	return
}

type UpdateServiceDetailRequest struct {
	Service   string `json:"service"`
	App       string `json:"app"`
	Env       string `json:"env"`
	Filter    string `json:"filter"`
	Config    string `json:"config"`
	ConfigVal string `json:"configval"`
	Type      string `json:"type"`
}

func UpdateServiceDetail(c *gin.Context) {
	param := &UpdateServiceDetailRequest{}
	err := c.Bind(&param)
	if err != nil {
		tools.ResponseParamError(c)
		return
	}
	//check param.ConfigVal
	configStoreVal := &mconfig.StoreVal{}
	err = json.Unmarshal([]byte(param.ConfigVal), configStoreVal)
	if err != nil {
		tools.ResponseDefaultFail(c, err)
		return
	}
	configStoreVal.Md5 = mconfig.GetInterfaceMd5(configStoreVal.Data)
	bs, _ := json.Marshal(configStoreVal)
	param.ConfigVal = string(bs)
	configStoreVal = &mconfig.StoreVal{}
	err = json.Unmarshal([]byte(param.Filter), configStoreVal)
	if err != nil {
		tools.ResponseDefaultFail(c, err)
		return
	}
	configStoreVal.Md5 = mconfig.GetInterfaceMd5(configStoreVal.Data)
	bs, _ = json.Marshal(configStoreVal)
	param.Filter = string(bs)

	withTimeout, _ := context.WithTimeout(context.Background(), time.Second*5)
	dial, err := grpc.DialContext(withTimeout, param.Service, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Info(err, " addr: ", param.Service)
		tools.ResponseDefaultFail(c, nil)
		return
	}
	mconfigService := server.NewMConfigClient(dial)
	withTimeout, _ = context.WithTimeout(context.Background(), time.Second*20)
	_, err = mconfigService.UpdateConfig(withTimeout, &server.UpdateConfigRequest{
		App:    param.App,
		Env:    param.Env,
		Config: param.Config,
		Filter: param.Filter,
		Val:    param.ConfigVal,
	})
	_ = dial.Close()
	if err != nil {
		log.Error(err)
		tools.ResponseDefaultFail(c, err)
		return
	}
	tools.ResponseDefaultSuccess(c, nil)
}

type DeleteServiceItemRequest struct {
	Service string `json:"service"`
	App     string `json:"app"`
	Env     string `json:"env"`
	Filter  string `json:"filter"`
	Config  string `json:"config"`
	Type    string `json:"type"`
}

func DeleteServiceItem(c *gin.Context) {
	param := &UpdateServiceDetailRequest{}
	err := c.Bind(&param)
	if err != nil {
		tools.ResponseParamError(c)
		return
	}
	withTimeout, _ := context.WithTimeout(context.Background(), time.Second*5)
	dial, err := grpc.DialContext(withTimeout, param.Service, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Info(err, " addr: ", param.Service)
		tools.ResponseDefaultFail(c, err)
		return
	}
	mconfigService := server.NewMConfigClient(dial)
	withTimeout, _ = context.WithTimeout(context.Background(), time.Second*20)

	if param.Type == "config" {
		_, err = mconfigService.DeletConfig(withTimeout, &server.DeletConfigRequest{
			App:    param.App,
			Env:    param.Env,
			Config: param.Config,
		})
		if err != nil {
			log.Error(err)
			tools.ResponseDefaultFail(c, err)
			return
		}
	}
	if param.Type == "filter" {
		_, err = mconfigService.DeletFilter(withTimeout, &server.DeletFilterRequest{
			App: param.App,
			Env: param.Env,
		})
		if err != nil {
			log.Error(err)
			tools.ResponseDefaultFail(c, err)
			return
		}
	}
	_ = dial.Close()
	tools.ResponseDefaultFail(c, "type 错误")
}

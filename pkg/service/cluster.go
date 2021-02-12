package service

import (
	"github.com/gin-gonic/gin"
	log "github.com/mhchlib/logger"
	"github.com/mhchlib/mconfig-admin/pkg/model"
	"github.com/mhchlib/register"
	"github.com/mhchlib/register/reg"
	"strconv"
)

type AddClusterRequest struct {
	Namespace string `form:"namespace"  binding:"required"`
	Register  string `form:"register"  binding:"required"`
	Desc      string `form:"desc"`
}

func AddCluster(c *gin.Context) {
	var param AddClusterRequest
	err := c.Bind(&param)
	if err != nil {
		responseParamError(c)
		return
	}
	err = model.InsertCluster(param.Namespace, param.Register, param.Desc)
	if err != nil {
		log.Error(err)
		responseDefaultFail(c, nil)
		return
	}
	responseDefaultSuccess(c, nil)
	return
}

type ListClusterRequest struct {
	Filter string `form:"filter"`
	Limit  int    `form:"limit"`
	Offset int    `form:"offset"`
}

type ListClusterResponse struct {
	Id         int    `json:"id"`
	Namespace  string `json:"namespace"`
	Register   string `json:"register"`
	Desc       string `json:"desc"`
	CreateTime int64  `json:"create_time"`
	UpdateTime int64  `json:"update_time"`
}

func ListCluster(c *gin.Context) {
	param := &ListClusterRequest{}
	err := c.Bind(&param)
	if err != nil {
		responseParamError(c)
		return
	}
	if param.Limit == 0 {
		param.Limit = DEFAULT_LIST_LIMIT
	}
	clusters, err := model.ListClusters(param.Filter, param.Limit, param.Offset)
	data := make([]*ListClusterResponse, 0)
	for _, cluster := range clusters {
		data = append(data, &ListClusterResponse{
			Id:         cluster.Id,
			Namespace:  cluster.Namespace,
			Desc:       cluster.Description,
			Register:   cluster.Register,
			CreateTime: cluster.CreateTime,
			UpdateTime: cluster.UpdateTime,
		})
	}
	if err != nil {
		responseDefaultFail(c, nil)
		return
	}
	responseDefaultSuccess(c, data)
	return
}

func DeleteCluster(c *gin.Context) {
	id := c.Param("id")
	log.Info(id)
	atoi, err := strconv.Atoi(id)
	if err != nil {
		responseParamError(c)
		return
	}
	err = model.DeleteCluster(atoi)
	if err != nil {
		responseDefaultFail(c, "删除失败")
		return
	}
	responseDefaultSuccess(c, nil)
}

type UpdateClusterRequest struct {
	Namespace string `form:"namespace"`
	Register  string `form:"register"`
	Desc      string `form:"desc"`
}

func UpdateCluster(c *gin.Context) {
	id := c.Param("id")
	param := &UpdateClusterRequest{}
	err := c.Bind(&param)
	if err != nil {
		responseParamError(c)
		return
	}
	log.Info("update", id, param.Namespace, param.Register, param.Desc)
	atoi, err := strconv.Atoi(id)
	if err != nil {
		responseParamError(c)
		return
	}
	err = model.UpdateCluster(atoi, param.Namespace, param.Register, param.Desc)
	if err != nil {
		responseDefaultFail(c, "更新失败")
		return
	}
	responseDefaultSuccess(c, nil)
	return
}

type GetClusterRepsonse struct {
	*model.Cluster
	Services []*reg.ServiceVal `json:"services"`
}

func GetCluster(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		responseParamError(c)
		return
	}
	response := &GetClusterRepsonse{}
	cluster, err := model.GetCluster(id)
	if err != nil {
		responseDefaultFail(c, "获取失败")
		return
	}
	response.Cluster = cluster

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
	response.Services = services
	responseDefaultSuccess(c, response)
	return
}

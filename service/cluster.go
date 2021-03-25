package service

import (
	"github.com/gin-gonic/gin"
	log "github.com/mhchlib/logger"
	"github.com/mhchlib/mconfig-admin/model"
	"github.com/mhchlib/mconfig-admin/pkg/tools"
	"github.com/mhchlib/register"
	"github.com/mhchlib/register/common"
	"github.com/mhchlib/register/registerOpts"
	"strconv"
)

// AddClusterRequest ...
type AddClusterRequest struct {
	Namespace string `form:"namespace"  binding:"required"`
	Register  string `form:"register"  binding:"required"`
	Desc      string `form:"desc"`
}

// AddCluster ...
func AddCluster(c *gin.Context) {
	var param AddClusterRequest
	err := c.Bind(&param)
	if err != nil {
		tools.ResponseParamError(c)
		return
	}
	err = model.InsertCluster(param.Namespace, param.Register, param.Desc)
	if err != nil {
		log.Error(err)
		tools.ResponseDefaultFail(c, nil)
		return
	}
	tools.ResponseDefaultSuccess(c, nil)
	return
}

// ListClusterRequest ...
type ListClusterRequest struct {
	Filter string `form:"filter"`
	Limit  int    `form:"limit"`
	Offset int    `form:"offset"`
}

// ListClusterResponse ...
type ListClusterResponse struct {
	Id         int    `json:"id"`
	Namespace  string `json:"namespace"`
	Register   string `json:"register"`
	Desc       string `json:"desc"`
	CreateTime int64  `json:"create_time"`
	UpdateTime int64  `json:"update_time"`
}

// ListCluster ...
func ListCluster(c *gin.Context) {
	param := &ListClusterRequest{}
	err := c.Bind(&param)
	if err != nil {
		tools.ResponseParamError(c)
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
		tools.ResponseDefaultFail(c, nil)
		return
	}
	tools.ResponseDefaultSuccess(c, data)
	return
}

// DeleteCluster ...
func DeleteCluster(c *gin.Context) {
	id := c.Param("id")
	log.Info(id)
	atoi, err := strconv.Atoi(id)
	if err != nil {
		tools.ResponseParamError(c)
		return
	}
	err = model.DeleteCluster(atoi)
	if err != nil {
		tools.ResponseDefaultFail(c, "删除失败")
		return
	}
	tools.ResponseDefaultSuccess(c, nil)
}

// UpdateClusterRequest ...
type UpdateClusterRequest struct {
	Namespace string `form:"namespace"`
	Register  string `form:"register"`
	Desc      string `form:"desc"`
}

// UpdateCluster ...
func UpdateCluster(c *gin.Context) {
	id := c.Param("id")
	param := &UpdateClusterRequest{}
	err := c.Bind(&param)
	if err != nil {
		tools.ResponseParamError(c)
		return
	}
	log.Info("update", id, param.Namespace, param.Register, param.Desc)
	atoi, err := strconv.Atoi(id)
	if err != nil {
		tools.ResponseParamError(c)
		return
	}
	err = model.UpdateCluster(atoi, param.Namespace, param.Register, param.Desc)
	if err != nil {
		tools.ResponseDefaultFail(c, "更新失败")
		return
	}
	tools.ResponseDefaultSuccess(c, nil)
	return
}

// GetClusterRepsonse ...
type GetClusterRepsonse struct {
	*model.Cluster
	Services []*common.ServiceVal `json:"services"`
}

// GetCluster ...
func GetCluster(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		tools.ResponseParamError(c)
		return
	}
	Response := &GetClusterRepsonse{}
	cluster, err := model.GetCluster(id)
	if err != nil {
		tools.ResponseDefaultFail(c, "获取失败")
		return
	}
	Response.Cluster = cluster

	//获取services
	regClient, err := register.InitRegister(
		registerOpts.Namespace(cluster.Namespace),
		registerOpts.ResgisterAddress(cluster.Register),
	)
	if err != nil {
		tools.ResponseDefaultFail(c, err)
		return
	}
	services, err := regClient.ListAllServices("mconfig-server")
	Response.Services = services
	tools.ResponseDefaultSuccess(c, Response)
	return
}

package service

import (
	"github.com/gin-gonic/gin"
	log "github.com/mhchlib/logger"
	"github.com/mhchlib/mconfig-admin/model"
	"github.com/mhchlib/mconfig-admin/pkg/common"
	"github.com/mhchlib/mconfig-admin/pkg/tools"
	"strconv"
)

// AddEnvRequest ...
type AddEnvRequest struct {
	App    int    `form:"app"  binding:"required"`
	Filter string `form:"filter"`
	Name   string `form:"name"  binding:"required"`
	Key    string `form:"key"`
	Weight int    `json:"weight"`
	Desc   string `form:"desc"`
}

// AddEnv ...
func AddEnv(c *gin.Context) {
	var param AddEnvRequest
	err := c.Bind(&param)
	if err != nil {
		tools.ResponseParamError(c)
		return
	}
	if param.App == 0 {
		tools.ResponseDefaultFail(c, "app key 不能为空")
		return
	}
	if param.Key == "" {
		param.Key = "env_" + common.GenKey()
	} else {
		unique := model.CheckEnvKeyUnique(param.App, param.Key)
		if !unique {
			tools.ResponseDefaultFail(c, "env key重复")
			return
		}
	}
	err = model.InsertEnv(param.App, param.Name, param.Desc, param.Key, param.Filter, param.Weight)
	if err != nil {
		log.Error(err)
		tools.ResponseDefaultFail(c, nil)
		return
	}
	tools.ResponseDefaultSuccess(c, nil)
	return
}

// ListEnvRequest ...
type ListEnvRequest struct {
	App    int    `form:"app"`
	Filter string `form:"filter"`
	Limit  int    `form:"limit"`
	Offset int    `form:"offset"`
}

// ListEnvResponse ...
type ListEnvResponse struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Desc       string `json:"desc"`
	Key        string `json:"key"`
	Weight     int    `json:"weight"`
	Filter     int    `json:"filter"`
	DeployTime int64  `json:"deploy_time"`
	CreateTime int64  `json:"create_time"`
	UpdateTime int64  `json:"update_time"`
}

// ListEnv ...
func ListEnv(c *gin.Context) {
	param := &ListEnvRequest{}
	err := c.Bind(&param)
	if err != nil {
		tools.ResponseParamError(c)
		return
	}
	if param.App == 0 || param.App == -1 {
		tools.ResponseDefaultFail(c, "app key 无效")
		return
	}
	if param.Limit == 0 {
		param.Limit = DEFAULT_LIST_LIMIT
	}
	envs, err := model.ListEnvs(param.App, param.Filter, param.Limit, param.Offset)
	data := make([]*ListEnvResponse, 0)
	for _, env := range envs {
		data = append(data, &ListEnvResponse{
			Id:         env.Id,
			Name:       env.Name,
			Desc:       env.Description,
			Key:        env.Key,
			Weight:     env.Weight,
			Filter:     env.Filter,
			DeployTime: env.DeployTime,
			CreateTime: env.CreateTime,
			UpdateTime: env.UpdateTime,
		})
	}
	if err != nil {
		tools.ResponseDefaultFail(c, err)
		return
	}
	tools.ResponseDefaultSuccess(c, data)
	return
}

// DeleteEnv ...
func DeleteEnv(c *gin.Context) {
	id := c.Param("id")
	log.Info(id)
	atoi, err := strconv.Atoi(id)
	if err != nil {
		tools.ResponseParamError(c)
		return
	}
	err = model.DeleteEnv(atoi)
	if err != nil {
		tools.ResponseDefaultFail(c, "删除失败")
		return
	}
	tools.ResponseDefaultSuccess(c, nil)
}

// UpdateEnvRequest ...
type UpdateEnvRequest struct {
	Name   string `form:"name"`
	Desc   string `form:"desc"`
	Weight int    `json:"weight"`
}

// UpdateEnv ...
func UpdateEnv(c *gin.Context) {
	idStr := c.Param("id")
	param := &UpdateEnvRequest{}
	err := c.Bind(&param)
	if err != nil {
		tools.ResponseParamError(c)
		return
	}
	log.Info("update", idStr, param.Name, param.Desc, param.Weight)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		tools.ResponseParamError(c)
		return
	}
	err = model.UpdateEnv(id, param.Name, param.Desc, param.Weight)
	if err != nil {
		tools.ResponseDefaultFail(c, "更新失败")
		return
	}
	tools.ResponseDefaultSuccess(c, nil)
	return
}

// UpdateEnvFilterRequest ...
type UpdateEnvFilterRequest struct {
	Filter int `form:"filter"`
}

// UpdateEnvFilter ...
func UpdateEnvFilter(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		tools.ResponseParamError(c)
		return
	}
	param := &UpdateEnvFilterRequest{}
	err = c.Bind(&param)
	if err != nil {
		tools.ResponseParamError(c)
		return
	}
	err = model.UpdateEnvFilter(id, param.Filter)
	if err != nil {
		tools.ResponseDefaultFail(c, "更新环境信息失败")
		return
	}
	tools.ResponseDefaultSuccess(c, nil)
	return
}

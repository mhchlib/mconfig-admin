package service

import (
	"github.com/gin-gonic/gin"
	log "github.com/mhchlib/logger"
	"github.com/mhchlib/mconfig-admin/pkg/common"
	"github.com/mhchlib/mconfig-admin/pkg/model"
	"github.com/mhchlib/mconfig-admin/pkg/tools"
	"strconv"
)

type AddConfigRequest struct {
	App  int    `form:"app" binding:"required"`
	Env  int    `form:"env" binding:"required"`
	Name string `form:"name" binding:"required"`
	Key  string `form:"key"`
	Desc string `form:"desc"`
}

func AddConfig(c *gin.Context) {
	var param AddConfigRequest
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
		param.Key = "config_" + common.GenKey()
	} else {
		unique := model.CheckConfigKeyUnique(param.App, param.Env, param.Key)
		if !unique {
			tools.ResponseDefaultFail(c, "config key重复")
			return
		}
	}
	err = model.InsertConfig(param.App, param.Env, param.Name, param.Desc, param.Key)
	if err != nil {
		log.Error(err)
		tools.ResponseDefaultFail(c, nil)
		return
	}
	tools.ResponseDefaultSuccess(c, nil)
	return
}

type ListConfigRequest struct {
	App    int    `form:"app"`
	Env    int    `form:"env"`
	Filter string `form:"filter"`
	Limit  int    `form:"limit"`
	Offset int    `form:"offset"`
}

type ListConfigResponse struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Desc       string `json:"desc"`
	Key        string `json:"key"`
	Config     string `json:"config"`
	Schema     string `json:"schema"`
	CreateTime int64  `json:"create_time"`
	UpdateTime int64  `json:"update_time"`
}

func ListConfig(c *gin.Context) {
	param := &ListConfigRequest{}
	err := c.Bind(&param)
	if err != nil {
		tools.ResponseParamError(c)
		return
	}
	if param.App == 0 || param.App == -1 {
		tools.ResponseDefaultFail(c, "app id 无效")
		return
	}
	if param.Env == 0 || param.Env == -1 {
		tools.ResponseDefaultFail(c, "env id 无效")
		return
	}
	if param.Limit == 0 {
		param.Limit = DEFAULT_LIST_LIMIT
	}
	Configs, err := model.ListConfigs(param.App, param.Env, param.Filter, param.Limit, param.Offset)
	data := make([]*ListConfigResponse, 0)
	for _, Config := range Configs {
		data = append(data, &ListConfigResponse{
			Id:         Config.Id,
			Name:       Config.Name,
			Desc:       Config.Description,
			Key:        Config.Key,
			Config:     Config.Val,
			Schema:     Config.Schema,
			CreateTime: Config.CreateTime,
			UpdateTime: Config.UpdateTime,
		})
	}
	if err != nil {
		tools.ResponseDefaultFail(c, err)
		return
	}
	tools.ResponseDefaultSuccess(c, data)
	return
}

func DeleteConfig(c *gin.Context) {
	id := c.Param("id")
	log.Info(id)
	atoi, err := strconv.Atoi(id)
	if err != nil {
		tools.ResponseParamError(c)
		return
	}
	err = model.DeleteConfig(atoi)
	if err != nil {
		tools.ResponseDefaultFail(c, "删除失败")
		return
	}
	tools.ResponseDefaultSuccess(c, nil)
}

type UpdateConfigRequest struct {
	Name string `form:"name"`
	Desc string `form:"desc"`
}

func UpdateConfig(c *gin.Context) {
	idStr := c.Param("id")
	param := &UpdateConfigRequest{}
	err := c.Bind(&param)
	if err != nil {
		tools.ResponseParamError(c)
		return
	}
	log.Info("update", idStr, param.Name, param.Desc)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		tools.ResponseParamError(c)
		return
	}
	err = model.UpdateConfig(id, param.Name, param.Desc)
	if err != nil {
		tools.ResponseDefaultFail(c, "更新失败")
		return
	}
	tools.ResponseDefaultSuccess(c, nil)
	return
}

type UpdateConfigValRequest struct {
	Val string `form:"config"`
}

func UpdateConfigVal(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		tools.ResponseParamError(c)
		return
	}
	param := &UpdateConfigValRequest{}
	err = c.Bind(&param)
	if err != nil {
		tools.ResponseParamError(c)
		return
	}
	err = model.UpdateConfigVal(id, param.Val)
	if err != nil {
		tools.ResponseDefaultFail(c, "更新信息失败")
		return
	}
	tools.ResponseDefaultSuccess(c, nil)
	return
}

type UpdateConfigSchemaRequest struct {
	Val string `form:"config"`
}

func UpdateConfigSchema(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		tools.ResponseParamError(c)
		return
	}
	param := &UpdateConfigValRequest{}
	err = c.Bind(&param)
	if err != nil {
		tools.ResponseParamError(c)
		return
	}
	err = model.UpdateConfigSchema(id, param.Val)
	if err != nil {
		tools.ResponseDefaultFail(c, "更新信息失败")
		return
	}
	tools.ResponseDefaultSuccess(c, nil)
	return
}

type UpdateConfigValAndConfigRequest struct {
	Config string `form:"config"`
	Schema string `form:"schema"`
}

func UpdateConfigValAndConfig(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		tools.ResponseParamError(c)
		return
	}
	param := &UpdateConfigValAndConfigRequest{}
	err = c.Bind(&param)
	if err != nil {
		tools.ResponseParamError(c)
		return
	}
	err = model.UpdateConfigValAndConfig(id, param.Config, param.Schema)
	if err != nil {
		tools.ResponseDefaultFail(c, "更新信息失败")
		return
	}
	tools.ResponseDefaultSuccess(c, nil)
	return
}

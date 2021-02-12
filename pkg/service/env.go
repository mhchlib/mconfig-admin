package service

import (
	"github.com/gin-gonic/gin"
	log "github.com/mhchlib/logger"
	"github.com/mhchlib/mconfig-admin/pkg/common"
	"github.com/mhchlib/mconfig-admin/pkg/model"
	"strconv"
)

type AddEnvRequest struct {
	App    int    `form:"app"  binding:"required"`
	Filter string `form:"filter"`
	Name   string `form:"name"  binding:"required"`
	Key    string `form:"key"`
	Weight int    `json:"weight"`
	Desc   string `form:"desc"`
}

func AddEnv(c *gin.Context) {
	var param AddEnvRequest
	err := c.Bind(&param)
	if err != nil {
		responseParamError(c)
		return
	}
	if param.App == 0 {
		responseDefaultFail(c, "app key 不能为空")
		return
	}
	if param.Key == "" {
		param.Key = "env_" + common.GenKey()
	} else {
		unique := model.CheckEnvKeyUnique(param.App, param.Key)
		if !unique {
			responseDefaultFail(c, "env key重复")
			return
		}
	}
	err = model.InsertEnv(param.App, param.Name, param.Desc, param.Key, param.Filter, param.Weight)
	if err != nil {
		log.Error(err)
		responseDefaultFail(c, nil)
		return
	}
	responseDefaultSuccess(c, nil)
	return
}

type ListEnvRequest struct {
	App    int    `form:"app"`
	Filter string `form:"filter"`
	Limit  int    `form:"limit"`
	Offset int    `form:"offset"`
}

type ListEnvResponse struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Desc       string `json:"desc"`
	Key        string `json:"key"`
	Weight     int    `json:"weight"`
	Filter     int    `json:"filter"`
	CreateTime int64  `json:"create_time"`
	UpdateTime int64  `json:"update_time"`
}

func ListEnv(c *gin.Context) {
	param := &ListEnvRequest{}
	err := c.Bind(&param)
	if err != nil {
		responseParamError(c)
		return
	}
	if param.App == 0 || param.App == -1 {
		responseDefaultFail(c, "app key 无效")
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
			CreateTime: env.CreateTime,
			UpdateTime: env.UpdateTime,
		})
	}
	if err != nil {
		responseDefaultFail(c, err)
		return
	}
	responseDefaultSuccess(c, data)
	return
}

func DeleteEnv(c *gin.Context) {
	id := c.Param("id")
	log.Info(id)
	atoi, err := strconv.Atoi(id)
	if err != nil {
		responseParamError(c)
		return
	}
	err = model.DeleteEnv(atoi)
	if err != nil {
		responseDefaultFail(c, "删除失败")
		return
	}
	responseDefaultSuccess(c, nil)
}

type UpdateEnvRequest struct {
	Name   string `form:"name"`
	Desc   string `form:"desc"`
	Weight int    `json:"weight"`
}

func UpdateEnv(c *gin.Context) {
	idStr := c.Param("id")
	param := &UpdateEnvRequest{}
	err := c.Bind(&param)
	if err != nil {
		responseParamError(c)
		return
	}
	log.Info("update", idStr, param.Name, param.Desc, param.Weight)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		responseParamError(c)
		return
	}
	err = model.UpdateEnv(id, param.Name, param.Desc, param.Weight)
	if err != nil {
		responseDefaultFail(c, "更新失败")
		return
	}
	responseDefaultSuccess(c, nil)
	return
}

type UpdateEnvFilterRequest struct {
	Filter int `form:"filter"`
}

func UpdateEnvFilter(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		responseParamError(c)
		return
	}
	param := &UpdateEnvFilterRequest{}
	err = c.Bind(&param)
	if err != nil {
		responseParamError(c)
		return
	}
	err = model.UpdateEnvFilter(id, param.Filter)
	if err != nil {
		responseDefaultFail(c, "更新环境信息失败")
		return
	}
	responseDefaultSuccess(c, nil)
	return
}

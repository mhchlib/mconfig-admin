package service

import (
	"github.com/gin-gonic/gin"
	log "github.com/mhchlib/logger"
	"github.com/mhchlib/mconfig-admin/model"
	"github.com/mhchlib/mconfig-admin/pkg/common"
	"github.com/mhchlib/mconfig-admin/pkg/tools"
	"strconv"
)

// AddAppRequest ...
type AddAppRequest struct {
	Name string `form:"name"  binding:"required"`
	Key  string `form:"key"`
	Desc string `form:"desc"`
}

// PREFIX_APP_KEY ...
const PREFIX_APP_KEY = "app_"

// AddApp ...
func AddApp(c *gin.Context) {
	var param AddAppRequest
	err := c.Bind(&param)
	if err != nil {
		tools.ResponseParamError(c)
		return
	}
	if param.Key == "" {
		param.Key = PREFIX_APP_KEY + common.GenKey()
	} else {
		unique := model.CheckAppKeyUnique(param.Key)
		if !unique {
			tools.ResponseDefaultFail(c, "app key重复")
			return
		}
	}
	err = model.InsertApp(param.Name, param.Desc, param.Key)
	if err != nil {
		log.Error(err)
		tools.ResponseDefaultFail(c, nil)
		return
	}
	tools.ResponseDefaultSuccess(c, nil)
	return
}

// ListAppRequest ...
type ListAppRequest struct {
	Filter string `form:"filter"`
	Limit  int    `form:"limit"`
	Offset int    `form:"offset"`
}

// ListAppResponse ...
type ListAppResponse struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Desc       string `json:"desc"`
	Key        string `json:"key"`
	CreateTime int64  `json:"create_time"`
	UpdateTime int64  `json:"update_time"`
}

// ListApp ...
func ListApp(c *gin.Context) {
	var param ListAppRequest
	err := c.Bind(&param)
	if err != nil {
		tools.ResponseParamError(c)
		return
	}
	if param.Limit == 0 {
		param.Limit = DEFAULT_LIST_LIMIT
	}
	apps, err := model.ListApps(param.Filter, param.Limit, param.Offset)
	data := make([]*ListAppResponse, 0)
	for _, app := range apps {
		data = append(data, &ListAppResponse{
			Id:         app.Id,
			Name:       app.Name,
			Desc:       app.Description,
			Key:        app.Key,
			CreateTime: app.CreateTime,
			UpdateTime: app.UpdateTime,
		})
	}
	if err != nil {
		tools.ResponseDefaultFail(c, nil)
		return
	}
	tools.ResponseDefaultSuccess(c, data)
	return
}

// DeleteAPP ...
func DeleteAPP(c *gin.Context) {
	id := c.Param("id")
	log.Info(id)
	atoi, err := strconv.Atoi(id)
	if err != nil {
		tools.ResponseParamError(c)
		return
	}
	err = model.DeleteApp(atoi)
	if err != nil {
		tools.ResponseDefaultFail(c, "删除失败")
		return
	}
	tools.ResponseDefaultSuccess(c, nil)
}

// UpdateAppRequest ...
type UpdateAppRequest struct {
	Name string `form:"name"`
	Desc string `form:"desc"`
}

// UpdateApp ...
func UpdateApp(c *gin.Context) {
	id := c.Param("id")
	param := &UpdateAppRequest{}
	err := c.Bind(&param)
	if err != nil {
		tools.ResponseParamError(c)
		return
	}
	log.Info("update", id, param.Name, param.Desc)
	atoi, err := strconv.Atoi(id)
	if err != nil {
		tools.ResponseParamError(c)
		return
	}
	err = model.UpdateApp(atoi, param.Name, param.Desc)
	if err != nil {
		tools.ResponseDefaultFail(c, "更新失败")
		return
	}
	tools.ResponseDefaultSuccess(c, nil)
	return
}

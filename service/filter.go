package service

import (
	"github.com/gin-gonic/gin"
	log "github.com/mhchlib/logger"
	"github.com/mhchlib/mconfig-admin/model"
	"github.com/mhchlib/mconfig-admin/pkg/tools"
	"strconv"
)

// AddFilterRequest ...
type AddFilterRequest struct {
	Filter string `form:"filter"`
	Mode   int    `json:"mode"`
}

// AddFilter ...
func AddFilter(c *gin.Context) {
	param := &AddFilterRequest{}
	err := c.Bind(&param)
	if err != nil {
		tools.ResponseParamError(c)
		return
	}
	id, err := model.InsertFilter(model.Mode_FILTER(param.Mode), param.Filter)
	if err != nil {
		log.Error(err)
		tools.ResponseDefaultFail(c, "创建失败")
		return
	}
	tools.ResponseDefaultSuccess(c, id)
	return
}

// UpdateFilterRequest ...
type UpdateFilterRequest struct {
	Id     int    `form:"id"`
	Filter string `form:"filter"`
	Mode   int    `json:"mode"`
}

// UpdateFilter ...
func UpdateFilter(c *gin.Context) {
	param := &UpdateFilterRequest{}
	err := c.Bind(&param)
	if err != nil {
		tools.ResponseParamError(c)
		return
	}
	err = model.UpdateFilter(param.Id, model.Mode_FILTER(param.Mode), param.Filter)
	if err != nil {
		tools.ResponseDefaultFail(c, "更新FILTER信息失败")
		return
	}
	tools.ResponseDefaultSuccess(c, nil)
	return
}

// GetFilter ...
func GetFilter(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		tools.ResponseParamError(c)
		return
	}
	item, err := model.GetFilter(id)
	if err != nil {
		tools.ResponseDefaultFail(c, "获取filter信息失败")
		return
	}
	tools.ResponseDefaultSuccess(c, item)
	return
}

// GetFilterMode ...
func GetFilterMode(c *gin.Context) {
	modes, err := model.GetFilterModes()
	if err != nil {
		tools.ResponseDefaultFail(c, "获取filter mode信息失败")
		return
	}
	tools.ResponseDefaultSuccess(c, modes)
	return
}

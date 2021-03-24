package service

import (
	"github.com/gin-gonic/gin"
	log "github.com/mhchlib/logger"
	"github.com/mhchlib/mconfig-admin/model"
	"github.com/mhchlib/mconfig-admin/pkg/tools"
	"strconv"
)

// BuildTagRequest ...
type BuildTagRequest struct {
	Tag    string `form:"tag"  binding:"required"`
	Desc   string `form:"desc"`
	Config int    `form:"config"  binding:"required"`
}

// BuildTag ...
func BuildTag(c *gin.Context) {
	var param BuildTagRequest
	err := c.Bind(&param)
	if err != nil {
		tools.ResponseParamError(c)
		return
	}
	config, err := model.GetConfig(param.Config)
	if err != nil {
		log.Error(err)
		tools.ResponseDefaultFail(c, nil)
		return
	}
	err = model.InsertTag(param.Tag, param.Desc, param.Config, config.Val, config.Schema)
	if err != nil {
		log.Error(err)
		tools.ResponseDefaultFail(c, nil)
		return
	}
	tools.ResponseDefaultSuccess(c, nil)
	return
}

// ListTagRequest ...
type ListTagRequest struct {
	Config int    `form:"config"  binding:"required"`
	Filter string `form:"filter"`
	Limit  int    `form:"limit"`
	Offset int    `form:"offset"`
}

// ListTagResponse ...
type ListTagResponse struct {
	Id         int    `json:"id"`
	Tag        string `json:"tag"`
	Desc       string `json:"desc"`
	CreateTime int64  `json:"create_time"`
	UpdateTime int64  `json:"update_time"`
}

// ListTag ...
func ListTag(c *gin.Context) {
	var param ListTagRequest
	err := c.Bind(&param)
	if err != nil {
		tools.ResponseParamError(c)
		return
	}
	if param.Limit == 0 {
		param.Limit = DEFAULT_LIST_LIMIT
	}
	tags, err := model.ListTags(param.Config, param.Filter, param.Limit, param.Offset)
	data := make([]*ListTagResponse, 0)
	for _, tag := range tags {
		data = append(data, &ListTagResponse{
			Id:         tag.Id,
			Tag:        tag.Tag,
			Desc:       tag.Description,
			CreateTime: tag.CreateTime,
			UpdateTime: tag.UpdateTime,
		})
	}
	if err != nil {
		tools.ResponseDefaultFail(c, nil)
		return
	}
	tools.ResponseDefaultSuccess(c, data)
	return
}

// DeleteTag ...
func DeleteTag(c *gin.Context) {
	id := c.Param("id")
	log.Info(id)
	atoi, err := strconv.Atoi(id)
	if err != nil {
		tools.ResponseParamError(c)
		return
	}
	err = model.DeleteTag(atoi)
	if err != nil {
		tools.ResponseDefaultFail(c, "删除失败")
		return
	}
	tools.ResponseDefaultSuccess(c, nil)
}

// GetTag ...
func GetTag(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		tools.ResponseParamError(c)
		return
	}
	item, err := model.GetTag(id)
	if err != nil {
		tools.ResponseDefaultFail(c, "获取tag信息失败")
		return
	}
	tools.ResponseDefaultSuccess(c, item)
	return
}

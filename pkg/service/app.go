package service

import (
	"github.com/gin-gonic/gin"
	log "github.com/mhchlib/logger"
	"github.com/mhchlib/mconfig-admin/pkg/common"
	"github.com/mhchlib/mconfig-admin/pkg/model"
)

type AddAppRequest struct {
	Name string `form:"name"`
	Key  string `form:"key"`
	Desc string `form:"desc"`
}

func AddApp(c *gin.Context) {
	param := &AddAppRequest{}
	err := c.Bind(&param)
	if err != nil {
		responseParamError(c)
		return
	}
	if param.Key == "" {
		param.Key = "app_" + common.GenKey()
	} else {
		unique := model.CheckKeyUnique(param.Key)
		if !unique {
			responseDefaultFail(c, "app key重复")
			return
		}
	}
	err = model.InsertApp(param.Name, param.Desc, param.Key)
	if err != nil {
		log.Error(err)
		responseDefaultFail(c, nil)
		return
	}
	responseDefaultSuccess(c, nil)
	return
}

type ListAppRequest struct {
	Filter string `form:"filter"`
	Limit  int    `form:"limit"`
	Offset int    `form:"offset"`
}

type ListAppResponse struct {
	Name       string
	Desc       string
	Key        string
	CreateTime int64
	UpdateTime int64
}

func ListApp(c *gin.Context) {
	param := &ListAppRequest{}
	err := c.Bind(&param)
	if err != nil {
		responseParamError(c)
		return
	}
	if param.Limit == 0 {
		param.Limit = DEFAULT_LIST_LIMIT
	}
	apps, err := model.ListApps(param.Filter, param.Limit, param.Offset)
	data := make([]*ListAppResponse, 0)
	for _, app := range apps {
		data = append(data, &ListAppResponse{
			Name:       app.Name,
			Desc:       app.Description,
			Key:        app.Key,
			CreateTime: app.CreateTime,
			UpdateTime: app.UpdateTime,
		})
	}
	if err != nil {
		responseDefaultFail(c, nil)
		return
	}
	responseDefaultSuccess(c, data)
	return
}

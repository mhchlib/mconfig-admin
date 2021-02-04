package service

import (
	"github.com/gin-gonic/gin"
	"github.com/mhchlib/mconfig-admin/pkg/common"
)

func response(c *gin.Context, code interface{}, msg interface{}, data interface{}) {
	c.JSON(200, &gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

func responseParamError(c *gin.Context) {
	response(c, common.CODE_ERROR_PARAM, common.GetResponseMsg(common.CODE_ERROR_PARAM), nil)
}

func responseDefaultFail(c *gin.Context, msg interface{}) {
	if msg != nil {
		response(c, common.CODE_FAIL_REQUEST, common.GetResponseMsg(common.CODE_FAIL_REQUEST), msg)
		return
	}
	response(c, common.CODE_FAIL_REQUEST, msg, nil)
}

func responseDefaultSuccess(c *gin.Context, data interface{}) {
	response(c, common.CODE_SUCCESS_REQUEST, common.GetResponseMsg(common.CODE_SUCCESS_REQUEST), data)
}

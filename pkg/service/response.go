package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseCode int
type ResponseMsg string

var resp map[ResponseCode]ResponseMsg

const (
	CODE_ERROR_PARAM     ResponseCode = 1000
	CODE_FAIL_REQUEST    ResponseCode = 1001
	CODE_SUCCESS_REQUEST ResponseCode = 1002
)

func init() {
	resp = map[ResponseCode]ResponseMsg{
		CODE_ERROR_PARAM:     "参数错误",
		CODE_FAIL_REQUEST:    "请求失败",
		CODE_SUCCESS_REQUEST: "请求成功",
	}
}

func GetResponseMsg(code ResponseCode) ResponseMsg {
	return resp[code]
}

func GetResponse(code ResponseCode) map[string]interface{} {
	m := make(map[string]interface{})
	m["code"] = code
	m["msg"] = resp[code]
	return m
}

func response(c *gin.Context, code interface{}, msg interface{}, data interface{}) {
	c.JSON(http.StatusOK, &gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

func responseParamErrorWithMsg(c *gin.Context, msg string) {
	response(c, CODE_ERROR_PARAM, string(GetResponseMsg(CODE_ERROR_PARAM))+", "+msg, nil)
}

func responseParamError(c *gin.Context) {
	response(c, CODE_ERROR_PARAM, GetResponseMsg(CODE_ERROR_PARAM), nil)
}

func responseDefaultFail(c *gin.Context, msg interface{}) {
	if msg != nil {
		response(c, CODE_FAIL_REQUEST, msg, nil)
		return
	}
	response(c, CODE_FAIL_REQUEST, GetResponseMsg(CODE_FAIL_REQUEST), nil)
}

func responseDefaultSuccess(c *gin.Context, data interface{}) {
	response(c, CODE_SUCCESS_REQUEST, GetResponseMsg(CODE_SUCCESS_REQUEST), data)
}

func responseCustom(c *gin.Context, code ResponseCode, msg interface{}, data interface{}) {
	if msg != nil {
		response(c, code, msg, data)
	}
	response(c, code, GetResponseMsg(code), data)
}

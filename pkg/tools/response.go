package tools

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// ResponseCode ...
type ResponseCode int

// ResponseMsg ...
type ResponseMsg string

var resp map[ResponseCode]ResponseMsg

const (
	// CODE_ERROR_PARAM ...
	CODE_ERROR_PARAM ResponseCode = 1000
	// CODE_FAIL_REQUEST ...
	CODE_FAIL_REQUEST ResponseCode = 1001
	// CODE_SUCCESS_REQUEST ...
	CODE_SUCCESS_REQUEST ResponseCode = 1002
	// CODE_FAIL_TOKEN_INVALID ...
	CODE_FAIL_TOKEN_INVALID ResponseCode = 1003
)

func init() {
	resp = map[ResponseCode]ResponseMsg{
		CODE_ERROR_PARAM:        "参数错误",
		CODE_FAIL_REQUEST:       "请求失败",
		CODE_SUCCESS_REQUEST:    "请求成功",
		CODE_FAIL_TOKEN_INVALID: "token无效",
	}
}

// GetResponseMsg ...
func GetResponseMsg(code ResponseCode) ResponseMsg {
	return resp[code]
}

// GetResponse ...
func GetResponse(code ResponseCode) map[string]interface{} {
	m := make(map[string]interface{})
	m["code"] = code
	m["msg"] = resp[code]
	return m
}

// Response ...
func Response(c *gin.Context, code interface{}, msg interface{}, data interface{}) {
	c.JSON(http.StatusOK, &gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

// Response403 ...
func Response403(c *gin.Context) {
	c.JSON(http.StatusForbidden, &gin.H{
		"msg":  "没有权限",
		"code": 403,
	})
}

// ResponseTokenInvalid ...
func ResponseTokenInvalid(c *gin.Context) {
	c.JSON(http.StatusOK, &gin.H{
		"msg":  resp[CODE_FAIL_TOKEN_INVALID],
		"code": CODE_FAIL_TOKEN_INVALID,
	})
}

// ResponseParamErrorWithMsg ...
func ResponseParamErrorWithMsg(c *gin.Context, msg string) {
	Response(c, CODE_ERROR_PARAM, string(GetResponseMsg(CODE_ERROR_PARAM))+", "+msg, nil)
}

// ResponseParamError ...
func ResponseParamError(c *gin.Context) {
	Response(c, CODE_ERROR_PARAM, GetResponseMsg(CODE_ERROR_PARAM), nil)
}

// ResponseDefaultFail ...
func ResponseDefaultFail(c *gin.Context, msg interface{}) {
	switch msg.(type) {
	case error:
		msg = msg.(error).Error()
	default:
	}
	if msg != nil {
		Response(c, CODE_FAIL_REQUEST, msg, nil)
		return
	}
	Response(c, CODE_FAIL_REQUEST, GetResponseMsg(CODE_FAIL_REQUEST), nil)
}

// ResponseDefaultSuccess ...
func ResponseDefaultSuccess(c *gin.Context, data interface{}) {
	Response(c, CODE_SUCCESS_REQUEST, GetResponseMsg(CODE_SUCCESS_REQUEST), data)
}

// ResponseCustom ...
func ResponseCustom(c *gin.Context, code ResponseCode, msg interface{}, data interface{}) {
	if msg != nil {
		Response(c, code, msg, data)
	}
	Response(c, code, GetResponseMsg(code), data)
}

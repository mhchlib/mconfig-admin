package common

type ResponseCode int
type ResponseMsg string

var response map[ResponseCode]ResponseMsg

const (
	CODE_ERROR_PARAM     ResponseCode = 1000
	CODE_FAIL_REQUEST    ResponseCode = 1001
	CODE_SUCCESS_REQUEST ResponseCode = 1002
)

func init() {
	response = map[ResponseCode]ResponseMsg{
		CODE_ERROR_PARAM:     "参数错误",
		CODE_FAIL_REQUEST:    "请求失败",
		CODE_SUCCESS_REQUEST: "请求成功",
	}
}

func GetResponseMsg(code ResponseCode) ResponseMsg {
	return response[code]
}

func GetResponse(code ResponseCode) map[string]interface{} {
	m := make(map[string]interface{})
	m["code"] = code
	m["msg"] = response[code]
	return m
}

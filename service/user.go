package service

import (
	"github.com/gin-gonic/gin"
	log "github.com/mhchlib/logger"
	"github.com/mhchlib/mconfig-admin/model"
	"github.com/mhchlib/mconfig-admin/pkg/middleware"
	"github.com/mhchlib/mconfig-admin/pkg/tools"
	"strconv"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	param := &LoginRequest{}
	err := c.Bind(&param)
	if err != nil {
		tools.ResponseParamError(c)
		return
	}
	user, err := model.GetUserByName(param.Username)
	if err != nil {
		log.Error(err)
		tools.ResponseDefaultFail(c, "登陆失败")
		return
	}
	passwd := tools.Md5Crypt(param.Password, user.Salt)
	if passwd != user.Password {
		tools.ResponseDefaultFail(c, "登陆失败")
		return
	}
	//login success
	jwtMiddleware := middleware.GetJwtMiddleware()
	token, _, err := jwtMiddleware.TokenGenerator(map[string]interface{}{
		"userId": strconv.Itoa(user.Id),
	})
	if err != nil {
		log.Error(err)
		tools.ResponseDefaultFail(c, "登陆失败")
		return
	}
	tools.ResponseDefaultSuccess(c,
		map[string]string{
			"token": token,
		},
	)
}

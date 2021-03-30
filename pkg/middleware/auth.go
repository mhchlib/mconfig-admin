package middleware

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v2"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/mhchlib/logger"
	jwtUtil "github.com/mhchlib/mconfig-admin/pkg/jwt_util"
	"github.com/mhchlib/mconfig-admin/pkg/tools"
	"github.com/micro/go-micro/v2/logger"
	"github.com/spf13/viper"
	"time"
)

var jwtMiddleware *jwtUtil.GinJWTMiddleware

var enforcer casbin.IEnforcer

// AuthInit ...
func AuthInit() error {
	middleware, err := jwtUtil.New(&jwtUtil.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte(viper.GetString("jwt.key")),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		PayloadFunc: PayloadFunc,
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		TimeFunc:    time.Now,
	})
	jwtMiddleware = middleware
	if err != nil {
		return err
	}
	//init casbin
	Apter, err := gormadapter.NewAdapter("mysql", viper.GetString("db.url"), true)
	if err != nil {
		return err
	}
	m := model.Model{}
	err = m.LoadModelFromText(CASBIN)
	if err != nil {
		return err
	}
	e, err := casbin.NewSyncedEnforcer(m, Apter)
	if err != nil {
		return err
	}
	e.EnableLog(true)
	err = e.LoadPolicy()
	if err != nil {
		logger.Infof("casbin rbac_model or policy init error, message: %v \r\n", err.Error())
		return err
	}
	enforcer = e
	return nil
}

// PayloadFunc ...
func PayloadFunc(data interface{}) jwtUtil.MapClaims {
	if v, ok := data.(map[string]interface{}); ok {
		return jwtUtil.MapClaims(jwt.MapClaims{
			"userId":  v["userId"],
			"roleKey": v["roleKey"],
		})
	}
	return jwtUtil.MapClaims{}
}

// Auth ...
func Auth(c *gin.Context) {
	//check white list
	whitelist := []string{"/api/v1/user/login"}
	for _, whiteUri := range whitelist {
		if whiteUri == c.FullPath() {
			c.Next()
			return
		}
	}
	claims, err := jwtMiddleware.GetClaimsFromJWT(c)
	if err != nil {
		log.Error(err)
		tools.ResponseTokenInvalid(c)
		c.Abort()
		return
	}
	log.Info(claims)
	c.Set("userId", claims["userId"])
	enforcer.AddRoleForUser(fmt.Sprintf("%v", claims["userId"]), "root")
	//check policy
	enforce, err := enforcer.Enforce(claims["userId"], c.FullPath(), c.Request.Method)
	if err != nil {
		log.Error(err)
		tools.Response403(c)
		c.Abort()
		return
	}
	if enforce == false {
		tools.Response403(c)
		c.Abort()
		return
	}
	//	generator, _, _ := jwtMiddleware.TokenGenerator(map[string]interface{}{"userId":"100","roleKey":"admin"})
	//	log.Info(generator)
	c.Next()
}

// GetJwtMiddleware ...
func GetJwtMiddleware() *jwtUtil.GinJWTMiddleware {
	return jwtMiddleware
}

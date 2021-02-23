module github.com/mhchlib/mconfig-admin

go 1.14

require (
	github.com/casbin/casbin/v2 v2.2.2
	github.com/casbin/gorm-adapter/v2 v2.1.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.6.3
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/protobuf v1.4.0
	github.com/jinzhu/gorm v1.9.16
	github.com/mhchlib/logger v0.0.1
	github.com/mhchlib/mconfig v0.0.0-00010101000000-000000000000
	github.com/mhchlib/mconfig-api v0.0.0-20201023050446-420de20374cc
	github.com/mhchlib/register v0.0.0-20201023050446-420de20374cc
	github.com/micro/go-micro/v2 v2.9.1
	github.com/satori/go.uuid v1.2.0
	github.com/spf13/viper v1.7.1
	google.golang.org/grpc v1.26.0
)

replace github.com/mhchlib/register v0.0.0-20201023050446-420de20374cc => ../register

replace github.com/mhchlib/mconfig => ../mconfig

replace github.com/mhchlib/logger v0.0.0-20201023050446-420de20374cc => ../logger

replace github.com/mhchlib/mconfig-api v0.0.0-20201023050446-420de20374cc => ../mconfig-api


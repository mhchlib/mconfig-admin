module github.com/mhchlib/mconfig-admin

go 1.14

require (
	github.com/casbin/casbin/v2 v2.2.2
	github.com/casbin/gorm-adapter/v2 v2.1.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.6.3
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/protobuf v1.4.2
	github.com/jinzhu/gorm v1.9.16
	github.com/mhchlib/logger v0.0.3-0.20210324103410-ddf65533f989
	github.com/mhchlib/mconfig v1.0.1-0.20210825174004-de924ea574e1
	github.com/mhchlib/mconfig-api v0.0.2-0.20210326111514-9f081bbd6da2
	github.com/mhchlib/mregister v0.0.2-0.20210825173657-52159fdd45bf
	github.com/micro/go-micro/v2 v2.9.1
	github.com/satori/go.uuid v1.2.0
	github.com/spf13/viper v1.7.1
	google.golang.org/grpc v1.27.0
)

replace google.golang.org/grpc v1.27.0 => google.golang.org/grpc v1.26.0

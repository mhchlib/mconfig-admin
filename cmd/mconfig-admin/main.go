package main

import (
	"github.com/gin-gonic/gin"
	log "github.com/mhchlib/logger"
	"github.com/mhchlib/mconfig-admin/cmd/mconfig-admin/internal/router"
	_ "github.com/mhchlib/mconfig-admin/pkg"
	"github.com/mhchlib/mconfig-admin/pkg/middleware"
)

func main() {
	engine := gin.New()
	middleware.InitMiddleware(engine)
	router.AddRouters(engine)
	err := engine.Run(":8087")
	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"github.com/mhchlib/mconfig-admin/router"
	log "github.com/mhchlib/logger"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	router.RegedistRouters(engine)
	err := engine.Run(":8087")
	if err != nil {
		log.Fatal(err)
	}
}

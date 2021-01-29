package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/mhchlib/logger"
	"github.com/spf13/viper"
)

var db *gorm.DB

func Init() {
	url := viper.GetString("db.url")
	mdb, err := gorm.Open("mysql", url)
	if err != nil {
		log.Fatal(err)
	}
	mdb.LogMode(true)
	mdb.SingularTable(true)
	mdb.DB().SetMaxIdleConns(10)
	mdb.DB().SetMaxOpenConns(100)
	db = mdb
}

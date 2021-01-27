package service

import (
	"github.com/gin-gonic/gin"
	"log"
)

type PostConfigRequest struct {
	AppId      string   `json:"appid" binding:"required"`
	ConfigName string   `json:"configName" binding:"required"`
	Desc       string   `json:"desc" binding:"required"`
	Tags       []string `json:"tags" binding:"required"`
	Config     string   `json:"config" binding:"required"`
	Schema     string   `json:"schema" binding:"required"`
}

type Config struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}

func AddConfig(c *gin.Context) {
	param := &PostConfigRequest{}
	err := c.Bind(&param)
	if err != nil {
		log.Println(err)
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "param error",
		})
		return
	}
	c.JSON(200, gin.H{
		"code":      200,
		"projectid": 10000,
	})
	log.Println(param)
}

func GetConfigById(c *gin.Context) {

	c.JSON(200, gin.H{
		"code": 200,
		"config": gin.H{
			"Published": gin.H{
				"config": "{\"name\":\"test111\"}",
				"schema": "{\"type\": \"object\",\"properties\":{\"a\":{\"type\":\"string\"}}}",
			},
			"Unpublished": gin.H{
				"config": "{\"name\":\"test222\"}",
				"schema": "{\"type\": \"object\",\"properties\":{\"a\":{\"type\":\"string\"}}}",
			},
			"Gray-Release": gin.H{
				"config": "{\"name\":\"test333\"}",
				"schema": "{\"type\": \"object\",\"properties\":{\"a\":{\"type\":\"string\"}}}",
			},
		},
	})
}

func DeleteConfig(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 200,
		"data": gin.H{},
	})
}

func PublishConfig(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 200,
		"data": gin.H{},
	})
}

func SaveConfig(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 200,
		"data": gin.H{},
	})
}

func PublishGrayConfig(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 200,
		"data": gin.H{},
	})
}

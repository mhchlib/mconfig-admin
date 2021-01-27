package service

import (
	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	param := &LoginRequest{}
	c.Bind(&param)
	if param.Username == "admin" && param.Password == "111111" {
		c.JSON(200, gin.H{
			"token": "pong",
		})
	}
}

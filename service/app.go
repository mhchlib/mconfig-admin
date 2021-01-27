package service

import (
	"github.com/gin-gonic/gin"
	"log"
)

type PostAppRequest struct {
	App  string `json:"app" binding:"required"`
	Desc string `json:"desc" binding:"required"`
}

type App struct {
	Id      int       `json:"id"`
	Name    string    `json:"name"`
	Desc    string    `json:"desc"`
	Configs []*Config `json:"configs"`
}

func AddApp(c *gin.Context) {
	param := &PostAppRequest{}
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
		"code":  200,
		"appid": 10000,
	})
	log.Println(param)
}

func GetAppList(c *gin.Context) {
	/*
	 app: [{
	        id: '12987122',
	        name: '好滋好味鸡蛋仔',
	        desc: '荷兰优质淡奶，奶香浓而不腻',
	        configs: [{
	            id: 1111,
	            name: "111",
	            desc: "333"
	          },
	          {
	            id: 2222,
	            name: "111",
	            desc: "333"
	          },
	          {
	            id: 3333,
	            name: "111",
	            desc: "333"
	          }
	        ]

	      }
	*/
	list := []*App{}
	configs01 := []*Config{}
	configs01 = append(configs01, &Config{
		Id:   100,
		Name: "config-100",
		Desc: "config-100-desc",
	})
	configs01 = append(configs01, &Config{
		Id:   101,
		Name: "config-101",
		Desc: "config-101-desc",
	}, &Config{
		Id:   102,
		Name: "config-102",
		Desc: "config-102-desc",
	}, &Config{
		Id:   103,
		Name: "config-103",
		Desc: "config-103-desc",
	}, &Config{
		Id:   104,
		Name: "config-104",
		Desc: "config-104-desc",
	})

	list = append(list, &App{
		Id:      1000,
		Name:    "app-1000",
		Desc:    "app-1000-desc",
		Configs: configs01,
	}, &App{
		Id:      2000,
		Name:    "app-2000",
		Desc:    "app-2000-desc",
		Configs: configs01,
	}, &App{
		Id:      3000,
		Name:    "app-3000",
		Desc:    "app-3000-desc",
		Configs: configs01,
	})
	c.JSON(200, gin.H{
		"code": 200,
		"data": gin.H{
			"apps": list,
		},
	})
}

func DeleteApp(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 200,
		"data": gin.H{},
	})
}

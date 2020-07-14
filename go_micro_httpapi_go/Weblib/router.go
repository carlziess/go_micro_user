package Weblib

import (
	"github.com/gin-gonic/gin"
	"http_api/Services"
)

func NewGinRoute(prodservice Services.ProdService)*gin.Engine{
	ginRoute := gin.Default()
	ware := InitMiddleWare(prodservice)
	//全局使用该中件间
	ginRoute.Use(ware)
	prodGroup := ginRoute.Group("/prod")
	{
		prodGroup.POST("/list", ProdListHandle())
		prodGroup.GET("/detail/",ProdDetailHandle())
	}
	return ginRoute
}
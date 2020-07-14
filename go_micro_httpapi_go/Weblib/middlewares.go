package Weblib

import (
	"github.com/gin-gonic/gin"
	"http_api/Services"
)

//gin框架中的中间件机制；在执行目标函数之前的操作；
//在micro中的装饰器wrapper同样可以实现

func InitMiddleWare(prodservice Services.ProdService)gin.HandlerFunc{
	return func(ginctx *gin.Context) {
		ginctx.Keys=make(map[string]interface{})
		ginctx.Keys["prodService"]=prodservice
		ginctx.Next()
	}
}


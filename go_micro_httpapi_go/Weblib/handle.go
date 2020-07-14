package Weblib

import (
	"context"
	"github.com/gin-gonic/gin"
	"http_api/Services"
	"net/http"

)

func ProdListHandle()gin.HandlerFunc{
	return func(ginContext *gin.Context) {
		//从中间件中取值
		prodService:= ginContext.Keys["prodService"].(Services.ProdService)
		var req Services.ProdReq
		err := ginContext.Bind(&req)
		rsp:=new(Services.ProdRsp)
		if err != nil {
			ginContext.JSON(http.StatusInternalServerError,gin.H{"status":err.Error()})
		}else {
			rsp, _ = prodService.GetProdList(context.Background(), &req)
			ginContext.JSON(http.StatusOK,gin.H{
				"data":rsp.Data,
			})
		}
		///*
		//	封装在wrapper中，这边只专注handle
		//*/
		////熔断器hystric的使用
		////1.设置规则config
		//configCommand:=hystrix.CommandConfig{
		//	Timeout: 1000, //表示允许访问的时间最大是1s
		//}
		////2.绑定config
		//hystrix.ConfigureCommand("getProdTime",configCommand)
		//rsp:=new(Services.ProdRsp)
		////3.Do执行，传入目标函数  fallback是降级方法，一般不要去调取过于复杂的任务，可以返回默认数据，不要出现err
		//err = hystrix.Do("getProdTime", func() error {
		//	rsp, err = prodService.GetProdList(context.Background(), &req)
		//	return err
		//}, nil)
		//
		//if err != nil {
		//	ginContext.JSON(http.StatusInternalServerError,gin.H{"status":err.Error()})
		//}else {
		//	ginContext.JSON(http.StatusOK,gin.H{
		//		"data":rsp.Data,
		//	})
		//}

	}
}

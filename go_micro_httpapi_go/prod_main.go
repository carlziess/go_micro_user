package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
	"http_api/Services"
	"http_api/Weblib"
	"http_api/Wrappers"
)

//http-api被客户端访问，之后通过rpc从consul中调取注册的grpc服务（业务逻辑）

func main(){
	//调取consul中注册的gprc服务
	myService := micro.NewService(
		micro.Name("prodServiceClient"),
		//log装饰器
		micro.WrapClient(Wrappers.NewLogWrapper),
		//prod访问时间控制熔断
		micro.WrapClient(Wrappers.NewProdWrapper),

	)
	prodService := Services.NewProdService("prods", myService.Client())

	newRegistry := consul.NewRegistry(
		registry.Addrs("192.168.241.129:8500"),
	)
	httpService := web.NewService(
		web.Name("httpApiService"),
		web.Registry(newRegistry),
		web.Handler(Weblib.NewGinRoute(prodService)),
		web.Address(":8010"),
	)
	httpService.Init()
	httpService.Run()

}






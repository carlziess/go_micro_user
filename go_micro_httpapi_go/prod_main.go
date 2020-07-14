package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
	"http_api/Services"
	"http_api/Weblib"
)

//http-api被客户端访问，之后通过rpc从consul中调取注册的grpc服务（业务逻辑）


//装饰器wrapper的使用
type logWrapper struct {
	client.Client
}
func (l *logWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	fmt.Println("调用接口，日志输出等操作")
	return l.Client.Call(ctx, req, rsp)
}
func NewLogWrapper(c client.Client) client.Client {
	return &logWrapper{c}
}


func main(){
	//调取consul中注册的gprc服务
	myService := micro.NewService(
		micro.Name("prodServiceClient"),
		//装饰器wrapper
		micro.WrapClient(NewLogWrapper),
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






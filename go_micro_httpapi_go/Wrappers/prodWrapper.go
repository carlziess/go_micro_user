package Wrappers

import (
	"context"
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
	"http_api/Services"
)

//装饰器wrapper的使用

//日志
type logWrapper struct {
	client.Client
}
func (l *logWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	fmt.Println("调用接口，日志输出等操作")
	md, _ := metadata.FromContext(ctx)
	fmt.Printf("[Log Wrapper] ctx: %v service: %s method: %s\n", md, req.Service(), req.Endpoint())
	return l.Client.Call(ctx, req, rsp)
}

func NewLogWrapper(c client.Client) client.Client {
	return &logWrapper{c}
}

//添加访问限制config
type ProdWrapper struct {
	client.Client
}

func defaultFunc(rsp interface{})error{
	var result []*Services.ProdModel
	arg:=Services.ProdModel{
		ProdID: int32(500),
		ProdName: "降级function",
	}
	result=append(result,&arg)
	rsp.(*Services.ProdRsp).Data=result
	return nil
}

func (p *ProdWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error{
	cmdName:=req.Service()+"_"+req.Endpoint()
	configA:= hystrix.CommandConfig{
		Timeout: 5000,
	}
	hystrix.ConfigureCommand(cmdName,configA)
	return hystrix.Do(cmdName,func() error{
		return p.Client.Call(ctx,req,rsp)
	}, func(err error) error {
		return defaultFunc(rsp)
	})
}

func NewProdWrapper(c client.Client) client.Client {
	return &ProdWrapper{c}
}





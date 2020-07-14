package main

import (
	"Gomicro_grpc/ServiceImpl"
	Services "Gomicro_grpc/service"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	"log"
)


func main(){
	newRegistry := consul.NewRegistry(
		registry.Addrs("192.168.241.129:8500"),
		)
	newService := micro.NewService(
		//micro.Address(":8002"),
		micro.Name("prods"),
		micro.Registry(newRegistry),
	)
	newService.Init()
	err := Services.RegisterProdServiceHandler(newService.Server(), new(ServiceImpl.ProdService))
	if err != nil {
		log.Fatal(err)
	}
	newService.Run()
}



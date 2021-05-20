package main

import (
	"log"

	"github.com/asim/go-micro/plugins/registry/consul/v3"
	httpServer "github.com/asim/go-micro/plugins/server/http/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
	"github.com/asim/go-micro/v3/server"
	"github.com/userserver/routers"
)

const (
	SERVER_NAME = "user-server" // server name
)

var consulReg registry.Registry
func init() {
	consulReg = consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"),
	)
	
}

func main() {
	srv := httpServer.NewServer(
		server.Name(SERVER_NAME),
		server.Address(":18000"),
	)

	ginRouter := routers.InitRouters()

	hd := srv.NewHandler(ginRouter)
	if err := srv.Handle(hd); err != nil {
		log.Fatalln(err)
	}

	service := micro.NewService(
		micro.Server(srv),
		micro.Registry(consulReg),
	)
	service.Init()
	service.Run()
}
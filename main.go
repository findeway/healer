package main

import (
	"healer/handler"
	"healer/subscriber"

	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"

	healer "healer/proto/healer"

	"github.com/micro/go-micro/v2/registry/etcd"
)

func main() {
	// New Service
	registre := etcd.NewRegistry()
	service := micro.NewService(
		micro.Registry(registre),
		micro.Name("go.micro.service.healer"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	healer.RegisterHealerHandler(service.Server(), new(handler.Healer))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.service.healer", service.Server(), new(subscriber.Healer))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

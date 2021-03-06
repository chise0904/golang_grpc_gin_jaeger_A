package main

import (
	"golang_grpc_gin_jaeger_A/grpcServer"
	// "golang_grpc_gin_jaeger_A/httpServer"
)

func main() {

	ch := make(chan struct{})

	go grpcServer.Run()
	// go httpServer.Run()

	<-ch
}

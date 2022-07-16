package main

import (
	pb "bigwork/gen/go/pet/v1"
	"bigwork/server"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

func main() {
	gs := grpc.NewServer()
	pb.RegisterPetStoreServiceServer(gs,server.NewHelloServer())
	//开启服务
	listen, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Panic(err)
	}

	log.Info("pet server start success!!")
	if err := gs.Serve(listen); err != nil {
		log.Panic(err)
	}
}
package main

import (
	"fmt"
	"log"
	"net"

	"github.com/Prasang-money/go-grpc-product-svc/pkg/config"
	"github.com/Prasang-money/go-grpc-product-svc/pkg/db"
	"github.com/Prasang-money/go-grpc-product-svc/pkg/pb"
	"github.com/Prasang-money/go-grpc-product-svc/pkg/services"
	"google.golang.org/grpc"
)

func main() {
	config, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("failed to load config ", err)
	}

	h := db.Init(config.DBUrl)

	lis, err := net.Listen("tcp", config.Port)

	if err != nil {
		log.Fatalln("failed to listen ", err)
	}

	fmt.Println("Product Service listening on ", config.Port)

	s := services.Server{
		H: h,
	}

	grpcServer := grpc.NewServer()
	pb.RegisterProductServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("failed to serve: ", err)
	}
}

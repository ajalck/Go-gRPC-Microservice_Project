package main

import (
	"log"
	"net"

	"github.com/ajalck/Go-gRPC-Microservice_Project/order_management/pkg/client"
	"github.com/ajalck/Go-gRPC-Microservice_Project/order_management/pkg/config"
	"github.com/ajalck/Go-gRPC-Microservice_Project/order_management/pkg/db"
	"github.com/ajalck/Go-gRPC-Microservice_Project/order_management/pkg/pb"
	"github.com/ajalck/Go-gRPC-Microservice_Project/order_management/pkg/services"
	"google.golang.org/grpc"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed to load config")
	}
	DB := db.InitDB(config.DBurl)
	if err = db.SyncDB(DB.DB); err != nil {
		log.Fatalln("Failed to sync :", err)
	}

	lis, err := net.Listen("tcp", config.Port)
	if err != nil {
		log.Fatalln("failed to listen :", err)
	}
	productServiceClient := client.InitProductServiceClient(config.Product_srv_url)
	grpcServer := grpc.NewServer()

	server := &services.OrderServer{
		DB:  DB.DB,
		PSC: productServiceClient,
	}
	pb.RegisterOrderServiceServer(grpcServer, server)

	log.Println("Server started listening at :", config.Port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}

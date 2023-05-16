package main

import (
	"log"
	"net"

	"github.com/ajalck/Go-gRPC-Microservice_Project/product_management/pkg/config"
	"github.com/ajalck/Go-gRPC-Microservice_Project/product_management/pkg/db"
	"github.com/ajalck/Go-gRPC-Microservice_Project/product_management/pkg/pb"
	"github.com/ajalck/Go-gRPC-Microservice_Project/product_management/pkg/services"
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

	grpcServer := grpc.NewServer()

	server := &services.ProductServer{
		DB: DB.DB,
	}
	pb.RegisterProductServiceServer(grpcServer, server)

	log.Println("Server started listening at :", config.Port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}

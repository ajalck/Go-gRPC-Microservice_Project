package main

import (
	"Go-gRPC-Microservice_Project/auth_management/pkg/Pb"
	"Go-gRPC-Microservice_Project/auth_management/pkg/config"
	"Go-gRPC-Microservice_Project/auth_management/pkg/db"
	"Go-gRPC-Microservice_Project/auth_management/pkg/service"
	"log"
	"net"

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
	server := &service.AuthServer{}
	Pb.RegisterAuthServiceServer(grpcServer, server)

	log.Println("Server started listening at :", config.Port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}

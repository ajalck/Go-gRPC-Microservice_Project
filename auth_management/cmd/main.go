package main

import (
	"go-grpc-microservice-auth_management/pkg/config"
	"go-grpc-microservice-auth_management/pkg/db"
	"go-grpc-microservice-auth_management/pkg/pb"
	"go-grpc-microservice-auth_management/pkg/service"
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
	pb.RegisterAuthServiceServer(grpcServer, server)

	log.Println("Server started listening at :", config.Port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}

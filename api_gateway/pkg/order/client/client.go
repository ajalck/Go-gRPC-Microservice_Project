package client

import (
	"log"

	"github.com/ajalck/Go-gRPC-Microservice_Project/api_gateway/pkg/config"
	"github.com/ajalck/Go-gRPC-Microservice_Project/order_management/pkg/pb"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.OrderServiceClient
}

func InitServiceClient(c *config.Config) pb.OrderServiceClient {
	cc, err := grpc.Dial(c.OrderSrvUrl, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal("Failed to initiate ProductServiceClient")
	}
	return pb.NewOrderServiceClient(cc)
}

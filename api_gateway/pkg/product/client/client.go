package client

import (
	"log"

	"github.com/ajalck/Go-gRPC-Microservice_Project/api_gateway/pkg/config"
	"github.com/ajalck/Go-gRPC-Microservice_Project/product_management/pkg/pb"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.ProductServiceClient
}

func InitServiceClient(c *config.Config) pb.ProductServiceClient {
	cc, err := grpc.Dial(c.ProductSrvUrl, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal("Failed to initiate ProductServiceClient")
	}
	return pb.NewProductServiceClient(cc)
}

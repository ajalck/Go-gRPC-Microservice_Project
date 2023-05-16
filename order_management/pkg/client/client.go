package client

import (
	"log"

	"github.com/ajalck/Go-gRPC-Microservice_Project/product_management/pkg/pb"
	"google.golang.org/grpc"
)

type PdtServiceClient struct {
	Client pb.ProductServiceClient
}

func InitProductServiceClient(url string) *PdtServiceClient {
	cc, err := grpc.Dial(url, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal("Could not connect to product service client")
		grpc.WithReturnConnectionError()
	}
	return &PdtServiceClient{
		Client: pb.NewProductServiceClient(cc),
	}
}

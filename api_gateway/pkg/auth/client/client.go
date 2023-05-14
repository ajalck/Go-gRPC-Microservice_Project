package client

import (
	"Go-gRPC-Microservice_Project/api_gateway/pkg/config"

	"github.com/ajalck/Go-gRPC-Microservice_Project/auth_management/pkg/Pb"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client Pb.AuthServiceClient
}

func InitServiceClient(c *config.Config, logger hclog.Logger) Pb.AuthServiceClient {

	// "Go-gRPC-Microservice_Project/auth_management/pkg/Pb"
	logger.Info("API Gateway : Initiated AuthService Client")

	cc, err := grpc.Dial(c.AuthSrvUrl, grpc.WithInsecure, grpc.WithBlock)
	if err != nil {
		logger.Error("Could not connect AuthService :", err)
	}

	return Pb.NewAuthServiceClient(cc)
}

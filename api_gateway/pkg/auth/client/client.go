package client

import (
	"github.com/ajalck/Go-gRPC-Microservice_Project/api_gateway/pkg/config"

	"github.com/ajalck/Go-gRPC-Microservice_Project/auth_management/pkg/pb"

	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

func InitServiceClient(c *config.Config, logger hclog.Logger) pb.AuthServiceClient {

	logger.Info("API Gateway : Initiated AuthService Client")

	cc, err := grpc.Dial(c.AuthSrvUrl, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		logger.Error("Could not connect AuthService :", err)
	}

	return pb.NewAuthServiceClient(cc)
}

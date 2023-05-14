package routes

import (
	"fmt"
	"go-grpc-microservice-api_gateway/pkg/auth/client"
	"go-grpc-microservice-api_gateway/pkg/config"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-hclog"
)

func AuthRoutes(r *gin.Engine, c *config.Config, logger hclog.Logger) {
	svc := &client.ServiceClient{
		Client: client.InitServiceClient(c, logger),
	}

	user := r.Group("/user")
	{
		user.POST("/register")
		user.POST("/login")
	}
	fmt.Println(svc)
}

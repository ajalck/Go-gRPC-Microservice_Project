package routes

import (
	"fmt"
	"Go-gRPC-Microservice_Project/api_gateway/pkg/auth/client"
	h "Go-gRPC-Microservice_Project/api_gateway/pkg/auth/handler"
	"Go-gRPC-Microservice_Project/api_gateway/pkg/config"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-hclog"
)

func AuthRoutes(r *gin.Engine, c *config.Config, logger hclog.Logger) {
	svc := &client.ServiceClient{
		Client: client.InitServiceClient(c, logger),
	}
	authHandler := &h.AuthHandler{C: svc.Client}
	user := r.Group("/user")
	{
		user.POST("/register", authHandler.Register)
		user.POST("/login")
	}
	fmt.Println(svc)
}

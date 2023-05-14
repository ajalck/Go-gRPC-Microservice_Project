package main

import (
	auth "go-grpc-microservice-api_gateway/pkg/auth/routes"
	"go-grpc-microservice-api_gateway/pkg/config"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-hclog"
)

func main() {
	logger := hclog.Default()
	logger.Info("Starting api gateway")

	c, err := config.LoadConfig()
	if err != nil {
		logger.Error("Failed to load config", err)
	}

	r := gin.Default()

	auth.AuthRoutes(r, &c, logger)

	r.Run(c.Port)

}

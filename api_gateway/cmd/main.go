package main

import (
	"go-grpc-microservice-api_gateway/pkg/config"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-hclog"
)

func main() {
	log := hclog.Default()
	log.Info("Starting api gateway")

	c, err := config.LoadConfig()
	if err != nil {
		log.Error("Failed to load config", err)
	}

	r := gin.Default()

	r.Run(c.Port)

}

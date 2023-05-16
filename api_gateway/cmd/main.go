package main

import (
	auth "github.com/ajalck/Go-gRPC-Microservice_Project/api_gateway/pkg/auth/routes"
	"github.com/ajalck/Go-gRPC-Microservice_Project/api_gateway/pkg/config"
	order "github.com/ajalck/Go-gRPC-Microservice_Project/api_gateway/pkg/order/routes"
	product "github.com/ajalck/Go-gRPC-Microservice_Project/api_gateway/pkg/product/routes"
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

	authSvcC := auth.AuthRoutes(r, &c, logger)
	product.ProductRoutes(r, &c, authSvcC)
	order.OrderRoutes(r, &c, authSvcC)

	r.Run(c.Port)

}

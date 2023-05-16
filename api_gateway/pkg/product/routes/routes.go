package routes

import (
	authClient "github.com/ajalck/Go-gRPC-Microservice_Project/api_gateway/pkg/auth/client"
	"github.com/ajalck/Go-gRPC-Microservice_Project/api_gateway/pkg/config"
	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.Engine, c *config.Config, authSvcC *authClient.ServiceClient){
	
}

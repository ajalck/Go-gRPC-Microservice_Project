package routes

import (
	authClient "github.com/ajalck/Go-gRPC-Microservice_Project/api_gateway/pkg/auth/client"
	"github.com/ajalck/Go-gRPC-Microservice_Project/api_gateway/pkg/auth/middleware"
	"github.com/ajalck/Go-gRPC-Microservice_Project/api_gateway/pkg/config"
	"github.com/ajalck/Go-gRPC-Microservice_Project/api_gateway/pkg/order/client"
	"github.com/ajalck/Go-gRPC-Microservice_Project/api_gateway/pkg/order/handler"
	"github.com/gin-gonic/gin"
)

func OrderRoutes(r *gin.Engine, c *config.Config, authSvcC *authClient.ServiceClient) {
	authorize := middleware.InitMiddleware(authSvcC)

	ordersvc := client.ServiceClient{
		Client: client.InitServiceClient(c),
	}
	ordrHandler := handler.OrderHandler{C: ordersvc.Client}

	user := r.Group("/user")
	user.Use(authorize.Authorize)
	{
		user.POST("/createorder", ordrHandler.CreateOrder)
		user.PATCH("/cancelorder", ordrHandler.CancelOrder)
	}
}

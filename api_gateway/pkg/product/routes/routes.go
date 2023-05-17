package routes

import (
	authClient "github.com/ajalck/Go-gRPC-Microservice_Project/api_gateway/pkg/auth/client"

	"github.com/ajalck/Go-gRPC-Microservice_Project/api_gateway/pkg/auth/middleware"
	"github.com/ajalck/Go-gRPC-Microservice_Project/api_gateway/pkg/config"
	"github.com/ajalck/Go-gRPC-Microservice_Project/api_gateway/pkg/product/client"
	"github.com/ajalck/Go-gRPC-Microservice_Project/api_gateway/pkg/product/handler"
	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.Engine, c *config.Config, authSvcC *authClient.ServiceClient) {
	authorize := middleware.InitMiddleware(authSvcC)

	pdtSvc := client.ServiceClient{
		Client: client.InitServiceClient(c),
	}
	pdtHandler := handler.ProductHandler{PdtClient: pdtSvc.Client}
	admin := r.Group("/admin")
	{
		admin.POST("/createproduct", pdtHandler.CreateProduct)
		admin.PATCH("/updatestock", pdtHandler.UpdateStock)
		admin.GET("/viewproduct", pdtHandler.ViewProductByID)
	}
	user := r.Group("/user")
	{

		user.GET("/listproducts", authorize.Authorize, pdtHandler.ListProduct)
	}
}

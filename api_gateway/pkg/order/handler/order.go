package handler

import (
	"github.com/ajalck/Go-gRPC-Microservice_Project/order_management/pkg/pb"
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	C pb.OrderServiceClient
}

type CreateOrderRequestBody struct{
	
}

func (h *OrderHandler)CreateOrder(c *gin.Context){

}
func (h *OrderHandler)CancelOrder(c *gin.Context){

}
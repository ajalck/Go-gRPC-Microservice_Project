package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/ajalck/Go-gRPC-Microservice_Project/order_management/pkg/pb"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/status"
)

type OrderHandler struct {
	C pb.OrderServiceClient
}

type CreateOrderRequestBody struct {
	ProductId int32 `json:"product_id"`
	Quantity  int32 `json:"quantity"`
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	body := &CreateOrderRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	userId, _ := strconv.Atoi(c.Writer.Header().Get("User_id"))
	res, err := h.C.CreateOrder(context.Background(), &pb.CreateOrderRequest{
		UserId:    int32(userId),
		ProductId: body.ProductId,
		Quantity:  body.Quantity,
	})
	if err != nil {
		grpcError, ok := status.FromError(err)
		if ok {
			errMessage := grpcError.Message()
			c.JSON(http.StatusBadGateway, errMessage)
		}
		return
	}
	c.JSON(200, res)
}

type CancelOrderRequestBody struct {
	OrderId int32 `json:"order_id"`
}

func (h *OrderHandler) CancelOrder(c *gin.Context) {
	body := &CancelOrderRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	res, err := h.C.CancelOrder(context.Background(), &pb.CancelOrderRequest{
		OrderId: body.OrderId,
	})
	if err != nil {
		grpcError, ok := status.FromError(err)
		if ok {
			errMessage := grpcError.Message()
			c.JSON(http.StatusBadGateway, errMessage)
		}
		return
	}
	c.JSON(200, res)
}

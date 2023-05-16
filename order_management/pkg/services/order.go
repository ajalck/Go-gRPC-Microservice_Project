package services

import (
	"context"

	"github.com/ajalck/Go-gRPC-Microservice_Project/order_management/pkg/client"
	"github.com/ajalck/Go-gRPC-Microservice_Project/order_management/pkg/models"
	"github.com/ajalck/Go-gRPC-Microservice_Project/order_management/pkg/pb"
	productPB "github.com/ajalck/Go-gRPC-Microservice_Project/product_management/pkg/pb"
	"gorm.io/gorm"
)

type OrderServer struct {
	DB  *gorm.DB
	PSC *client.PdtServiceClient
	*pb.UnimplementedOrderServiceServer
}

func (s *OrderServer) CreateOrder(c context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {

	res, err := s.PSC.Client.ViewProductByID(context.Background(), &productPB.ViewProductRequest{
		ProductId: req.ProductId,
	})
	if err != nil {
		return &pb.CreateOrderResponse{Message: "Product not found"}, err
	}
	toOrder := models.Order{
		UserID:      req.UserId,
		ProductID:   req.ProductId,
		Quantity:    req.Quantity,
		TotalPrice:  (res.GetProduct().GetPrice()) * float32(req.Quantity),
		OrderStatus: "success",
	}
	result := s.DB.Create(&toOrder)
	if result.Error != nil {
		return &pb.CreateOrderResponse{Message: "Couldn't create order"}, err
	}
	if _, err = s.PSC.Client.UpdateStock(context.Background(), &productPB.UpdateStockRequest{
		ProductId: toOrder.ProductID,
		Stock:     res.Product.Stock - toOrder.Quantity,
	}); err != nil {
		return &pb.CreateOrderResponse{Message: "Failed to update order"}, err
	}
	return &pb.CreateOrderResponse{
		Message:     "Order Created Successfully",
		UserId:      toOrder.UserID,
		ProductId:   toOrder.ProductID,
		ProductName: res.Product.ProductName,
		Quantity:    toOrder.Quantity,
		TotalPrice:  toOrder.TotalPrice,
		OrderId:     int32(toOrder.ID),
		OrderStatus: toOrder.OrderStatus,
	}, nil
}
func (s *OrderServer) CancelOrder(c context.Context, req *pb.CancelOrderRequest) (*pb.CancelOrderResponse, error) {
	order := models.Order{}
	result := s.DB.Where("id", req.OrderId).First(&order)
	if result.Error != nil || order.OrderStatus == "cancelled" {
		return &pb.CancelOrderResponse{Message: "Order not found"}, result.Error
	}

	result = s.DB.Where("id", req.OrderId).Update("order_status", "cancelled")
	if result.Error != nil {
		return &pb.CancelOrderResponse{Message: "Failed to cancell order"}, result.Error
	}

	res, _ := s.PSC.Client.ViewProductByID(context.Background(), &productPB.ViewProductRequest{
		ProductId: order.ProductID,
	})

	if _, err := s.PSC.Client.UpdateStock(context.Background(), &productPB.UpdateStockRequest{
		ProductId: order.ProductID,
		Stock:     res.Product.Stock + order.Quantity,
	}); err != nil {
		return &pb.CancelOrderResponse{Message: "Failed to update order"}, err
	}

	return &pb.CancelOrderResponse{
		Message:     "Order Cancelled Successfully",
		UserId:      order.UserID,
		ProductId:   order.ProductID,
		Quantity:    order.Quantity,
		OrderStatus: order.OrderStatus,
	}, nil
}

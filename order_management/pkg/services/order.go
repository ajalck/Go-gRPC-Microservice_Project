package services

import (
	"context"

	"github.com/ajalck/Go-gRPC-Microservice_Project/order_management/pkg/client"
	"github.com/ajalck/Go-gRPC-Microservice_Project/order_management/pkg/pb"

	"gorm.io/gorm"
)

type OrderServer struct {
	DB  *gorm.DB
	PSC *client.PdtServiceClient
	*pb.UnimplementedOrderServiceServer
}

func (s *OrderServer) CreateOrder(c context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	return &pb.CreateOrderResponse{}, nil
}
func (s *OrderServer) CancelOrder(c context.Context, req *pb.CancelOrderRequest) (*pb.CancelOrderResponse, error) {
	return &pb.CancelOrderResponse{}, nil
}

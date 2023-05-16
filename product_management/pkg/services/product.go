package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/ajalck/Go-gRPC-Microservice_Project/product_management/pkg/models"
	"github.com/ajalck/Go-gRPC-Microservice_Project/product_management/pkg/pb"
	"gorm.io/gorm"
)

type ProductServer struct {
	DB *gorm.DB
	pb.UnimplementedProductServiceServer
}

func (s *ProductServer) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {

	product := &models.Products{
		ProductName: req.ProductName,
		Stock:       req.Stock,
		Price:       req.Price,
	}
	result := s.DB.Create(&product)
	if result.Error != nil {
		return &pb.CreateProductResponse{
			Message: "Could'nt add new product",
		}, result.Error
	}
	return &pb.CreateProductResponse{
		Message:   "New Product added successfully",
		ProductId: int32(product.ID),
	}, nil
}
func (s *ProductServer) ListProduct(ctx context.Context, req *pb.ListProductRequest) (*pb.ListProductResponse, error) {
	products := []models.Products{}
	result := s.DB.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	var pd []*pb.ProductDetails
	for _, product := range products {
		p := pb.ProductDetails{
			ProductId:   int32(product.ID),
			ProductName: product.ProductName,
			Stock:       product.Stock,
			Price:       product.Price,
		}
		pd = append(pd, &p)
	}
	return &pb.ListProductResponse{
		Products: pd,
	}, nil
}
func (s *ProductServer) ViewProductByID(ctx context.Context, req *pb.ViewProductRequest) (*pb.ViewProductResponse, error) {
	product := models.Products{}
	result := s.DB.Table("products").Where("id", req.ProductId).First(&product)
	if result.Error != nil {
		return &pb.ViewProductResponse{Message: "Product not found"}, result.Error
	}
	return &pb.ViewProductResponse{
		Product: &pb.ProductDetails{
			ProductId:   int32(product.ID),
			ProductName: product.ProductName,
			Stock:       product.Stock,
			Price:       product.Price,
		},
	}, nil
}
func (s *ProductServer) UpdateStock(ctx context.Context, req *pb.UpdateStockRequest) (*pb.UpdateStockResponse, error) {
	product := models.Products{}
	result := s.DB.Where("id", req.ProductId).First(&product)
	if result.Error != nil {
		fmt.Println(result.Error)
		return nil, errors.New("Product not found")
	}
	result = s.DB.Table("products").Where("id", req.ProductId).Update("stock", req.Stock)
	if result.Error != nil {
		fmt.Println(result.Error)
		return nil, result.Error
	}

	return &pb.UpdateStockResponse{
		Message:   "Product Updated Successfully",
		ProductId: int32(product.ID),
	}, nil
}

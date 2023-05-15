package service

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ajalck/Go-gRPC-Microservice_Project/auth_management/pkg/models"
	"github.com/ajalck/Go-gRPC-Microservice_Project/auth_management/pkg/pb"
	"gorm.io/gorm"
)

type AuthServer struct {
	DB *gorm.DB
	pb.UnimplementedAuthServiceServer
}

func (s *AuthServer) Register(c context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	fmt.Println("Auth Register : Register")

	user := models.User{
		Email:    req.Email,
		Password: req.Password,
	}
	result := s.DB.Create(&user)
	if result.Error != nil {
		return &pb.RegisterResponse{
			Status: http.StatusConflict,
			Error:  result.Explain(result.Error.Error()),
		}, result.Error
	}

	return &pb.RegisterResponse{
		Status: 200,
	}, nil
}
func (s *AuthServer) Login(c context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	return &pb.LoginResponse{}, nil
}
func (s *AuthServer) Validate(c context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	return &pb.ValidateResponse{}, nil
}

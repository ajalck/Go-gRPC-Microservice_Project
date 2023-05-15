package service

import (
	"context"
	"fmt"

	"github.com/ajalck/Go-gRPC-Microservice_Project/auth_management/pkg/models"
	"github.com/ajalck/Go-gRPC-Microservice_Project/auth_management/pkg/pb"
	"gorm.io/gorm"
)

type AuthServer struct {
	DB *gorm.DB
	pb.UnimplementedAuthServiceServer
}

func (s *AuthServer) Register(c context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	fmt.Println("Auth Service : Register")

	user := models.User{
		Email:    req.Email,
		Password: req.Password,
	}
	result := s.DB.Create(&user)
	if result.Error != nil {
		return &pb.RegisterResponse{
			Status:  400,
			Message: "Failed to Register new user",
		}, result.Error
	}
	var userid int64
	s.DB.Select("id").Where("email", req.Email).First(userid)
	return &pb.RegisterResponse{
		Status:  200,
		Message: "Registration successfull",
		Userid:  userid,
	}, nil
}
func (s *AuthServer) Login(c context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	return &pb.LoginResponse{}, nil
}
func (s *AuthServer) Validate(c context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	return &pb.ValidateResponse{}, nil
}

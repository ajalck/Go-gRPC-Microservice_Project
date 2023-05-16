package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/ajalck/Go-gRPC-Microservice_Project/auth_management/pkg/models"
	"github.com/ajalck/Go-gRPC-Microservice_Project/auth_management/pkg/pb"
	"github.com/ajalck/Go-gRPC-Microservice_Project/auth_management/pkg/utils"
	"gorm.io/gorm"
)

type AuthServer struct {
	DB         *gorm.DB
	Jwtwrapper utils.JWTWrapper
	pb.UnimplementedAuthServiceServer
}

func (s *AuthServer) Register(c context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	fmt.Println("Auth Service : Register")

	password, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, errors.New("Error in hashing password")
	}
	user := models.User{
		Email:    req.Email,
		Password: password,
	}
	result := s.DB.Create(&user)
	if result.Error != nil {
		return &pb.RegisterResponse{
			Status:  400,
			Message: "Failed to Register new user",
		}, result.Error
	}

	return &pb.RegisterResponse{
		Status:  200,
		Message: "Registration successfull",
		Userid:  int64(user.ID),
	}, nil
}
func (s *AuthServer) Login(c context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	user := &models.User{}
	result := s.DB.Where("email", req.Email).First(&user)
	if result.Error != nil {
		return &pb.LoginResponse{
			Status:  400,
			Message: "Failed to Login",
		}, errors.New("User not found")
	}
	if err := utils.CheckPasswordHash(req.Password, user.Password); err != nil {
		return &pb.LoginResponse{
			Status:  400,
			Message: "Failed to Login",
		}, errors.New("Mismatch in password")
	}
	token, err := s.Jwtwrapper.GenerateToken(*user)
	if err != nil {
		return &pb.LoginResponse{
			Status:  400,
			Message: "Failed to Login",
		}, err
	}
	return &pb.LoginResponse{
		Status:  200,
		Message: "Logged in Successfully",
		Userid:  int32(user.ID),
		Token:   token,
	}, nil
}
func (s *AuthServer) Validate(c context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {

	claims, err := s.Jwtwrapper.ValidateToken(req.Token)
	if err != nil {
		return nil, err
	}
	return &pb.ValidateResponse{
		Status:  200,
		Message: "Autherization Successfull",
		UserId:  claims.Id,
	}, nil
}

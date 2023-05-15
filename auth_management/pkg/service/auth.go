package service

import (
	"context"
	"github.com/ajalck/Go-gRPC-Microservice_Project/auth_management/pkg/pb"
)

type AuthServer struct {
	pb.UnimplementedAuthServiceServer
}

func (s *AuthServer) Register(c context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	return &pb.RegisterResponse{}, nil
}
func (s *AuthServer) Login(c context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	return &pb.LoginResponse{}, nil
}
func (s *AuthServer) Validate(c context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	return &pb.ValidateResponse{}, nil
}

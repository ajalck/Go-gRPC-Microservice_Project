package service

import (
	"context"
	"Go-gRPC-Microservice_Project/auth_management/pkg/Pb"
)

type AuthServer struct {
	Pb.UnimplementedAuthServiceServer
}

func (s *AuthServer) Register(c context.Context, req *Pb.RegisterRequest) (*Pb.RegisterResponse, error) {
	return &Pb.RegisterResponse{}, nil
}
func (s *AuthServer) Login(c context.Context, req *Pb.LoginRequest) (*Pb.LoginResponse, error) {
	return &Pb.LoginResponse{}, nil
}
func (s *AuthServer) Validate(c context.Context, req *Pb.ValidateRequest) (*Pb.ValidateResponse, error) {
	return &Pb.ValidateResponse{}, nil
}

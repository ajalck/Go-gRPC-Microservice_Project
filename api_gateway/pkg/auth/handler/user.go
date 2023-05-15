package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/ajalck/Go-gRPC-Microservice_Project/auth_management/pkg/pb"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	C pb.AuthServiceClient
}
type RegisterRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *AuthHandler) Register(ctx *gin.Context) {
	body := &RegisterRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.C.Register(context.Background(), &pb.RegisterRequest{
		Email:    body.Email,
		Password: body.Password,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	ctx.Writer.Header().Add("Content-Type", "application/json")
	ctx.Writer.Header().Set("user_id", strconv.Itoa(int(res.Userid)))
	ctx.JSON(int(res.Status), res.Message)

}

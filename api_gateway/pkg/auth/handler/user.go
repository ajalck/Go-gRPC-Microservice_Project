package handler

import (
	"context"
	"github.com/ajalck/Go-gRPC-Microservice_Project/auth_management/pkg/Pb"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	C Pb.AuthServiceClient
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

	res, err := h.C.Register(context.Background(), &Pb.RegisterRequest{
		Email:    body.Email,
		Password: body.Password,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	ctx.JSON(int(res.Status), res)

}

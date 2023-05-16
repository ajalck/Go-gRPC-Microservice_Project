package handler

import (
	"context"
	"net/http"

	"github.com/ajalck/Go-gRPC-Microservice_Project/product_management/pkg/pb"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	PdtClient pb.ProductServiceClient
}
type CreateProductRequestBody struct {
	ProductName string  `json:"product_name"`
	Stock       int32   `json:"stock"`
	Price       float32 `json:"price"`
}

func (h *ProductHandler) CreateProduct(ctx *gin.Context) {
	body := &CreateProductRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Inputs"})
		return
	}

	res, err := h.PdtClient.CreateProduct(context.Background(), &pb.CreateProductRequest{
		ProductName: body.ProductName,
		Stock:       body.Stock,
		Price:       body.Price,
	})
	if err != nil {
		ctx.JSON(http.StatusNotFound, err)
		return
	}
	ctx.JSON(200, res)
}
func (h *ProductHandler) ListProduct(ctx *gin.Context) {
	res, err := h.PdtClient.ListProduct(context.Background(), &pb.ListProductRequest{})
	if err != nil {
		ctx.JSON(http.StatusNotFound, err)
		return
	}
	ctx.JSON(200, res)
}

type ViewProductRequestBody struct {
	ProductId int32 `json:"product_id" gorm:"not null"`
}

func (h *ProductHandler) ViewProductByID(c *gin.Context) {
	body := &ViewProductRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	res, err := h.PdtClient.ViewProductByID(context.Background(), &pb.ViewProductRequest{
		ProductId: body.ProductId,
	})
	if err != nil {
		c.JSON(http.StatusBadGateway, res)
		return
	}
	c.JSON(200, res)
}

type UpdateProductRequestBody struct {
	ProductId int32 `json:"product_id" gorm:"not null" binding:"required,numeric"`
	Stock     int32 `json:"stock" gorm:"not null"`
}

func (h *ProductHandler) UpdateStock(ctx *gin.Context) {
	body := &UpdateProductRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	res, err := h.PdtClient.UpdateStock(context.Background(), &pb.UpdateStockRequest{
		ProductId: body.ProductId,
		Stock:     body.Stock,
	})
	if err != nil {
		ctx.JSON(http.StatusBadGateway, err)
		return
	}
	ctx.JSON(200, res)
}

package middleware

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	auth "github.com/ajalck/Go-gRPC-Microservice_Project/api_gateway/pkg/auth/client"
	"github.com/ajalck/Go-gRPC-Microservice_Project/auth_management/pkg/pb"
	"github.com/gin-gonic/gin"
)

type MiddlewareServiceClient struct {
	authClient *auth.ServiceClient
}

func InitMiddleware(client *auth.ServiceClient) *MiddlewareServiceClient {
	return &MiddlewareServiceClient{
		authClient: client,
	}
}
func (m *MiddlewareServiceClient) Authorize(c *gin.Context) {
	authtoken := c.Request.Header.Get("Authorization")
	if authtoken == "" {
		c.JSON(http.StatusUnauthorized, "Failed to Authorize")
		return
	}
	token := strings.Split(authtoken, "Bearer ")
	if len(token) > 2 {
		c.JSON(http.StatusUnauthorized, "Failed to Authorize")
		return
	}
	res, err := m.authClient.Client.Validate(context.Background(), &pb.ValidateRequest{
		Token: token[1],
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, err)
		return
	}
	c.Writer.Header().Set("User_id", strconv.Itoa(int(res.UserId)))
	c.JSON(int(res.Status), res.GetMessage())
}

package utils

import (
	"errors"
	"time"

	"github.com/ajalck/Go-gRPC-Microservice_Project/auth_management/pkg/models"
	"github.com/golang-jwt/jwt"
)

type JWTWrapper struct {
	SecretKey string
	Issuer    string
}
type jwtClaims struct {
	Id    int64
	Email string
	jwt.StandardClaims
}

func (w *JWTWrapper) GenerateToken(user models.User) (signedToken string, err error) {
	claims := &jwtClaims{
		Id:    int64(user.ID),
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * 1).Unix(),
			Issuer:    w.Issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err = token.SignedString([]byte(w.SecretKey))
	if err != nil {
		return signedToken, err
	}
	return signedToken, err
}
func (w *JWTWrapper) ValidateToken(signedToken string) (claims *jwtClaims, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken, &jwtClaims{}, func(t *jwt.Token) (interface{}, error) {
			return []byte(w.SecretKey), nil
		})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*jwtClaims)
	if !ok {
		return nil, errors.New("Could'nt parse claims")
	}
	if claims.ExpiresAt < time.Now().Unix() {
		return nil, errors.New("Oops! Session expired")
	}
	return claims, nil
}

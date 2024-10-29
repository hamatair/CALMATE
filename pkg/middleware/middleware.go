package middleware

import (
	"github.com/bccfilkom-be/go-server/internal/usecase"
	"github.com/bccfilkom-be/go-server/pkg/jwt"
	"github.com/gin-gonic/gin"
)

type Interface interface {
	AuthenticateUser(ctx *gin.Context)
	Timeout() gin.HandlerFunc
}

type middleware struct {
	jwtAuth jwt.Interface
	usecase *usecase.Usecase
}

func Init(jwtAuth jwt.Interface, usecase *usecase.Usecase) Interface {
	return &middleware{
		jwtAuth: jwtAuth,
		usecase: usecase,
	}
}

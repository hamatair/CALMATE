package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/bccfilkom-be/go-server/pkg/model"
	"github.com/bccfilkom-be/go-server/pkg/response"
	"github.com/gin-gonic/gin"
)

func (m *middleware) AuthenticateUser(ctx *gin.Context) {
	bearer := ctx.GetHeader("Authorization")
	if bearer == "" {
		response.Error(ctx, http.StatusUnauthorized, "empty token", errors.New(""))
		ctx.Abort()
	}

	token := strings.Split(bearer, " ")[1]
	idPengguna, err := m.jwtAuth.ValidateToken(token)
	if err != nil {
		response.Error(ctx, http.StatusUnauthorized, "failed validate token", err)
		ctx.Abort()
	}
	
	pengguna, err := m.usecase.PenggunaUsecase.GetPengguna(model.PenggunaParam{
		IDPengguna: idPengguna.String(),
	})
	if err != nil {
		response.Error(ctx, http.StatusUnauthorized, "failed get pengguna", err)
		ctx.Abort()
	}

	param := model.PenggunaParam{
		IDPengguna: pengguna.IDPengguna,
		Email: pengguna.Email,
		Password: pengguna.Password,
	}

	ctx.Set("pengguna", param)

	ctx.Next()
}

package main

import (
	"github.com/bccfilkom-be/go-server/internal/repository"
	"github.com/bccfilkom-be/go-server/internal/rest"
	"github.com/bccfilkom-be/go-server/internal/usecase"
	"github.com/bccfilkom-be/go-server/pkg/bcrypt"
	"github.com/bccfilkom-be/go-server/pkg/config"
	"github.com/bccfilkom-be/go-server/pkg/database/mysql"
	"github.com/bccfilkom-be/go-server/pkg/jwt"
	"github.com/bccfilkom-be/go-server/pkg/middleware"
)

func main() {

	config.LoadEnv()

	jwtAuth := jwt.Init()

	bcrypt := bcrypt.Init()

	db := mysql.ConnectDatabase()

	repository := repository.NewRepository(db)

	usecase := usecase.NewUsecase(usecase.InitParam{
		Repository: repository,
		JwtAuth: jwtAuth,
		Bcrypt: bcrypt,
	})

	middleware := middleware.Init(jwtAuth, usecase)

	rest := rest.NewRest(usecase, middleware)

	mysql.Migration(db)

	rest.MountEndpoint()

	rest.Serve()
}

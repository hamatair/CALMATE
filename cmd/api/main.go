package main

import (
	// "net/http"

	"github.com/bccfilkom-be/go-server/internal/repository"
	"github.com/bccfilkom-be/go-server/internal/rest"
	"github.com/bccfilkom-be/go-server/internal/usecase"
	"github.com/bccfilkom-be/go-server/pkg/bcrypt"
	"github.com/bccfilkom-be/go-server/pkg/config"
	"github.com/bccfilkom-be/go-server/pkg/database/mysql"
	"github.com/bccfilkom-be/go-server/pkg/jwt"
	"github.com/bccfilkom-be/go-server/pkg/middleware"
	// "github.com/go-chi/chi/v5"
	// "github.com/go-chi/chi/v5/middleware"
	// "gorm.io/gorm"
)

// func mux() http.Handler {
// 	r := chi.NewRouter()
// 	r.Use(middleware.RequestID)
// 	r.Use(middleware.Logger)

// 	v1 := chi.NewRouter()
// 	v1.Get("/status", func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([]byte("Healthy"))
// 	})
// 	r.Mount("/api/v1", v1)

// 	return r
// }

func main() {
	// server := &http.Server{Addr: "0.0.0.0:8080", Handler: mux()}
	// ctx, cancel := context.WithCancel(context.Background())
	// sig := make(chan os.Signal, 1)
	// signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	// go func() {
	// 	<-sig
	// 	shutdownCtx, cancelShutdownCtx := context.WithTimeout(ctx, 30*time.Second)

	// 	go func() {
	// 		<-shutdownCtx.Done()
	// 		if shutdownCtx.Err() == context.DeadlineExceeded {
	// 			log.Fatal("graceful shutdown timed out.. forcing exit.")
	// 		}
	// 	}()

	// 	err := server.Shutdown(shutdownCtx)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	cancelShutdownCtx()
	// 	cancel()
	// }()

	// fmt.Printf("server running on %s\n", server.Addr)
	// err := server.ListenAndServe()
	// if err != nil && err != http.ErrServerClosed {
	// 	log.Fatal(err)
	// }

	// <-ctx.Done()

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

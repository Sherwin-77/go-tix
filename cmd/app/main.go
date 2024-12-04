package main

import (
	"context"
	"fmt"
	echoSwagger "github.com/swaggo/echo-swagger"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4/middleware"
	"github.com/sherwin-77/go-echo-template/configs"
	"github.com/sherwin-77/go-echo-template/internal/builder"
	"github.com/sherwin-77/go-echo-template/internal/http/handler"
	"github.com/sherwin-77/go-echo-template/pkg/caches"
	"github.com/sherwin-77/go-echo-template/pkg/database"
	"github.com/sherwin-77/go-echo-template/pkg/server"

	_ "github.com/sherwin-77/go-echo-template/docs"
)

//	@title			go-echo-template
//	@version		1.0
//	@description	This is a sample server for go-echo-template.

//	@host		localhost:8080
//	@BasePath	/api/v1

// @securityDefinitions.bearerauth
func main() {
	config := configs.LoadConfig()

	db, err := database.InitDB(config.Postgres)
	if err != nil {
		panic(err)
	}

	cache := caches.NewCache(caches.InitRedis(config.Redis))

	echoServer := server.NewServer()
	echoServer.GET("/swagger/*", echoSwagger.WrapHandler)
	echoServer.Use(middleware.LoggerWithConfig(configs.GetEchoLoggerConfig()))
	echoServer.Use(middleware.RecoverWithConfig(configs.GetEchoRecoverConfig()))
	echoServer.Validator = configs.NewAppValidator()
	echoServer.HTTPErrorHandler = handler.HTTPErrorHandler

	group := echoServer.Group("/api")
	builder.BuildV1Routes(config, db, cache, group)

	runServer(echoServer, config)
	waitForShutdown(echoServer)
}

func runServer(s *server.Server, config *configs.Config) {
	go func() {
		if err := s.Start(fmt.Sprintf("0.0.0.0:%s", config.Port)); err != nil {
			s.Logger.Fatal(err)
		}
	}()
}

func waitForShutdown(s *server.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	go func() {
		if err := s.Shutdown(ctx); err != nil {
			s.Logger.Fatal(err)
		}
	}()
}

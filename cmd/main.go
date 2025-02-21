package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/emaforlin/bussiness-service/config"
	endpoints "github.com/emaforlin/bussiness-service/endpoint"
	"github.com/emaforlin/bussiness-service/pb"
	"github.com/emaforlin/bussiness-service/repository"
	"github.com/emaforlin/bussiness-service/service"
	"github.com/emaforlin/bussiness-service/transport"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config.Init()

	cfg := config.GetConfig()

	logger := zap.NewExample()
	defer logger.Sync()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT, syscall.SIGALRM)

	database := repository.NewMySQLConnection(logger)
	err := database.Connect()
	if err != nil {
		logger.Fatal("error connecting to the database", zap.Error(err))
	}

	svc := service.NewService(logger, database)
	businessEdpoint := endpoints.MakeEndpoints(svc)
	grpcServer := transport.NewGRPCServer(businessEdpoint, logger)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.App.Port))
	if err != nil {
		logger.Fatal("error to listen on", zap.Uint16("port", cfg.App.Port))
	}

	baseServer := grpc.NewServer()
	pb.RegisterBusinessServer(baseServer, grpcServer)

	reflection.Register(baseServer)

	go func() {
		if err := baseServer.Serve(listener); err != nil {
			logger.Fatal("Failed to start the server", zap.Error(err))
		}
	}()

	logger.Info("Server started successfully")

	// Handle graceful shutdown
	<-quit
	logger.Info("Shutting down the server")
	baseServer.GracefulStop()
	listener.Close()
	close(quit)
	logger.Info("Server stopped successfully")

}

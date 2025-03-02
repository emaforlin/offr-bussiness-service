package server

import (
	"context"
	"fmt"
	"net"

	"github.com/emaforlin/bussiness-service/config"
	"github.com/emaforlin/bussiness-service/pb"
	"github.com/emaforlin/bussiness-service/service"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	pb.UnimplementedBusinessServer
	businessService *service.BusinessService
}

func NewServer(bs *service.BusinessService, cfg *config.Config) *Server {
	return &Server{businessService: bs}
}

func StartGRPCServer() fx.Option {
	return fx.Invoke(func(lc fx.Lifecycle, cfg *config.Config, businessSvc *service.BusinessService, log *zap.Logger) {
		grpcServer := grpc.NewServer()

		serverImpl := NewServer(businessSvc, cfg)

		// Register reflection
		if cfg.App.DevMode {
			reflection.Register(grpcServer)
		}

		pb.RegisterBusinessServer(grpcServer, serverImpl)

		lc.Append(fx.Hook{
			OnStart: func(ctx context.Context) error {
				addr := fmt.Sprintf(":%d", cfg.App.Port)
				lis, err := net.Listen("tcp", addr)
				if err != nil {
					return err
				}

				go grpcServer.Serve(lis)
				log.Info("Listening on ", zap.String("address", addr))
				return nil
			},
			OnStop: func(ctx context.Context) error {
				grpcServer.GracefulStop()
				log.Info("Server stopped.")
				return nil
			},
		})
	})
}

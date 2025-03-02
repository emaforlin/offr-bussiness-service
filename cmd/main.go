package main

import (
	"github.com/emaforlin/bussiness-service/config"
	"github.com/emaforlin/bussiness-service/logger"
	"github.com/emaforlin/bussiness-service/repository"
	"github.com/emaforlin/bussiness-service/server"
	"github.com/emaforlin/bussiness-service/service"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
		config.ProvideConfig(),
		logger.ProvideZapLogger(),
		repository.ProvideRepo(),
		service.ProvideBusinessService(),
		server.StartGRPCServer(),
	).Run()
}

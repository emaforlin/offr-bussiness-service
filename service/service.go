package service

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type BusinessService struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewBusinessService(db *gorm.DB, log *zap.Logger) *BusinessService {
	return &BusinessService{db: db, logger: log}
}

func ProvideBusinessService() fx.Option {
	return fx.Provide(NewBusinessService)
}

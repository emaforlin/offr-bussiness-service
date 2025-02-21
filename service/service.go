package service

import (
	"context"
	"errors"

	"github.com/emaforlin/bussiness-service/entities"
	"github.com/emaforlin/bussiness-service/repository"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Service interface {
	CreateBusiness(ctx context.Context, data entities.Business) (*entities.Business, error)
}

type service struct {
	logger *zap.Logger
	db     repository.GMDatabase
}

func (s *service) CreateBusiness(ctx context.Context, business entities.Business) (*entities.Business, error) {
	if err := s.db.Cursor().Create(&business).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, errors.New("already exists a business with that address")
		}
	}
	return &business, nil
}

func NewService(logger *zap.Logger, db repository.GMDatabase) Service {
	return &service{
		logger: logger,
		db:     db,
	}
}

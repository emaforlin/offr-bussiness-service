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
	CreateBusiness(ctx context.Context, data entities.CreateBusinessDto) (*entities.Business, error)
}

type service struct {
	logger *zap.Logger
	db     repository.GMDatabase
}

func (s *service) CreateBusiness(ctx context.Context, business entities.CreateBusinessDto) (*entities.Business, error) {
	var dbBusiness = &entities.Business{
		Name:    business.Name,
		Address: business.Address,
		Staff:   business.Staff,
	}
	if err := s.db.Cursor().Create(dbBusiness).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, errors.New("already exists a business with that address")
		}
	}
	return dbBusiness, nil
}

func NewService(logger *zap.Logger, db repository.GMDatabase) Service {
	return &service{
		logger: logger,
		db:     db,
	}
}

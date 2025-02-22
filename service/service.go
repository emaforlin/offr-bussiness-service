package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/emaforlin/bussiness-service/entities"
	"github.com/emaforlin/bussiness-service/repository"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Service interface {
	CreateBusiness(ctx context.Context, data entities.CreateBusinessDto) (*entities.Business, error)
	DeleteBusiness(ctx context.Context, deleteID uint) (*uint, error)
}

type service struct {
	logger *zap.Logger
	db     repository.GMDatabase
}

// DeleteBusiness implements Service.
func (s *service) DeleteBusiness(ctx context.Context, deleteID uint) (*uint, error) {
	err := s.db.Cursor().Delete(&entities.Business{}, deleteID).Error

	if err != nil {
		return nil, errors.New(fmt.Sprintf("business with id %d not found", deleteID))
	}
	s.logger.Debug("Business deleted")
	return &deleteID, nil
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
	s.logger.Debug("Business created")
	return dbBusiness, nil
}

func NewService(logger *zap.Logger, db repository.GMDatabase) Service {
	return &service{
		logger: logger,
		db:     db,
	}
}

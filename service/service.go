package service

import (
	"context"

	"github.com/emaforlin/bussiness-service/entities"
	"github.com/emaforlin/bussiness-service/repository"
	"go.uber.org/zap"
)

type Service interface {
	CreateBusiness(ctx context.Context, data entities.CreateBusinessDto) (*entities.Business, error)
	DeleteBusiness(ctx context.Context, deleteID uint) (*uint, error)
	SendJoinInvitation(ctx context.Context, invitation entities.InvitationDto) (string, error)
}

type service struct {
	logger *zap.Logger
	db     repository.GMDatabase
}

func NewService(logger *zap.Logger, db repository.GMDatabase) Service {
	return &service{
		logger: logger,
		db:     db,
	}
}

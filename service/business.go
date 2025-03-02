package service

import (
	"github.com/emaforlin/bussiness-service/dto"
	"github.com/emaforlin/bussiness-service/entities"
	"go.uber.org/zap"
)

func (s *BusinessService) CreateBusiness(data dto.CreateBusinessDto) (*entities.Business, error) {
	var dbBusiness = &entities.Business{
		Name:    data.Name,
		Address: data.Address,
	}

	if err := dbBusiness.Create(s.db); err != nil {
		s.logger.Error("failed to create New Business", zap.Error(err))
		return nil, err
	}
	s.logger.Info("New Business created.")
	return dbBusiness, nil
}

func (s *BusinessService) DeleteBusiness(id uint64) error {
	var dbBusiness = new(entities.Business)
	err := dbBusiness.Delete(s.db, id)

	if err != nil {
		// err = ErrOnBusinessDelete
		s.logger.Error("failed to delete Business", zap.Error(err))
		return err
	}
	s.logger.Info("Business deleted.")
	return nil
}

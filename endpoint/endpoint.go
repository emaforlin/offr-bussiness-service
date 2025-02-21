package endpoints

import (
	"context"

	"github.com/emaforlin/bussiness-service/entities"
	"github.com/emaforlin/bussiness-service/service"
	gkendpoint "github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	CreateBusiness gkendpoint.Endpoint
}

func MakeEndpoints(s service.Service) Endpoints {
	return Endpoints{
		CreateBusiness: makeCreateBusinessEndpoint(s),
	}
}

func makeCreateBusinessEndpoint(s service.Service) gkendpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(entities.Business)

		res, err := s.CreateBusiness(ctx, req)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
}

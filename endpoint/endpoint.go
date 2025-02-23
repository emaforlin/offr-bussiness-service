package endpoints

import (
	"context"

	"github.com/emaforlin/bussiness-service/entities"
	"github.com/emaforlin/bussiness-service/service"
	gkendpoint "github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	CreateBusiness     gkendpoint.Endpoint
	DeleteBusiness     gkendpoint.Endpoint
	SendJoinInvitation gkendpoint.Endpoint
}

func MakeEndpoints(s service.Service) Endpoints {
	return Endpoints{
		CreateBusiness:     makeCreateBusinessEndpoint(s),
		DeleteBusiness:     makeDeleteBusinessEndpoint(s),
		SendJoinInvitation: makeSendJoinInvitationEndpoint(s),
	}
}

func makeSendJoinInvitationEndpoint(s service.Service) gkendpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(entities.InvitationDto)
		token, err := s.SendJoinInvitation(ctx, req)
		if err != nil {
			return nil, err
		}
		return token, nil
	}
}

func makeDeleteBusinessEndpoint(s service.Service) gkendpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		id := request.(uint64)
		res, err := s.DeleteBusiness(ctx, uint(id))
		if err != nil {
			return nil, err
		}
		return res, nil

	}
}

func makeCreateBusinessEndpoint(s service.Service) gkendpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(entities.CreateBusinessDto)

		res, err := s.CreateBusiness(ctx, req)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
}

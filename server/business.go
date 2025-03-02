package server

import (
	"context"
	"fmt"

	"github.com/emaforlin/bussiness-service/dto"
	"github.com/emaforlin/bussiness-service/pb"
	"go.uber.org/fx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteBusiness(ctx context.Context, req *pb.DeleteBusinessRequest) (*pb.DeleteBusinessResponse, error) {
	err := s.businessService.DeleteBusiness(req.GetId())
	if err != nil {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("couldn't delete, business with id %d not found", req.GetId()))
	}

	return &pb.DeleteBusinessResponse{
		Id: &req.Id,
	}, nil
}

func (s *Server) CreateNewBusiness(ctx context.Context, req *pb.CreateBusinessRequest) (*pb.CreateBusinessResponse, error) {
	created, err := s.businessService.CreateBusiness(dto.CreateBusinessDto{
		Name:      req.GetName(),
		Address:   req.GetAddress(),
		ManagerID: req.GetBusinessManager(),
	})

	if err != nil {
		return nil, status.Error(codes.AlreadyExists, err.Error())
	}
	fx.Invoke()
	return &pb.CreateBusinessResponse{
		Id: uint64(created.ID),
	}, nil
}

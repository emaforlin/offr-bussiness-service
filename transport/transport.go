package transport

import (
	"context"

	endpoints "github.com/emaforlin/bussiness-service/endpoint"
	"github.com/emaforlin/bussiness-service/entities"
	"github.com/emaforlin/bussiness-service/pb"
	gt "github.com/go-kit/kit/transport/grpc"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type gRPCServer struct {
	pb.UnimplementedBusinessServer
	createBusiness     gt.Handler
	deleteBusiness     gt.Handler
	sendJoinInvitation gt.Handler
}

func (s *gRPCServer) SendJoinInvitation(ctx context.Context, req *pb.InvitationRequest) (*pb.InvitationResponse, error) {
	_, resp, err := s.sendJoinInvitation.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.InvitationResponse), nil
}

func decodeInvitationReq(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.InvitationRequest)
	return entities.InvitationDto{
		InviterID:      req.GetInviterID(),
		RecipientEmail: req.GetRecipientEmail(),
		AssignedRole:   req.GetAssignedRole(),
	}, nil
}

func encodeInvitationResp(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(string)
	if res == "" {
		return nil, status.Error(codes.InvalidArgument, "failed to send invitation")
	}

	return &pb.InvitationResponse{
		Token: res,
	}, nil
}

func (s *gRPCServer) DeleteBusiness(ctx context.Context, req *pb.DeleteBusinessRequest) (*pb.DeleteBusinessResponse, error) {
	_, resp, err := s.deleteBusiness.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.DeleteBusinessResponse), nil
}

func decodeDeleteBusinessReq(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.DeleteBusinessRequest)
	return req.GetId(), nil
}

func encodeDeleteBusinessResp(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(*uint)
	id := uint64(*res)

	if res == nil {
		return nil, status.Error(codes.InvalidArgument, "could not delete business")
	}
	return &pb.DeleteBusinessResponse{
		Id: &id,
	}, nil
}

func (s *gRPCServer) CreateNewBusiness(ctx context.Context, req *pb.CreateBusinessRequest) (*pb.CreateBusinessResponse, error) {
	_, resp, err := s.createBusiness.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.CreateBusinessResponse), nil
}

func decodeCreateBusinessReq(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.CreateBusinessRequest)

	staff := make([]entities.Staff, 1)
	staff[0] = entities.Staff{
		Auth0ID: req.GetBusinessManager(),
	}
	return entities.CreateBusinessDto{
		Name:    req.GetName(),
		Address: req.GetAddress(),
		Staff:   staff,
	}, nil
}

func encodeCreateBusinessResp(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(*entities.Business)
	if res.ID == 0 {
		return nil, status.Error(codes.InvalidArgument, "failed to create a new business")
	}
	return &pb.CreateBusinessResponse{
		Id: uint64(res.ID),
	}, nil
}

func NewGRPCServer(endpoints endpoints.Endpoints, logger *zap.Logger) pb.BusinessServer {
	return &gRPCServer{
		createBusiness: gt.NewServer(
			endpoints.CreateBusiness,
			decodeCreateBusinessReq,
			encodeCreateBusinessResp,
		),
		deleteBusiness: gt.NewServer(
			endpoints.DeleteBusiness,
			decodeDeleteBusinessReq,
			encodeDeleteBusinessResp,
		),
		sendJoinInvitation: gt.NewServer(
			endpoints.SendJoinInvitation,
			decodeInvitationReq,
			encodeInvitationResp,
		),
	}
}

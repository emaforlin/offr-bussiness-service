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
	createBusiness gt.Handler
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
	return entities.Business{
		Name:    req.GetName(),
		Address: req.GetAddress(),
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
	}
}

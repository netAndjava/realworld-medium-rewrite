package transport

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	gt "github.com/go-kit/kit/transport/grpc"
	"iohttps.com/live/realworld-medium-rewrite/domain"
	"iohttps.com/live/realworld-medium-rewrite/endpoints"
	pb "iohttps.com/live/realworld-medium-rewrite/service/api"
)

type grpcServer struct {
	handler gt.Handler
}

func NewArticleGRPCServer(endpoint endpoint.Endpoint, logger log.Logger) pb.ArticleServiceServer {
	return &grpcServer{
		handler: gt.NewServer(endpoint, decodeSaveReq, encodeSaveResp, nil),
	}
}

func (s *grpcServer) Save(ctx context.Context, req *pb.SaveRequest) (*pb.SaveResponse, error) {
	_, resp, err := s.handler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.SaveResponse), nil

}

func decodeSaveReq(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.SaveRequest)
	return endpoints.SaveReq{Article: domain.Article{ID: domain.NUUID(req.Article.Id), Title: req.Article.Title, Content: req.Article.Content, Status: domain.Draft, AuthorID: domain.NUUID(req.Article.AuthorId)}}, nil
}

func encodeSaveResp(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.SaveResponse)
	return endpoints.SaveResp{ID: domain.NUUID(resp.Id)}, nil
}

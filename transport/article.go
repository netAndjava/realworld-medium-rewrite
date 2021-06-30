package transport

import (
	"context"

	"github.com/go-kit/kit/log"
	gt "github.com/go-kit/kit/transport/grpc"
	"iohttps.com/live/realworld-medium-rewrite/domain"
	"iohttps.com/live/realworld-medium-rewrite/endpoints"
	pb "iohttps.com/live/realworld-medium-rewrite/service/api"
)

type grpcServer struct {
	write gt.Handler
}

func NewArticleGRPCServer(endpoints endpoints.Endpoints, logger log.Logger) pb.ArticleServiceServer {
	return &grpcServer{
		write: gt.NewServer(endpoints.Write, decodeWriteReq, encodeWriteResp, nil),
	}
}

func (s *grpcServer) Save(ctx context.Context, req *pb.WriteRequest) (*pb.WriteResponse, error) {
	_, resp, err := s.write.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.WriteResponse), nil

}

func decodeWriteReq(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.WriteRequest)
	return endpoints.WriteReq{Article: domain.Article{ID: domain.NUUID(req.Article.Id), Title: req.Article.Title, Content: req.Article.Content, Status: domain.Draft, AuthorID: domain.NUUID(req.Article.AuthorId)}}, nil
}

func encodeWriteResp(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.WriteResponse)
	return endpoints.WriteResp{ID: domain.NUUID(resp.Id)}, nil
}

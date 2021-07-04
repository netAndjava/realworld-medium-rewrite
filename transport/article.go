package transport

import (
	"context"
	"log"

	gt "github.com/go-kit/kit/transport/grpc"
	"iohttps.com/live/realworld-medium-rewrite/domain"
	"iohttps.com/live/realworld-medium-rewrite/endpoints"
	pb "iohttps.com/live/realworld-medium-rewrite/service/api"
)

type grpcServer struct {
	write                    gt.Handler
	viewDraftArticles        gt.Handler
	view                     gt.Handler
	publish                  gt.Handler
	viewPublicArticles       gt.Handler
	viewRecentArticles       gt.Handler
	viewDraftOfPublicArticle gt.Handler
	republish                gt.Handler
	drop                     gt.Handler
}

func NewArticleGRPCServer(endpoints endpoints.Endpoints, logger log.Logger) pb.ArticleServiceServer {
	return &grpcServer{
		write:             gt.NewServer(endpoints.Write, decodeWriteReq, encodeWriteResp, nil),
		viewDraftArticles: gt.NewServer(endpoints.ViewDraftArticles, decodeViewDraftArticlesReq, encodeViewDraftArticleResp, nil),
	}
}

func (s *grpcServer) Write(ctx context.Context, req *pb.WriteRequest) (*pb.WriteResponse, error) {
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
	resp := response.(endpoints.WriteResp)
	return &pb.WriteResponse{Id: int64(resp.ID)}, nil
}

func (s *grpcServer) ViewDraftedArticles(ctx context.Context, req *pb.ViewDraftedArticlesRequest) (*pb.ViewDraftedArticlesResponse, error) {
	_, resp, err := s.viewDraftArticles.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ViewDraftedArticlesResponse), nil
}

func decodeViewDraftArticlesReq(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ViewDraftedArticlesRequest)
	return endpoints.ViewDraftArticlesReq{UserID: domain.NUUID(req.UserId)}, nil
}

func encodeViewDraftArticleResp(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.ViewDraftArticlesResp)
	return &pb.ViewDraftedArticlesResponse{Articles: ConvertArticles(resp.Articles)}, nil
}

func (s *grpcServer) View(ctx context.Context, req *pb.ViewRequest) (*pb.ViewResponse, error) {
	_, resp, err := s.view.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ViewResponse), nil
}

func (s *grpcServer) Publish(_ context.Context, _ *pb.PublishRequest) (*pb.PublishResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (s *grpcServer) ViewPublicArticles(_ context.Context, _ *pb.ViewPublicArticlesRequest) (*pb.ViewPublicArticlesResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (s *grpcServer) ViewRecentArticles(_ context.Context, _ *pb.ViewRecentArticlesRequest) (*pb.ViewRecentArticlesResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (s *grpcServer) ViewDraftOfPublicArticle(_ context.Context, _ *pb.ViewDraftOfPublicArticleRequest) (*pb.ViewDraftOfPublicArticleResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (s *grpcServer) Republish(_ context.Context, _ *pb.RepublishRequest) (*pb.RepublishResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (s *grpcServer) Drop(_ context.Context, _ *pb.DropArticleRequest) (*pb.DropArticleResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (s *grpcServer) mustEmbedUnimplementedArticleServiceServer() {
	panic("not implemented") // TODO: Implement
}

//ConvertArticles .....
func ConvertArticles(arts []domain.Article) []*pb.Article {

	articles := make([]*pb.Article, len(arts))
	for i, art := range arts {
		articles[i] = &pb.Article{Id: int64(art.ID), Title: art.Title, Content: art.Content, Status: int32(art.Status), AuthorId: int64(art.AuthorID)}
	}
	return articles
}

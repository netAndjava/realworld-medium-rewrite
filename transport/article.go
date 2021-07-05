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
		write:                    gt.NewServer(endpoints.Write, decodeWriteReq, encodeWriteResp, nil),
		viewDraftArticles:        gt.NewServer(endpoints.ViewDraftArticles, decodeViewDraftArticlesReq, encodeViewDraftArticleResp, nil),
		view:                     gt.NewServer(endpoints.View, decodeViewReq, encodeViewResp),
		publish:                  gt.NewServer(endpoints.Publish, decodePublishReq, encodePublishResp),
		viewPublicArticles:       gt.NewServer(endpoints.ViewPublicArticles, decodeViewPublicArticleReq, encodeViewPublicArticleResp, nil),
		viewRecentArticles:       gt.NewServer(endpoints.ViewRecentArticles, nil, encodeViewRecentArticlesResp, nil),
		viewDraftOfPublicArticle: gt.NewServer(endpoints.ViewDraftOfPublicArticle, decodeViewReq, encodeViewResp, nil),
		republish:                gt.NewServer(endpoints.Republish, decodePublishReq, encodePublishResp, nil),
		drop:                     gt.NewServer(endpoints.Drop, decodeDropReq, encodeDropResp, nil),
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

func decodeViewReq(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ViewRequest)
	return endpoints.ViewReq{ID: domain.NUUID(req.Id)}, nil
}

func encodeViewResp(ctx context.Context, response interface{}) (interface{}, error) {
	req := response.(endpoints.ViewResp)
	return &pb.ViewResponse{Article: ConvertArticle(req.Article)}, nil
}

func (s *grpcServer) Publish(ctx context.Context, req *pb.PublishRequest) (*pb.PublishResponse, error) {
	_, resp, err := s.publish.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.PublishResponse), nil
}

func decodePublishReq(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.PublishRequest)
	return endpoints.PublishReq{Article: domain.Article{ID: domain.NUUID(req.Article.Id), Content: req.Article.Content, Title: req.Article.Content, AuthorID: domain.NUUID(req.Article.AuthorId)}}, nil
}

func encodePublishResp(ctx context.Context, response interface{}) (interface{}, error) {
	return &pb.PublishResponse{}, nil
}

func (s *grpcServer) ViewPublicArticles(ctx context.Context, req *pb.ViewPublicArticlesRequest) (*pb.ViewPublicArticlesResponse, error) {
	_, resp, err := s.viewPublicArticles.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ViewPublicArticlesResponse), nil
}

func decodeViewPublicArticleReq(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ViewPublicArticlesRequest)
	return endpoints.ViewPublicArticlesReq{UserID: domain.NUUID(req.UserId)}, nil
}

func encodeViewPublicArticleResp(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.ViewPublicArticlesResp)
	return &pb.ViewPublicArticlesResponse{Articles: ConvertArticles(resp.Articles)}, nil
}

func (s *grpcServer) ViewRecentArticles(ctx context.Context, req *pb.ViewRecentArticlesRequest) (*pb.ViewRecentArticlesResponse, error) {
	_, resp, err := s.viewRecentArticles.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ViewRecentArticlesResponse), nil
}

func encodeViewRecentArticlesResp(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.ViewRecentArticlesResp)
	return &pb.ViewRecentArticlesResponse{Articles: ConvertArticles(resp.Articles)}, nil
}

func (s *grpcServer) ViewDraftOfPublicArticle(ctx context.Context, request *pb.ViewDraftOfPublicArticleRequest) (*pb.ViewDraftOfPublicArticleResponse, error) {
	_, resp, err := s.viewDraftOfPublicArticle.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ViewDraftOfPublicArticleResponse), nil
}

func (s *grpcServer) Republish(ctx context.Context, request *pb.RepublishRequest) (*pb.RepublishResponse, error) {
	_, resp, err := s.republish.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.RepublishResponse), nil
}

func (s *grpcServer) Drop(ctx context.Context, request *pb.DropArticleRequest) (*pb.DropArticleResponse, error) {
	_, resp, err := s.drop.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.DropArticleResponse), nil
}

func decodeDropReq(ctx context.Context, req interface{}) (interface{}, error) {
	request := req.(*pb.DropArticleRequest)
	return endpoints.DropReq{ID: domain.NUUID(request.ArticleId), UserID: domain.NUUID(request.UserId)}, nil
}

func encodeDropResp(ctx context.Context, resp interface{}) (interface{}, error) {
	return endpoints.DropResp{}, nil
}

func (s *grpcServer) mustEmbedUnimplementedArticleServiceServer() {
	//不知道这个函数什么作用
	panic("not implemented") // TODO: Implement
}

func ConvertArticle(art domain.Article) *pb.Article {
	return &pb.Article{Id: int64(art.ID), Title: art.Title, Content: art.Content, Status: int32(art.Status), AuthorId: int64(art.AuthorID)}
}

//ConvertArticles .....
func ConvertArticles(arts []domain.Article) []*pb.Article {

	articles := make([]*pb.Article, len(arts))
	for i, art := range arts {
		articles[i] = &pb.Article{Id: int64(art.ID), Title: art.Title, Content: art.Content, Status: int32(art.Status), AuthorId: int64(art.AuthorID)}
	}
	return articles
}

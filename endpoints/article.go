package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"iohttps.com/live/realworld-medium-rewrite/domain"
	"iohttps.com/live/realworld-medium-rewrite/service/article"
	"iohttps.com/live/realworld-medium-rewrite/usecases"
)

type Endpoints struct {
	Write endpoint.Endpoint
}

type WriteReq struct {
	Article domain.Article
}

type WriteResp struct {
	ID domain.NUUID
}

func makeEndpoints(s article.ArticleService) Endpoints {
	return Endpoints{
		Write: makeWriteEndpoint(s),
	}
}

func makeWriteEndpoint(s article.ArticleService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(WriteReq)
		ID, err := s.Write(ctx, usecases.GenerateUUID, req.Article)
		return WriteResp{ID: ID}, err
	}
}

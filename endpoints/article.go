package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"iohttps.com/live/realworld-medium-rewrite/domain"
	"iohttps.com/live/realworld-medium-rewrite/usecases"
)

type SaveDraftReq struct {
	Article domain.Article
}

type SaveDraftResp struct {
}

type Endpoints struct {
	Add endpoint.Endpoint
}

func MakeEndpoints(s usecases.ArticleService) Endpoints {
	return Endpoints{
		Add: makeAddEndpoints(s),
	}
}

func makeAddEndpoints(s usecases.ArticleService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SaveDraftReq)
		err := s.SaveDraft(req.Article)
		return SaveDraftResp{}, err
	}
}

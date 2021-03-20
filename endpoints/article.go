package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"iohttps.com/live/realworld-medium-rewrite/domain"
	"iohttps.com/live/realworld-medium-rewrite/usecases"
)

type SaveReq struct {
	Article domain.Article
}

type SaveResp struct {
	ID domain.NUUID
}

func makeSaveDraftEndpoints(s usecases.ArticleService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SaveReq)
		err := s.SaveDraft(req.Article)
		return SaveResp{}, err
	}
}

package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"iohttps.com/live/realworld-medium-rewrite/domain"
	"iohttps.com/live/realworld-medium-rewrite/service/article"
)

type Endpoints struct {
	Write                    endpoint.Endpoint
	ViewDraftArticles        endpoint.Endpoint
	View                     endpoint.Endpoint
	Publish                  endpoint.Endpoint
	ViewPublicArticles       endpoint.Endpoint
	ViewRecentArticles       endpoint.Endpoint
	ViewDraftOfPublicArticle endpoint.Endpoint
	Republish                endpoint.Endpoint
	Drop                     endpoint.Endpoint
}

func makeEndpoints(s article.ArticleService) Endpoints {
	return Endpoints{
		Write:                    makeWriteEndpoint(s),
		ViewDraftArticles:        makeViewDraftArticlesEndpoint(s),
		View:                     makeViewEndpoint(s),
		Publish:                  makePublishEndpoint(s),
		ViewPublicArticles:       makeViewPublicArticlesEndpoint(s),
		ViewRecentArticles:       makeViewRecentArticlesEndpoint(s),
		ViewDraftOfPublicArticle: makeViewDraftOfPublicArticleEndpoint(s),
		Republish:                makePublishEndpoint(s),
		Drop:                     makeDropEndpoint(s),
	}
}

func makeWriteEndpoint(s article.ArticleService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(WriteReq)
		ID, err := s.Write(ctx, req.Article)
		return WriteResp{ID: ID}, err
	}
}

type WriteReq struct {
	Article domain.Article
}

type WriteResp struct {
	ID domain.NUUID
}

func makeViewDraftArticlesEndpoint(s article.ArticleService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(ViewDraftArticlesReq)
		articles, err := s.ViewDraftArticles(ctx, req.UserID)
		return ViewDraftArticlesResp{Articles: articles}, err
	}
}

type ViewDraftArticlesReq struct {
	UserID domain.NUUID
}

type ViewDraftArticlesResp struct {
	Articles []domain.Article
}

func makeViewEndpoint(s article.ArticleService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(ViewReq)
		article, err := s.View(ctx, req.ID)
		return ViewResp{Article: article}, err
	}
}

type ViewReq struct {
	ID domain.NUUID
}

type ViewResp struct {
	Article domain.Article
}

func makePublishEndpoint(s article.ArticleService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(PublishReq)
		err := s.Publish(ctx, req.Article)
		return PunlishResp{}, err
	}
}

type PublishReq struct {
	Article domain.Article
}

func makeViewPublicArticlesEndpoint(s article.ArticleService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ViewPublicArticlesReq)
		articles, err := s.ViewPublicArticles(ctx, req.UserID)
		return ViewPublicArticlesResp{Articles: articles}, err
	}
}

type PunlishResp struct {
}

type ViewPublicArticlesReq struct {
	UserID domain.NUUID
}

func makeViewRecentArticlesEndpoint(s article.ArticleService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		articles, err := s.ViewRecentArticles(ctx)
		return ViewRecentArticlesResp{Articles: articles}, err
	}
}

type ViewPublicArticlesResp struct {
	Articles []domain.Article
}

type ViewRecentArticlesResp struct {
	Articles []domain.Article
}

func makeViewDraftOfPublicArticleEndpoint(s article.ArticleService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ViewReq)
		article, err := s.ViewDraftOfPublicArticle(ctx, req.ID)
		return ViewResp{Article: article}, err
	}
}

func makeRepublishEndpoint(s article.ArticleService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(PublishReq)
		err := s.Republish(ctx, req.Article)
		return PunlishResp{}, err
	}
}

func makeDropEndpoint(s article.ArticleService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DropReq)
		err := s.Drop(ctx, req.ID, req.UserID)
		return DropResp{}, err
	}
}

type DropReq struct {
	ID     domain.NUUID
	UserID domain.NUUID
}

type DropResp struct {
}

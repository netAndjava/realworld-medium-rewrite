// Package article provides ...
package article

import (
	"context"

	"github.com/go-kit/kit/log"
	"iohttps.com/live/realworld-medium-rewrite/domain"
	"iohttps.com/live/realworld-medium-rewrite/usecases"
)

type ArticleService interface {
	Write(ctx context.Context, a domain.Article) (domain.NUUID, error)
	ViewDraftArticles(ctx context.Context, userID domain.NUUID) ([]domain.Article, error)
	View(ctx context.Context, ID domain.NUUID) (domain.Article, error)
	Publish(ctx context.Context, a domain.Article) error
	ViewPublicArticles(ctx context.Context, userID domain.NUUID) ([]domain.Article, error)
	ViewRecentArticles(ctx context.Context) ([]domain.Article, error)
	ViewDraftOfPublicArticle(ctx context.Context, ID domain.NUUID) (domain.Article, error)
	Republish(ctx context.Context, a domain.Article) error
	Drop(ctx context.Context, ID, userID domain.NUUID) error
}

type articleService struct {
	Itor   usecases.ArticleInteractor
	logger log.Logger
}

func NewArticleService(logger log.Logger, itor usecases.ArticleInteractor) ArticleService {
	return &articleService{Itor: itor, logger: logger}
}

func (as *articleService) Write(ctx context.Context, a domain.Article) (domain.NUUID, error) {
	art, _ := as.Itor.View(a.ID)
	if art.ID == 0 {
		return as.Itor.Write(usecases.GenerateUUID, a)
	}

	if art.ID != 0 && art.Status == domain.Draft {
		//2.保存编辑的文章
		err := as.Itor.EditDraftArticle(a)
		return a.ID, err
	}

	err := as.Itor.EditPublicArticle(a)
	return a.ID, err
}

func (as *articleService) ViewDraftArticles(ctx context.Context, userID domain.NUUID) ([]domain.Article, error) {
	return as.Itor.ViewDraftArticles(userID)
}

func (as *articleService) View(ctx context.Context, ID domain.NUUID) (domain.Article, error) {
	return as.Itor.View(ID)
}

func (as *articleService) Publish(ctx context.Context, a domain.Article) error {
	_, err := as.Write(ctx, a)
	if err != nil {
		return err
	}
	return as.Itor.Publish(a.ID, a.AuthorID)
}

func (as *articleService) ViewPublicArticles(ctx context.Context, userID domain.NUUID) ([]domain.Article, error) {
	return as.Itor.ViewPublicArticles(userID)
}

func (as *articleService) ViewRecentArticles(ctx context.Context) ([]domain.Article, error) {
	return as.Itor.ViewRecentArticles()
}

func (as *articleService) ViewDraftOfPublicArticle(ctx context.Context, ID domain.NUUID) (domain.Article, error) {
	return as.Itor.ViewDraftOfPublicArticle(ID)
}

func (as *articleService) Republish(ctx context.Context, a domain.Article) error {
	return as.Itor.Republish(a)
}

func (as *articleService) Drop(ctx context.Context, ID, userID domain.NUUID) error {
	return as.Itor.Drop(ID, userID)
}

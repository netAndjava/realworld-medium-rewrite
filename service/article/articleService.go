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
}

type articleService struct {
	Itor   usecases.ArticleInteractor
	logger log.Logger
}

// find a bug
// 你先手动吧

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

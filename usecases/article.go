// Package usecases provides ...
package usecases

import (
	"errors"

	"iohttps.com/live/realworld-medium-rewrite/domain"
)

type ArticleService interface {
	SaveDraft(a domain.Article) error
}

// ArticleInteractor article interactor
type ArticleInteractor struct {
	ArticleRepo domain.ArticleRepository
}

// Write the user writes an article
func (itor ArticleInteractor) Wrtie(generate func() domain.NUUID, a domain.Article) (domain.NUUID, error) {
	if len(a.Title) == 0 || len(a.Content) == 0 {
		return domain.NUUID(0), errors.New("用户内容为空")
	}
	a.ID = generate()
	err := itor.ArticleRepo.Create(a)

	return a.ID, err
}

// EditDraftArticle the user edits a draft article
func (itor ArticleInteractor) EditDraftArticle(a domain.Article) error {
	art, err := itor.ArticleRepo.Get(a.ID)
	if err != nil {
		return err
	}
	if art.AuthorID != a.AuthorID {
		return errors.New("用户没有修改文章权限")
	}
	return itor.ArticleRepo.Save(art)
}

//Publish:the user publishes a draft article
func (itor ArticleInteractor) Publish(articleID, userID domain.NUUID) error {
	article, err := itor.ArticleRepo.Get(articleID)
	if err != nil {
		return err
	}
	if err := article.Check(); err != nil {
		return err
	}
	if article.AuthorID != userID {
		return errors.New("用户没有发布文章权限")
	}
	return itor.ArticleRepo.Publish(articleID)
}

// ViewDraftArticles the author views draft articles
func (itor ArticleInteractor) ViewDraftArticles(userID domain.NUUID) ([]domain.Article, error) {
	return itor.ArticleRepo.ViewDraftArticles(userID)
}

//View the user views an article
func (itor ArticleInteractor) View(ID domain.NUUID) (domain.Article, error) {
	if ID == 0 {
		return domain.Article{}, errors.New("文章不存在")
	}
	return itor.ArticleRepo.Get(ID)
}

// ViewPublicArticle 获取作者的已发布文章
func (itor ArticleInteractor) ViewPublicArticles(userID domain.NUUID) ([]domain.Article, error) {
	return itor.ArticleRepo.ViewPublicArticles(userID)
}

// ViewRecentArticles 读者查看最近发布的文章
func (itor ArticleInteractor) ViewRecentArticles() ([]domain.Article, error) {
	return itor.ArticleRepo.GetAllPublicArticles()
}

// ViewDraftOfPublicArticle 作者查看已发布文章的草稿
func (itor ArticleInteractor) ViewDraftOfPublicArticle(ID domain.NUUID) (domain.Article, error) {
	if ID == 0 {
		return domain.Article{}, errors.New("文章不存在")
	}
	art, err := itor.ArticleRepo.ViewDraftOfPublicArticle(ID)
	if err != nil && err != domain.ErrNotFound {
		return domain.Article{}, err
	}
	if err == domain.ErrNotFound {
		return itor.View(ID)
	}
	return art, nil
}

// EditPublicArticle 用户编辑已发布文章
func (itor ArticleInteractor) EditPublicArticle(a domain.Article) error {
	art, err := itor.ArticleRepo.Get(a.ID)
	if err != nil {
		return err
	}

	if art.AuthorID != a.AuthorID {
		return errors.New("用户没有创建已发布文章草稿权限")
	}
	art, err = itor.ViewDraftOfPublicArticle(a.ID)
	if err != nil && err == domain.ErrNotFound {
		return itor.ArticleRepo.UpdateDraftOfPublicArticle(a)
	}

	if err != nil {
		return err
	}
	return itor.ArticleRepo.CreateDraftOfPublicArticle(a)
}

// Republish 用户重新发布文章
func (itor ArticleInteractor) Republish(a domain.Article) error {
	art, err := itor.ViewDraftOfPublicArticle(a.ID)
	if err != nil {
		return err
	}
	if art.AuthorID != a.AuthorID {
		return errors.New("用户没有发布权限")
	}
	if err := art.Check(); err != nil {
		return err
	}
	return itor.ArticleRepo.Republish(art)
}

//Drop 删除文章
func (itor ArticleInteractor) Drop(ID, userID domain.NUUID) error {
	art, err := itor.View(ID)
	if err != nil && err == domain.ErrNotFound {
		// TODO:分开处理not found和连接错误  <15-03-21, nqq> //
		return errors.New("已删除")
	}

	if err != nil {
		return err
	}

	if art.AuthorID != userID {
		return errors.New("用户没有删除权限")
	}
	return itor.ArticleRepo.Drop(ID)
}

// GenerateUUID 生成Number类型的id
func GenerateUUID() domain.NUUID {
	return domain.NUUID(0)
}

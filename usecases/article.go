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

// SaveDraft 保存草稿
func (itor ArticleInteractor) SaveDraft(a domain.Article) error {
	art, err := itor.ArticleRepo.Get(a.ID)
	if err != nil {
		return err
	}
	if art.AuthorID != a.AuthorID {
		return errors.New("用户没有修改文章权限")
	}
	return itor.ArticleRepo.Save(art)
}

// CreateDraft 创建草稿
func (itor ArticleInteractor) CreateDraft(generate func() domain.NUUID, a domain.Article) (domain.NUUID, error) {
	if len(a.Title) == 0 || len(a.Content) == 0 {
		return domain.NUUID(0), errors.New("用户内容为空")
	}
	a.ID = generate()
	err := itor.ArticleRepo.Create(a)

	return a.ID, err
}

//Publish 发布文章
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

// GetAuthorDrafts 获取作者的草稿列表
func (itor ArticleInteractor) GetAuthorDrafts(userID domain.NUUID) ([]domain.Article, error) {
	return itor.ArticleRepo.GetAuthorDrafts(userID)
}

// GetAuthorPublicArticles 获取作者的已发布文章
func (itor ArticleInteractor) GetAuthorPublicArticles(userID domain.NUUID) ([]domain.Article, error) {
	return itor.ArticleRepo.GetAuthorPublicArticles(userID)
}

//GetArticle 获取文章详情
func (itor ArticleInteractor) GetArticle(ID domain.NUUID) (domain.Article, error) {
	return itor.ArticleRepo.Get(ID)
}

// GetAllPublicArticles 获取所有已发布文章
func (itor ArticleInteractor) GetAllPublicArticles() ([]domain.Article, error) {
	return itor.ArticleRepo.GetAllPublicArticles()
}

// GetPublicArticleDraft 获取对修改已发布文章编辑的草稿
func (itor ArticleInteractor) GetPublicArticleDraft(ID domain.NUUID) (domain.Article, error) {
	art, err := itor.ArticleRepo.GetDraftOfPublicArticle(ID)
	if err != nil && err != domain.ErrNotFound {
		return domain.Article{}, err
	}
	if err == domain.ErrNotFound {
		return itor.GetArticle(ID)
	}
	return art, nil
}

// SavePublicArticleDraft 保存已发布文章草稿
func (itor ArticleInteractor) SavePublicArticleDraft(a domain.Article, userID domain.NUUID) error {
	art, err := itor.ArticleRepo.Get(a.ID)
	if err != nil {
		return err
	}

	if art.AuthorID != userID {
		return errors.New("用户没有创建已发布文章草稿权限")
	}
	return itor.ArticleRepo.UpdateDraftOfPublicArticle(a)
}

//CreatePublicArticleDraft 创建已发布文章草稿
func (itor ArticleInteractor) CreatePublicArticleDraft(a domain.Article, userID domain.NUUID) error {
	art, err := itor.ArticleRepo.Get(a.ID)
	if err != nil {
		return err
	}
	if art.AuthorID != userID {
		return errors.New("用户没有创建已发布文章草稿权限")
	}
	return itor.ArticleRepo.CreateDraftOfPublicArticle(a)
}

// PublishPublicArticleDraft 发布对已发布文章的修改草稿
func (itor ArticleInteractor) PublishPublicArticleDraft(a domain.Article, userID domain.NUUID) error {
	art, err := itor.GetPublicArticleDraft(a.ID)
	if err != nil {
		return err
	}
	if art.AuthorID != userID {
		return errors.New("用户没有发布权限")
	}
	if err := art.Check(); err != nil {
		return err
	}
	return itor.ArticleRepo.PublishPublicArticleDraft(art)
}

//Drop 删除文章
func (itor ArticleInteractor) Drop(ID, userID domain.NUUID) error {
	art, err := itor.GetArticle(ID)
	if err != nil {
		// TODO:分开处理not found和连接错误  <15-03-21, nqq> //
		return errors.New("已删除")
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

// Package usecases provides ...
package usecases

import "iohttps.com/live/realworld-medium-rewrite/domain"

// ArticleInteractor article interactor
type ArticleInteractor struct {
	ArticleRepo domain.ArticleRepository
}

// SaveDraft 保存草稿
func (itor ArticleInteractor) SaveDraft(generate func() domain.NUUID, a domain.Article) (domain.NUUID, error) {
	// TODO:  <12-11-20, nqq> //
	return domain.NUUID(0), nil
}

//Publish 发布文章
func (itor ArticleInteractor) Publish(ID, userID domain.NUUID) error {
	return nil
}

// GetAuthorDrafts 获取作者的草稿列表
func (itor ArticleInteractor) GetAuthorDrafts(userID domain.NUUID) ([]domain.Article, error) {
	return []domain.Article{}, nil
}

// GetAuthorPublicArticles 获取作者的已发布文章
func (itor ArticleInteractor) GetAuthorPublicArticles(userID domain.NUUID) ([]domain.Article, error) {
	return []domain.Article{}, nil
}

//GetArticleDetail 获取文章详情
func (itor ArticleInteractor) GetArticle()

// GetAllPublicArticles 获取所有已发布文章
func (itor ArticleInteractor) GetAllPublicArticles()

// GetPublicArticleDraft 获取对修改已发布文章编辑的草稿
func (itor ArticleInteractor) GetPublicArticleDraft()

// SavePublicArticleDraft 保存已发布文章草稿
func (itor ArticleInteractor) SavePublicArticleDraft()

// PublishPublicArticleDraft 发布对已发布文章的修改草稿
func (itor ArticleInteractor) PublishPublicArticleDraft()

// GenerateUUID 生成树枝类型的id
func GenerateUUID() domain.NUUID {
}

// Package uscecases provides ...
package uscecases

//ArticleInteractor article interactor
type ArticleInteractor struct {
	ArticleRepo domain.ArticleRepository
}

//SaveDraft 保存草稿
func(itor ArticleInteractor)SaveDraft

//Public 发布文章
func(itor ArticleInteractor)Public

//GetAuthorDrafts 获取作者的草稿列表
func(itor ArticleInteractor)GetAuthorDrafts

//GetAuthorPublicArticles 获取作者的已发布文章
func(itor ArticleInteractor)GetAuthorPublicArticles

//GetAllPublicArticles 获取所有已发布文章
func(itor ArticleInteractor)GetAllPublicArticles

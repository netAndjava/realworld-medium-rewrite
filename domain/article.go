// Package domain provides ...
package domain

import "errors"

//PublicStatus 文章发布状态
type PublicStatus int8

const (
	//Draft 草稿
	Draft PublicStatus = iota
	//Public 已发布
	Public
)

//NUUID 数字类型UUID
type NUUID int64

//Article 文章实体
type Article struct {
	ID      NUUID
	Title   string
	Content string
	Status  PublicStatus
	Author  User
}

func (art Article) Check() error {
	if len(art.Title) == 0 {
		return errors.New("文章标题为空")
	}
	if len(art.Content) == 0 {
		return errors.New("文章内容为空")
	}
	return nil
}

//ArticleRepository article repository
type ArticleRepository interface {
	Create(a Article) error
	Save(a Article) error
	Publish(ID NUUID) error
	GetAuthorDrafts(userID NUUID) ([]Article, error)
	GetAuthorPublicArticles(userID NUUID) ([]Article, error)
	GetAllPublicArticles() ([]Article, error)
	Get(ID NUUID) (Article, error)
	GetDraftOfPublicArticle(ID NUUID) (Article, error)
	CreateDraftOfPublicArticle(a Article) error
	UpdateDraftOfPublicArticle(a Article) error
}

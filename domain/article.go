// Package domain provides ...
package domain

import "iohttps.com/nqq/product-news-collect/domain"

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
}

//ArticleRepository article repository
type ArticleRepository interface {
	Create(a Article) error
	Publish(ID NUUID) error
	GetAuthorDrafts(userID NUUID) ([]domain.Article, error)
}

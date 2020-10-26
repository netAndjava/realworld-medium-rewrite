// Package domain provides ...
package domain

//PublicStatus 文章发布状态
type PublicStatus int8

const (
	//Draft 草稿
	Draft PublicStatus = iota
	//Public 已发布
	Public
)

//Article 文章实体
type Article struct {
	Title   string
	Content string
	Status  PublicStatus
}

//ArticleRepository article repository
type ArticleRepository interface {
}

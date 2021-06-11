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

var (
	// ErrNotFound article not found? // TODO: <10-12-20, bantana>
	ErrNotFound = errors.New("没找到")
)

// TODO: UUID是(universally unique identifier 通用唯一标识,前面再加个number让人更难理解了,
// 还有在domain里面的命名一般都是terms, 如果这个是domain的terms,应该给个domain的命名,
// 如果不是domain直接相关的通常是放到common或者base中去) <10-12-20, bantana> //

// NUUID is a Number UUID
//NUUID 数字类型UUID
type NUUID int64

//Article 文章实体
type Article struct {
	ID       NUUID
	Title    string
	Content  string
	Status   PublicStatus
	AuthorID NUUID
}

// Check is valid method
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

	ViewDraftArticles(userID NUUID) ([]Article, error)
	Get(ID NUUID) (Article, error)
	ViewPublicArticles(userID NUUID) ([]Article, error)

	GetAllPublicArticles() ([]Article, error)
	ViewDraftOfPublicArticle(ID NUUID) (Article, error)
	CreateDraftOfPublicArticle(a Article) error
	UpdateDraftOfPublicArticle(a Article) error

	Republish(a Article) error
	Drop(ID NUUID) error
}

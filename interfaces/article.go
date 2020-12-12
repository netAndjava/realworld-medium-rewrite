// Package interfaces provides ...
package interfaces

import (
	"iohttps.com/live/realworld-medium-rewrite/domain"
	"iohttps.com/live/realworld-medium-rewrite/infrastructure/database"
)

//ArticleRepo ......
type ArticleRepo database.DbRepo

//NewArticleRepo ......
func NewArticleRepo(helper database.DbHandler) domain.ArticleRepository {
	return &ArticleRepo{DbHandler}
}

//Create .....
func (repo *ArticleRepo) Create(a domain.Article) error {
	_, err := repo.Handler.Execute(`insert into t_article (id,title,content,status,userID) values(?,?,?,?,?)`, a.ID, a.Title, a.Status, a.AuthorID)
	return err
}

//Save ......
func (repo *ArticleRepo) Save(a domain.Article) error {
	_, err := repo.Handler.Execute(`update t_article set title=?,content=?`, a.Title, a.Content)
	return err
}

//Publish ......
func (repo *ArticleRepo) Publish(ID domain.NUUID) error {
	_, err := repo.Handler.Execute(`update t_article set status=? where ID=?`, domain.Draft, ID)
	return err
}

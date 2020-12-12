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
	_, err := repo.Handler.Execute(`update t_article set status=? where id=?`, domain.Draft, ID)
	return err
}

//GGetAuthorDrafts .....
func (repo *ArticleRepo) GetAuthorDrafts(userID domain.NUUID) ([]domain.Article, error) {
	return repo.GetAuthorArticleByStatus(userID, domain.Draft)
}

//GetAuthorPublicArticles .......
func (repo *ArticleRepo) GetAuthorPublicArticles(userID domain.NUUID) ([]domain.Article, error) {
	return repo.GetAuthorArticleByStatus(userID, domain.Public)
}

func (repo *ArticleRepo) GetAuthorArticleByStatus(userID domain.NUUID, status domain.PublicStatus) ([]domain.Article, error) {
	var articles []domain.Article
	rows, err := repo.Handler.Query(`select id,title,content where status=?,userId=?`, status, userID)
	if err != nil {
		return []domain.Article{}, err
	}

	for rows.Next() {
		var article domain.Article
		err = rows.Scan(&article.ID, &article.Title, &article.Content)
		if err != nil {
			return []domain.Article{}, err
		}
		article.Status = domain.Draft
		article.AuthorID = userID
		articles = append(articles, article)
	}

	return articles, nil
}
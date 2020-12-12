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

func (repo *ArticleRepo) GetAllPublicArticles() ([]domain.Article, error) {

	var articles []domain.Article
	rows, err := repo.Handler.Query(`select id,title,content,status,userId from t_article`)
	if err != nil {
		return []domain.Article{}, err
	}

	for rows.Next() {
		var article domain.Article
		err := rows.Scan(&article.ID, &article.Title, &article.Content, &article.Status, &article.AuthorID)
		if err != nil {
			return []domain.Article{}, err
		}
		articles = append(articles, article)

	}
	return articles, nil

}

//Get ......
func (repo *ArticleRepo) Get(ID domain.NUUID) (domain.Article, error) {
	row := repo.Handler.QueryRow(`select title,content,status,userId from t_article where id=?`, ID)
	var a domain.Article
	err := row.Scan(&a.Title, &a.Content, &a.Status, &a.AuthorID)
	a.ID = ID
	return a, err
}

func (repo *ArticleRepo) GetDraftOfPublicArticle(ID domain.NUUID) (domain.Article, error) {
	row := repo.Handler.QueryRow(`select title,content from t_draft where id=?`, ID)

	var art domain.Article
	art.ID = ID
	err := row.Scan(&art.Title, &art.Content)
	return art, err
}

func (repo *ArticleRepo) CreateDraftOfPublicArticle(a domain.Article) error {
	_, err := repo.Handler.Execute(`insert into t_draft (id,title,content) values(?,?,?)`, a.ID, a.Title, a.Content)
	return err
}

func (repo *ArticleRepo) UpdateDraftOfPublicArticle(a domain.Article) error {
	_, err := repo.Handler.Execute(`update t_draft set title=?,content=? where id=?`, a.Title, a.Content, a.ID)
	return err
}

func (repo *ArticleRepo)PublishPublicArticleDraft(ID domain.NUUID)error{
	repo.Handler.Execute(, args ...interface{})
}

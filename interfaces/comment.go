// Package interfaces provides ...
package interfaces

import (
	"iohttps.com/live/realworld-medium-rewrite/domain"
	"iohttps.com/live/realworld-medium-rewrite/infrastructure/database"
)

type CommentRepo database.DbRepo

func NewCommentRepo(helper database.DbHandler) domain.CommentRepository {
	return &CommentRepo{Handler: helper}
}

func (repo *CommentRepo) Add(c domain.Comment) error {
	_, err := repo.Handler.Execute(`insert into t_comment (id,pid,articleID,content,userID) values(?,?,?,?,?)`, c.ID, c.PID, c.ArticleID, c.Content, c.UserID)
	return err
}

func (repo *CommentRepo) Get(ID domain.NUUID) (domain.Comment, error) {
	row := repo.Handler.QueryRow(`select pid,articleID,content,userID from t_comment where id=?`, id)
	var c domain.Comment
	c.ID = id
	err := row.Scan(&c.PID, &c.ArticleID, &c.Content, &c.Creator)
	return c, err
}

func (repo *CommentRepo) GetCommentByPID(PID domain.NUUID) ([]domain.Comment, error) {
	rows, err := repo.Handler.Query(`select id,articleID,content,userID from t_comment where pid=?`, PID)
	if err != nil {
		return []domain.Comment{}, err
	}
	var comments []domain.Comment
	for rows.Next() {
		var c domain.Comment
		err := rows.Scan(&c.ID, &c.ArticleID, &c.Content, &c.Creator)
		if err != nil {
			// TODO:记录错误日志  <16-12-20, nqq> //
			continue
		}
		c.PID = PID
		comments = append(comments, c)
	}
	return comments, nil
}

func (repo *CommentRepo) GetCommentsByArticleID(articleID domain.NUUID) ([]domain.Comment, error) {
	rows, err := repo.Handler.Query(`select id,pid,content,userID from t_comment where articleID=?`, articleID)
	if err != nil {
		return []domain.Comment{}, err
	}
	var comments []domain.Comment
	for rows.Next() {
		var c domain.Comment
		err := rows.Scan(&c.ID, &c.ArticleID, &c.Content, &c.Creator)
		if err != nil {
			// TODO:记录错误日志  <16-12-20, nqq> //
			continue
		}
		c.PID = PID
		comments = append(comments, c)
	}
	return comments, nil
}

func (repo *CommentRepo) Drop(ID domain.NUUID) error {
	_, err := repo.Handler.Execute(`delete from t_comment where id=?`, ID)
	return err
}

func (repo *CommentRepo) DropByPID(PID domain.NUUID) error {
	_, err := repo.Handler.Execute(`delete from t_comment where pid=?`, PID)
	return err
}

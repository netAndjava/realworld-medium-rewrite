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

// func (repo *CommentRepo)Get(ID domain.NUUID)(domain.Comment,error){
// 	repo.Handler.QueryRow(`select pid,articleID,content,userID from t_comment where id=?`, id)
// }

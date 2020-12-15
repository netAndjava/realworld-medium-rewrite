// Package interfaces provides ...
package interfaces

import (
	"iohttps.com/live/realworld-medium-rewrite/infrastructure/database"
	"iohttps.com/live/realworld-medium-rewrite/usecases"
)

type TokenRepo database.DbRepo

func NewTokenRepo(handler database.DbHandler) usecases.TokenRepository {
	return &TokenRepo{Handler: handler}
}

func (repo *TokenRepo) Save(t usecases.Token) error {
	_, err := repo.Handler.Execute(`insert into t_token (id,userID,expiredAt) values(?,?,?)`, t.ID, t.UserID, t.ExpiredAt)
	return err
}

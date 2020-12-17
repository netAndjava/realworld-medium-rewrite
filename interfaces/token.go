// Package interfaces provides ...
package interfaces

import (
	"iohttps.com/live/realworld-medium-rewrite/infrastructure/database"
	"iohttps.com/live/realworld-medium-rewrite/usecases"
)

//TokenRepo .........
type TokenRepo database.DbRepo

//NewTokenRepo ........
func NewTokenRepo(handler database.DbHandler) usecases.TokenRepository {
	return &TokenRepo{Handler: handler}
}

//Save .........
func (repo *TokenRepo) Save(t usecases.Token) error {
	_, err := repo.Handler.Execute(`insert into t_token (id,userID,expiredAt) values(?,?,?)`, t.ID, t.UserID, t.ExpiredAt)
	return err
}

//Get ........
func (repo *TokenRepo) Get(tokenID usecases.SUUID) (usecases.Token, error) {
	row := repo.Handler.QueryRow(`select userID,expiredAt from t_token where id=?`, tokenID)
	var token usecases.Token
	token.ID = tokenID
	err := row.Scan(&token.UserID, &token.ExpiredAt)
	return token, err
}

//Delete .......
func (repo *TokenRepo) Delete(tokenID usecases.SUUID) error {
	_, err := repo.Handler.Execute(`delete from t_token where id=?`, tokenID)
	return err
}

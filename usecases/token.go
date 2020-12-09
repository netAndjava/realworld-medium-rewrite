// Package usecases provides ...
package usecases

import (
	"errors"

	"iohttps.com/live/realworld-medium-rewrite/domain"
)

type Token struct {
	ID     SUUID
	UserID domain.NUUID
}

type SUUID string

type TokenRepository interface {
	Save(t Token) error
}

type TokenInteractor struct {
	TokenRepos TokenRepository
}

//Login 判断用户登录
func (itor TokenInteractor) Login(userID domain.NUUID, generate func() SUUID) (SUUID, error) {
	token := generate()
	err := itor.TokenRepos.Save(Token{ID: token, UserID: userID})
	return token, err
}

// TODO:因为token来校验用户是否登录，放在token中实现  <15-11-20, nqq> //
//CheckIfLoggedin 校验用户是否登录
func (itor TokenInteractor) CheckIfLoggedin(tokenID SUUID) (Token, error) {
	if len(tokenID) == 0 {
		return errors.New("token不能为空")
	}
	token, err := itor.TokenRepos.Get(tokenID)
	return token, err
}

func GenerateToken() SUUID {
	return ""
}

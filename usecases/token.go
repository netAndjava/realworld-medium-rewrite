// Package usecases provides ...
package usecases

import "iohttps.com/live/realworld-medium-rewrite/domain"

type Token struct {
	ID     SUUID
	UserID domain.NUUID
}

type SUUID string

type TokenRepository interface {
	Save(t Token) error
}

type TokeInteractor struct {
	TokenRepos TokenRepository
}

//Login 判断用户登录
func (itor Token) Login(userID domain.NUUID, generate func() SUUID) (SUUID, error) {
	token := generate()
	err := itor.TokenRepos.Save(Token{ID: token, UserID: userID})
	return token, err
}

// TODO:因为token来校验用户是否登录，放在token中实现  <15-11-20, nqq> //
//CheckIfLoggedin 校验用户是否登录
func (itor UserInteractor) CheckIfLoggedin()

func GenerateToken() SUUID {
	return ""
}

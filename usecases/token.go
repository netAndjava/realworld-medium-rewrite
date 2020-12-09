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

func GenerateToken() SUUID {
	return ""
}

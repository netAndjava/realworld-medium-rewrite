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

// TODO: SUUID 这种难以理解的id请加上注释,要么就语义可理解的命名. <10-12-20, bantana> //
type SUUID string

type TokenRepository interface {
	Save(t Token) error
	Get(tokenID SUUID) (Token, error)
	Delete(tokenID SUUID) error
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

//IsLoggedin 检查用户是否登录
func (itor TokenInteractor) IsLoggedin(tokenID SUUID) (Token, error) {
	if len(tokenID) == 0 {
		return Token{}, errors.New("token不能为空")
	}
	token, err := itor.TokenRepos.Get(tokenID)
	return token, err
}

//Logout 用户退出登录
func (itor TokenInteractor) Logout(tokenID SUUID) error {
	if len(tokenID) == 0 {
		return errors.New("token不能为空")
	}
	return itor.TokenRepos.Delete(tokenID)
}

func GenerateToken() SUUID {
	return ""
}

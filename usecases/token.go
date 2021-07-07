// Package usecases provides ...
package usecases

import (
	"errors"
	"time"

	"iohttps.com/live/realworld-medium-rewrite/domain"
)

//Token .....
type Token struct {
	ID        SUUID
	UserID    domain.NUUID
	ExpiredAt int64 //过期时间
}

//SUUID ....
// TODO: SUUID 这种难以理解的id请加上注释,要么就语义可理解的命名. <10-12-20, bantana> //
type SUUID string

//TokenRepository token存储器
type TokenRepository interface {
	Save(t Token) error
	Get(tokenID SUUID) (Token, error)
	Delete(tokenID SUUID) error
}

//TokenInteractor token交互器
type TokenInteractor struct {
	TokenRepos TokenRepository
}

// uid, tokenService get a token banding uid,
//Login 判断用户登录
func (itor TokenInteractor) Generate(userID domain.NUUID, generate func() SUUID) (SUUID, error) {
	token := generate()
	err := itor.TokenRepos.Save(Token{ID: token, UserID: userID})
	return token, err
}

//IsLogedin 检查用户是否登录
func (itor TokenInteractor) IsLogedin(tokenID SUUID) (Token, error) {
	if len(tokenID) == 0 {
		return Token{}, errors.New("token不能为空")
	}
	token, err := itor.TokenRepos.Get(tokenID)
	if token.ExpiredAt < time.Now().Unix() {
		return Token{}, errors.New("登录已过期，请重新登录")
	}
	return token, err
}

//Logout 用户退出登录
func (itor TokenInteractor) Logout(tokenID SUUID) error {
	if len(tokenID) == 0 {
		return errors.New("token不能为空")
	}
	return itor.TokenRepos.Delete(tokenID)
}

//GenerateToken 生成token id
func GenerateToken() SUUID {
	return ""
}

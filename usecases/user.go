// Package usecases provides ...
package usecases

import "iohttps.com/live/realworld-medium-rewrite/domain"

//UserInteractor ....
type UserInteractor struct {
	UserRepo domain.UserRepository
}

//Token ...
type Token struct {
	ID     SUUID
	UserID domain.NUUID
}

type TokenRepository interface {
	CreateToken(t Token) (SUUID, error)
}

//SUUID string类型uuid
type SUUID string

//Register 用户注册
func (itor UserInteractor) Register(GenerateUUID func() domain.NUUID, user domain.User) (domain.NUUID, error) {
	return domain.NUUID(0), nil
}

//Login 用户登录
func (itor UserInteractor) Login(checkIdentity func(user domain.User) error, u domain.User) (SUUID, error) {
	return SUUID(""), nil
}

//CheckIdentityByEmail 通过email来校验身份
func (itor UserInteractor) CheckIdentityByEmail()

//GenerateToken 生成token
func (itor UserInteractor) GenerateToken()

//CheckIfLoggedin 校验用户是否登录
func (itor UserInteractor) CheckIfLoggedin()

//getUserByToken 通过token获取用户身份
func (itor UserInteractor) getUserByToken()

//Logout 用户退出登录
func (itor UserInteractor) Logout()

//GenerateVericationCode 生成验证码
func (itor UserInteractor) GenerateVericationCode()

//CheckIfVerCodeIsCorrect 判断用户验证码是否正确
func (itor UserInteractor) CheckIfVerCodeIsCorrect()

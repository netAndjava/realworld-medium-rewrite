// Package usecases provides ...
package usecases

import "iohttps.com/live/realworld-medium-rewrite/domain"

//UserInteractor ....
type UserInteractor struct {
	UserRepo domain.UserRepository
}

//Register 用户注册
func (itor UserInteractor) Register(GenerateUUID func() domain.NUUID, user domain.User) (domain.NUUID, error) {
	return domain.NUUID(0), nil
}

//CheckIdentityByEmail 通过email来校验身份
func (itor UserInteractor) CheckIdentityByEmail(u domain.User) error {
	return nil
}

//CheckIfLoggedin 校验用户是否登录
//// TODO:因为token来校验用户是否登录，放在token中实现  <15-11-20, nqq> //
// func (itor UserInteractor) CheckIfLoggedin()

//getUserByToken 通过token获取用户身份
//// TODO: 通过token来获取用户信息，放在token中实现 <15-11-20, nqq> //
// func (itor UserInteractor) getUserByToken()

//Logout 用户退出登录
//// TODO: 放在token中实现 <15-11-20, nqq> //
// func (itor UserInteractor) Logout()

//GenerateVericationCode 生成验证码
func (itor UserInteractor) GenerateVericationCode()

//CheckIfVerCodeIsCorrect 判断用户验证码是否正确
func (itor UserInteractor) CheckIfVerCodeIsCorrect()

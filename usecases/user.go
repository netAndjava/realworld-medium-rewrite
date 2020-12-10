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
func (itor UserInteractor) CheckIdentityByEmail(name, password string) (domain.User, error) {
	return domain.User{}, nil
}

func (itor UserInteractor) GetUserByPhone(phone domain.PhoneNumber) (domain.User, error) {
	// TODO:  <10-12-20, bantana> //
	return domain.User{}, nil
}

//GenerateVericationCode 生成验证码
func (itor UserInteractor) GenerateVericationCode()

//CheckIfVerCodeIsCorrect 判断用户验证码是否正确
func (itor UserInteractor) CheckIfVerCodeIsCorrect()

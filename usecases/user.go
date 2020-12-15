// Package usecases provides ...
package usecases

import (
	"errors"

	"iohttps.com/live/realworld-medium-rewrite/domain"
)

//UserInteractor ....
type UserInteractor struct {
	UserRepo domain.UserRepository
}

//Register 用户注册
func (itor UserInteractor) Register(generate func() domain.NUUID, user domain.User) (domain.NUUID, error) {
	if err := user.Check(); err != nil {
		return domain.NUUID(0), err
	}
	u, err := itor.UserRepo.GetByEmail(user.Email)
	if u.Email == user.Email {
		return nil, errors.New("该邮箱已注册过")
	}

	user.ID = generate()
	err := itor.UserRepo.Create(u)

	return user, err
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
func (itor UserInteractor) GenerateVericationCode() {
}

//CheckIfVerCodeIsCorrect 判断用户验证码是否正确
func (itor UserInteractor) CheckIfVerCodeIsCorrect() {
}

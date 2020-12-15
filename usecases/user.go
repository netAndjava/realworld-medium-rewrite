// Package usecases provides ...
package usecases

import (
	"errors"
	"strings"

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
	if len(strings.TrimSpace(name)) == 0 {
		return domain.User{}, errors.New("请输入用户名")
	}
	if len(strings.TrimSpace(password)) == 0 {
		return domain.User{}, errors.New("请输入密码")
	}
	user, err := itor.UserRepo.GetByEmail(name)
	if err != nil && err == domain.ErrNotFoundj {
		return domain.User{}, errors.New("该邮箱还未注册")
	}

	if err != nil {
		return domain.User{}, err
	}

	if user.Name == name && user.Password != password {
		return domain.User{}, errors.New("用户名密码不匹配")
	}

	return user, nil
}

func (itor UserInteractor) GetUserByPhone(phone domain.PhoneNumber) (domain.User, error) {
	if err := iphone.Check(); err != nil {
		return domain.User{}, err
	}
	user, err := itor.UserRepo.FindByPhone(phone)
	if err != nil && err == domain.ErrNotFound {
		return domain.User{}, errors.New("改用户还未注册")
	}

	return user, err
}

//GenerateVericationCode 生成验证码
func (itor UserInteractor) GenerateVericationCode() {
}

//CheckIfVerCodeIsCorrect 判断用户验证码是否正确
func (itor UserInteractor) CheckIfVerCodeIsCorrect() {
}

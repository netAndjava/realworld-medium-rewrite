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

//SUUID string类型uuid
type SUUID string

//Register 用户注册
func (itor UserInteractor) Register()

//Login 用户登录
func (itor UserInteractor) Login()

//CheckIdentityByEmail 通过email来校验身份
func (itor UserInteractor) CheckIdentityByEmail()

//CheckLoginStatus 校验用户是否登录
func (itor UserInteractor) CheckIfLoggedin()

//getUserByToken 通过token获取用户身份
func (itor UserInteractor) getUserByToken()

//Logout 用户退出登录
func (itor UserInteractor) Logout()

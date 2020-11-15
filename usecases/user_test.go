// Package usecases provides ...
package usecases

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"iohttps.com/live/realworld-medium-rewrite/domain"
)

var userItor UserInteractor

func TestRegister(t *testing.T) {
	//测试注册失败
	//1. 用户信息为空
	//2. 邮箱格式不正确
	//2. 邮箱已经被注册过
	u := domain.User{Email: "1040@qq.com"}
	// 测试注册成功
	_, err := userItor.Register(GenerateUUID, u)
	assert.Nil(err)
}

func TestLogin(t *testing.T) {
	//测试登录成功
	u := domain.User{Email: "1040qq.com", Password: "123456"}
	token, err := itor.Login(CheckIdentityByEmail, u)
	a := assert.New(t)
	a.True(a.Nil(err) && a.True(len(token) > 0))
}

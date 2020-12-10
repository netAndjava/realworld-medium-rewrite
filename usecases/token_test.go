// Package usecases provides ...
package usecases

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tokenItor TokenInteractor

//TestLogin 测试登录
func TestLogin(t *testing.T) {
	a := assert.New(t)
	_, err := tokenItor.Login(10, GenerateToken)
	a.Nil(err)
}

// TODO:  login, logout, logged in, logged out<10-12-20, bantana> //
// Noun: login, logout, loggedIn, loggedOut
// Verb: log in, log out, logged in, logged out
// Derived forms: logins, logged in, logging in, logs in
// 重新理解命名的原因:
// UI/UX中用户识别和用户体验中的影响
// 1. sign on, sign in, sign out
// 2. register, login, logout
// 3. register, sign in, sign out (best)

// TestCheckIfLoggedin tokenItor.CheckIfLogged 改为tokenItor.IsLoggedIn or tokenItor.IsOnline ?
func TestCheckIfLoggedin(t *testing.T) {
	a := assert.New(t)
	tokenID := ""
	_, err := tokenItor.IsLoggedin(SUUID(tokenID))
	a.NotNil(err)
	tokenID = "test"
	_, err := tokenItor.IsLoggedin(SUUID(tokenID))
	a.Nil(err)
}

func TestLoggout(t *testing.T) {
	a := assert.New(t)
	tokenID := ""
	err := tokenItor.Logout(tokenID)
	a.NotNil(err)
	tokenID = "test"
	err = tokenItor.Logout(tokenID)
	a.Nil(err)
}

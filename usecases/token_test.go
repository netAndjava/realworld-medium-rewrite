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

func TestCheckIfLoggedin(t *testing.T) {
	a := assert.New(t)
	tokenID = ""
	_, err := tokenItor.CheckIfLoggedin(tokenID)
	a.NotNil(err)
	tokenID = "test"
	_, err := tokenItor.CheckIfLoggedin(tokenID)
	a.Nil(err)
}

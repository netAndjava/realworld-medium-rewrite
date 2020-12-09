// Package usecases provides ...
package usecases

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tokenItor TokeInteractor

//TestLogin 测试登录
func TestLogin(t *testing.T) {
	a := assert.New(t)
	_, err := tokenItor.Login(10, GenerateToken)
	a.Nil(err)
}

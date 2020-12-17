// Package usecases provides ...
package usecases

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"iohttps.com/live/realworld-medium-rewrite/domain"
)

var commentItor CommentInteractor

func TestAdd(t *testing.T) {
	a := assert.New(t)

	c := domain.Comment{}
	//测试添加不成功
	//1. 添加的内容为空
	_, err := commentItor.Add(GenerateUUID, c, 1)
	a.NotNil(err)

	//添加成功
	c.Content = "test"
	_, err = commentItor.Add(GenerateUUID, c, 1)
	a.Nil(err)
}

func TestGetCommentsOfArticle(t *testing.T) {
	a := assert.New(t)
	_, err := commentItor.GetCommentsOfArticle(1)
	a.Nil(err)
}

func TestDrop(t *testing.T) {
	a := assert.New(t)
	//测试删除不成功
	//没有删除权限
	err := commentItor.Drop(10)
	a.NotNil(err)
	err = commentItor.Drop(1)
	a.Nil(err)
}

func TestDropByCreator(t *testing.T) {
	a := assert.New(t)
	//测试删除不成功
	//没有删除权限
	err := commentItor.DropByCreator(10, 0)
	a.NotNil(err)
	err = commentItor.DropByCreator(10, 1)
	a.Nil(err)
}

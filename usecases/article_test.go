// Package usecases provides ...
package usecases

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"iohttps.com/live/realworld-medium-rewrite/domain"
)

var itor ArticleInteractor

//TestEditDraftArticle 测试保存草稿
func TestEditDraftArticle(t *testing.T) {
	a := assert.New(t)

	//2. 更新草稿文章
	//2.1 用户没有更新权限，更新失败
	//2.2 更新成功
	article := domain.Article{}
	article.Title = "testja"
	article.Content = "test"
	err := itor.EditDraftArticle(article)
	a.Nil(err)

}

func TestWrite(t *testing.T) {
	a := assert.New(t)
	//1. 创建草稿态文章
	//1.1 创建的内容为空的文章，创建失败
	article := domain.Article{}
	_, err := itor.Wrtie(GenerateUUID, article)
	a.NotNil(err)

	//创建成功
	article.Title = "time"
	ID, err := itor.Wrtie(GenerateUUID, article)
	a.True(a.Nil(err) && a.IsType(new(domain.NUUID), &ID, nil))
}

func TestPublish(t *testing.T) {
	a := assert.New(t)
	//1.测试发布失败
	//1.1发布的文章不存在
	art := domain.Article{}
	err := itor.Publish(art.ID, 0)
	a.NotNil(err)
	//1.2 作者没有权限
	//1.3 文章的标题或者内容为空
	//测试发布成功

}

func TestViewDraftArticles(t *testing.T) {
	a := assert.New(t)
	//测试成功
	arts, err := itor.ViewDraftArticles(domain.NUUID(1))
	a.True(a.Nil(err) && a.IsType(arts, []domain.Article{}, nil))
}

//ViewPublicArticles 获取作者已发布文章
func ViewPublicArticles(t *testing.T) {
	a := assert.New(t)
	//测试成功
	arts, err := itor.ViewPublicArticles(domain.NUUID(1))
	a.True(a.Nil(err) && a.IsType(arts, []domain.Article{}, nil))
}

//TestView 获取文章详情
func TestView(t *testing.T) {
	a := assert.New(t)
	//测试失败
	//1.1获取的文章不存在
	_, err := itor.View(domain.NUUID(0))
	a.NotNil(err)
	_, err = itor.View(domain.NUUID(1))
	a.Nil(err)
}

func TestViewRecentArticles(t *testing.T) {
	a := assert.New(t)
	arts, err := itor.ViewRecentArticles()
	a.True(a.Nil(err) && a.IsType(arts, []domain.Article{}, nil))
}

func TestViewDraftOfPublicArticle(t *testing.T) {
	a := assert.New(t)
	//测试失败
	//查找的文章不存在
	_, err := itor.ViewDraftOfPublicArticle(0)
	a.NotNil(err)
	_, err = itor.ViewDraftOfPublicArticle(1)
	a.Nil(err)
}

//TestEditPublicArticle 测试保存已发布文章草稿
func TestEditPublicArticle(t *testing.T) {
	a := assert.New(t)

	//1.测试失败
	//1.1草稿不存在
	//1.2用户没有权限

	art := domain.Article{}
	err := itor.EditPublicArticle(art)
	a.NotNil(err)

	art.ID = 1
	err = itor.EditPublicArticle(art)
	a.Nil(err)

}

func TestRepublish(t *testing.T) {
	a := assert.New(t)
	//测试发布失败
	// 1. 发布为文章不存在
	art := domain.Article{}
	err := itor.Republish(art)
	a.NotNil(err)
	//2.发布的人不是文章作者
	itor.Republish(art)
	a.NotNil(err)

	//测试发布成功
	err = itor.Republish(art)
	a.Nil(err)
}

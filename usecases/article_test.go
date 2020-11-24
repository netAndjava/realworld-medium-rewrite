// Package usecases provides ...
package usecases

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"iohttps.com/live/realworld-medium-rewrite/domain"
)

var itor ArticleInteractor

//TestSaveDraft 测试保存草稿
func TestSaveDraft(t *testing.T) {
	a := assert.New(t)

	//2. 更新草稿文章
	//2.1 用户没有更新权限，更新失败
	//2.2 更新成功
	article := domain.Article{}
	article.Title = "testja"
	article.Content = "test"
	err := itor.SaveDraft(article, 10)
	a.Nil(err)

}

func TestCreateDraft(t *testing.T) {
	a := assert.New(t)
	//1. 创建草稿态文章
	//1.1 创建的内容为空的文章，创建失败
	article := domain.Article{}
	_, err := itor.CreateDraft(GenerateUUID, article)
	a.NotNil(err)

	//创建成功
	article.Title = "time"
	ID, err := itor.CreateDraft(GenerateUUID, article)
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

func TestGetAuthorDrafts(t *testing.T) {
	a := assert.New(t)
	//测试成功
	arts, err := itor.GetAuthorDrafts(domain.NUUID(1))
	a.True(a.Nil(err) && a.IsType(arts, []domain.Article{}, nil))
}

//TestGetAuthorPublicArticles 获取作者已发布文章
func TestGetAuthorPublicArticles(t *testing.T) {
	a := assert.New(t)
	//测试成功
	arts, err := itor.GetAuthorPublicArticles(domain.NUUID(1))
	a.True(a.Nil(err) && a.IsType(arts, []domain.Article{}, nil))
}

//TestGetArticle 获取文章详情
func TestGetArticle(t *testing.T) {
	a := assert.New(t)
	//测试失败
	//1.1获取的文章不存在
	_, err := itor.GetArticle(domain.NUUID(0))
	a.NotNil(err)
	_, err = itor.GetArticle(domain.NUUID(1))
	a.Nil(err)
}

func TestGetAllPublicArticles(t *testing.T) {
	a := assert.New(t)
	arts, err := itor.GetAllPublicArticles()
	a.True(a.Nil(err) && a.IsType(arts, []domain.Article{}, nil))
}

func TestGetPublicArticleDraft(t *testing.T) {
	a := assert.New(t)
	//测试失败
	//查找的文章不存在
	_, err := itor.GetPublicArticleDraft(0)
	a.NotNil(err)
	_, err = itor.GetPublicArticleDraft(1)
	a.Nil(err)
}

//TestSavePublicArticleDraft 测试保存已发布文章草稿
func TestSavePublicArticleDraft(t *testing.T) {
	a := assert.New(t)

	//1.测试失败
	//1.1草稿不存
	//1.2用户没有权限

	art := domain.Article{}
	err := itor.SavePublicArticleDraft(art, domain.NUUID(1))
	a.NotNil(err)

	art.ID = 1
	err = itor.SavePublicArticleDraft(art, domain.NUUID(1))
	a.Nil(err)

}

func TestCreatePublicArticleDraft(t *testing.T) {
	//测试失败
	//1.1已发布文章不存在
	//1.2 没有修改已发布文章权限
}

func TestPublishPublicArticleDraft(t *testing.T) {
	a := assert.New(t)
	//测试发布失败
	// 1. 发布为文章不存在
	err := itor.PublishPublicArticleDraft(10000, 1)
	a.NotNil(err)
	//2.发布的人不是文章作者
	itor.PublishPublicArticleDraft(3, 1)
	a.NotNil(err)

	//测试发布成功
	err = itor.PublishPublicArticleDraft(3, 3)
	a.Nil(err)
}

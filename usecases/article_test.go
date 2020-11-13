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
	//1. 创建草稿态文章
	//1.1 保存空内容失败
	article := domain.Article{}
	_, err := itor.SaveDraft(GenerateUUID, article)
	a.NotNil(err)

	//1.2 创建草稿文章成功
	article.Title = "test"
	ID, err := itor.SaveDraft(GenerateUUID, article)
	a.True(a.Nil(err) && a.IsType(new(domain.NUUID), &ID, nil))

	//2. 更新草稿文章
	article.Title = "testja"
	article.Content = "test"
	ID, err = itor.SaveDraft(GenerateUUID, article)
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

	//测试失败
	//已发布文章不存在
	art := domain.Article{}
	err := itor.SavePublicArticleDraft(art)
	a.NotNil(err)

	art.ID = 1
	err = itor.SavePublicArticleDraft(art)
	a.Nil(err)

}

func TestPublishPublicArticleDraft(t *testing.T) {
	a := assert.New(t)
	//测试发布失败
	art := domain.Article{}
	//1. 发布的内容为空
	err := itor.PublishPublicArticleDraft(art, 1)
	a.NotNil(err)
	//2.发布的文章不存在
	art.Title = "test"
	art.Content = "test"
	itor.PublishPublicArticleDraft(a, 1)
	a.NotNil(err)
	//3.发布的人不是文章作者
	art.Author.ID = 3
	itor.PublishPublicArticleDraft(a, 1)
	a.NotNil(err)

	//测试发布成功
	err = itor.PublishPublicArticleDraft(a, 3)
	a.Nil(err)
}

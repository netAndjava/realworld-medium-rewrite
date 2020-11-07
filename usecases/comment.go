// Package usecases provides ...
package usecases

import "iohttps.com/live/realworld-medium-rewrite/domain"

//CommentInteractor ......
type CommentInteractor struct {
	CommentRepos domain.CommentRepository
}

//Add 添加评论
func (itor CommentInteractor) Add()

//DropByCreator 评论长作者删除评论
func (itor CommentInteractor) DropByCreator()

//GetCommentsOfArticle 获取文章的评论列表
func (itor CommentInteractor) GetCommentsOfArticle()

//DropByArticleAuthor 文章作者删除针对文章的评论
func (itor CommentInteractor) DropByArticleAuthor()

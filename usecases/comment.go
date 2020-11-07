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

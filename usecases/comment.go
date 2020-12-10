// Package usecases provides ...
package usecases

import "iohttps.com/live/realworld-medium-rewrite/domain"

type CommentInteractor struct {
	CommentRepos domain.CommentRepository
}

//CommentInteractor ......
type CommentInteractor struct {
	CommentRepos domain.CommentRepository
}

//Add 添加评论
func (itor CommentInteractor) Add(generate func() domain.NUUID, c domain.Comment, userID domain.NUUID) (domain.NUUID, error) {
	if err := c.Check(); err != nil {
		return domain.NUUID{}, err
	}
	c.Creator.ID = userID
	c.ID = generate()
	err := itor.CommentRepos.Add(c)
	return c.ID, err
}

//DropByCreator 评论文作者删除评论
func (itor CommentInteractor) DropByCreator()

//GetCommentsOfArticle 获取文章的评论列表
func (itor CommentInteractor) GetCommentsOfArticle()

//DropByArticleAuthor 文章作者删除针对文章的评论
func (itor CommentInteractor) DropByArticleAuthor()

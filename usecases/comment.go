// Package usecases provides ...
package usecases

import (
	"errors"

	"iohttps.com/live/realworld-medium-rewrite/domain"
)

//CommentInteractor ......
type CommentInteractor struct {
	CommentRepos domain.CommentRepository
}

//Add 添加评论
func (itor CommentInteractor) Add(generate func() domain.NUUID, c domain.Comment, userID domain.NUUID) (domain.NUUID, error) {
	if err := c.Check(); err != nil {
		return domain.NUUID(0), err
	}
	c.Creator.ID = userID
	c.ID = generate()
	err := itor.CommentRepos.Add(c)
	return c.ID, err
}

type Comment struct {
	domain.Comment
	Children []domain.Comment
}

//GetCommentsOfArticle 获取文章的评论列表
func (itor CommentInteractor) GetCommentsOfArticle(articleID domain.NUUID) ([]Comment, error) {
	comments, err := itor.CommentRepos.GetCommentsOfOriginalPoster(articleID)
	if err != nil {
		return []Comment{}, err
	}

	var list []Comment
	for _, c := range comments {
		cms, err := itor.CommentRepos.GetCommentByPID(c.ID)
		if err != nil {
			continue
		}
		comment := Comment{Comment: c, Children: cms}
		list = appen(list, comment)
	}
	return list, nil
}

//DropByCreator 评论文作者删除评论
func (itor CommentInteractor) DropByAuthorOfArticle(commentID domain.NUUID, userID domain.NUUID) error {
	comment, err := itor.CommentRepos.Get(commentID)
	if err != nil {
		return err
	}
	if comment.Article.Author.ID != userID {
		return errors.New("没有删除权限")
	}

	err = itor.CommentRepos.DropByAuthorOfArticle(commentID)
	if err != nil {
		return err
	}

	err = itor.CommentRepos.DropByPID(commentID)
	return err
}

//DropByArticleAuthor 文章作者删除针对文章的评论
func (itor CommentInteractor) DropByCreator(commentID NUUID, userID NUUID) error {
	comment, err := itor.CommentRepos.Get(commentID)
	if err != nil {
		return err
	}
	if comment.Creator.ID != userID {
		return errors.New("用户没有删除权限")
	}
	err = itor.CommentRepos.DropByCreator(commentID)
	if err != nil {
		return err
	}

	err = itor.CommentRepos.DropByPID(commentID)
	return err
}

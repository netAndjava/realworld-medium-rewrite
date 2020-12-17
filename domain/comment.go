// Package domain provides ...
package domain

import "errors"

//Comment 文章评论
// 文章在文章下面直接平路是楼主评论
// 在楼主中回复评论，回复的是楼主评论
// 如果有@评论作者,则需要知道回复的评论人
type Comment struct {
	ID        NUUID
	PID       NUUID  //父ID，用来知道回复的评论
	ArticleID NUUID  //评论的文章
	Content   string //Content 评论的内容
	Creator   NUUID  //评论的创建者
}

func (c Comment) Check() error {
	if len(c.Content) == 0 {
		return errors.New("评论内容不能为空")
	}
	return nil
}

//CommentRepository 评论行为
type CommentRepository interface {
	Add(c Comment) error
	Get(ID NUUID) (Comment, error)
	GetCommentByPID(PID NUUID) ([]Comment, error)
	GetCommentsByArticleID(articleID NUUID) ([]Comment, error)
	Drop(ID NUUID) error
	DropByPID(PID NUUID) error
}

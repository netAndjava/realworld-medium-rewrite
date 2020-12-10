// Package domain provides ...
package domain

//Comment 文章评论
// 文章在文章下面直接平路是楼主评论
// 在楼主中回复评论，回复的是楼主评论，需要知道回复哪一楼
// 如果有@评论作者,则需要知道回复的评论人
type Comment struct {
	ID        NUUID
	PID       NUUID   //父ID，用来知道回复的评论
	Article   Article //评论的文章
	ReplyUser User    //回复的用户ID
	Content   string  //Content 评论的内容
	Creator   User    //评论的创建者
}

//CommentRepository 评论行为
type CommentRepository interface {
	Add(c Comment) error
	GetCommentByPID(PID NUUID) ([]Comment, error)
	GetCommentsOfOriginalPoster(articleID NUUID) ([]Comment, error)
}

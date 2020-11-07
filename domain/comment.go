// Package domain provides ...
package domain

//Comment 文章评论
type Comment struct {
	ID        NUUID
	PID       NUUID  //父ID，用来知道回复的评论
	ArticleID NUUID  //评论的文章
	Content   string //Content 评论的内容
	Creator   User   //评论的创建者
}

//CommentRepository 评论行为
type CommentRepository interface {
	Create()
	Delete()
}

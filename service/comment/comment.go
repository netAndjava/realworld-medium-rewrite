// Package comment provides ...
package comment

import (
	"log"

	"iohttps.com/live/realworld-medium-rewrite/interfaces"
	pb "iohttps.com/live/realworld-medium-rewrite/service/api"
	"iohttps.com/live/realworld-medium-rewrite/usecases"
)

type commentServer struct {
	pb.UnimplementedCommentServer
	commentItor usecases.CommentInteractor
}

func Start(port int) {
	handler, err := mysql.NewMysqlHanelr("root@/real_world_medium?charset=utf8")
	if err != nil {
		log.Fatal("connect database err:", err)
	}
	repo := interfaces.NewCommentRepo(handler)

}

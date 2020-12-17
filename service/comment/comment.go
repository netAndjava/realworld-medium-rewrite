// Package comment provides ...
package comment

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"iohttps.com/live/realworld-medium-rewrite/domain"
	"iohttps.com/live/realworld-medium-rewrite/infrastructure/database"
	"iohttps.com/live/realworld-medium-rewrite/interfaces"
	pb "iohttps.com/live/realworld-medium-rewrite/service/api"
	"iohttps.com/live/realworld-medium-rewrite/usecases"
)

type commentServer struct {
	pb.UnimplementedCommentServer
	commentItor usecases.CommentInteractor
}

// {
// 	handler, err := mysql.NewMysqlHandler("root@/real_world_medium?charset=utf8")
// 	if err != nil {
// 		log.Fatal("connect database err:", err)
// 	}
//
// }
// // started is start service for comment
// func  started(ip:port, dbhandle){
// 	// 1.  init environment variables listen{ip:port}
//
// 	conn,err := net.Listen("tcp", fmt.Sprintf(":%d", port))
// 	// 1.1 db
// 	repo = interfaces.NewCommentRepo(dbhandler)
//
// }

//Start ....
func Start(address string, handler database.DbHandler) {

	repo := interfaces.NewCommentRepo(handler)

	itor := usecases.CommentInteractor{CommentRepos: repo}

	conn, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("listen to address:%s err:%v\n", address, err)
	}
	log.Println("start server on address:", address)

	server := grpc.NewServer()
	pb.RegisterCommentServer(server, &commentServer{commentItor: itor})
	server.Serve(conn)
}

//Add 添加评论
func (server commentServer) Add(ctx context.Context, c *pb.Comment) (*pb.AddRep, error) {
	ID, err := server.commentItor.Add(usecases.GenerateUUID, domain.Comment{PID: domain.NUUID(c.Pid), ArticleID: domain.NUUID(c.ArticleID), Content: c.Content, Creator: domain.NUUID(c.UserID)}, domain.NUUID(c.UserID))
	if err != nil {
		return nil, err
	}
	return &pb.AddRep{Id: int64(ID)}, nil
}

func (server commentServer) GetCommentsOfArticle(ctx context.Context, req *pb.GetCommentsOfArticleReq) (*pb.GetCommentsOfArticleRep, error) {
	comments, err := server.commentItor.GetCommentsOfArticle(domain.NUUID(req.ArticleID))
	if err != nil {
		return nil, err
	}

	cms := make([]*pb.Comment, len(comments))
	for i, c := range comments {
		cms[i] = &pb.Comment{Id: int64(c.ID), Pid: int64(c.PID), ArticleID: int64(c.ArticleID), UserID: int64(c.Creator), Content: c.Content}
	}
	return &pb.GetCommentsOfArticleRep{Comments: cms}, nil
}

func (server commentServer) Drop(ctx context.Context, req *pb.DropReq) (*pb.DropRep, error) {
	err := server.commentItor.Drop(domain.NUUID(req.Id))
	if err != nil {
		return nil, err
	}
	return &pb.DropRep{}, nil
}

func (server commentServer) DropByCreator(ctx context.Context, req *pb.DropByCreatorReq) (*pb.DropByCreatorRep, error) {
	err := server.commentItor.DropByCreator(domain.NUUID(req.Id), domain.NUUID(req.UserID))
	return &pb.DropByCreatorRep{}, err

}

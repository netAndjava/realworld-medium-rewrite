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
	pb.UnimplementedCommentServiceServer
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
	pb.RegisterCommentServiceServer(server, &commentServer{commentItor: itor})
	server.Serve(conn)
}

//Add 添加评论
func (server commentServer) Add(ctx context.Context, in *pb.AddRequest) (*pb.AddResponse, error) {
	ID, err := server.commentItor.Add(usecases.GenerateUUID, domain.Comment{PID: domain.NUUID(in.Comment.Pid), ArticleID: domain.NUUID(in.Comment.ArticleId), Content: in.Comment.Content, Creator: domain.NUUID(in.Comment.UserId)})
	if err != nil {
		return nil, err
	}
	return &pb.AddResponse{Id: int64(ID)}, nil
}

func (server commentServer) ViewComments(ctx context.Context, req *pb.ViewCommentsRequest) (*pb.ViewCommentsResponse, error) {
	comments, err := server.commentItor.GetCommentsOfArticle(domain.NUUID(req.ArticleId))
	if err != nil {
		return nil, err
	}

	cms := make([]*pb.Comment, len(comments))
	for i, c := range comments {
		cms[i] = &pb.Comment{Id: int64(c.ID), Pid: int64(c.PID), ArticleId: int64(c.ArticleID), UserId: int64(c.Creator), Content: c.Content}
	}
	return &pb.ViewCommentsResponse{Comments: cms}, nil
}

func (server commentServer) Drop(ctx context.Context, req *pb.DropRequest) (*pb.DropResponse, error) {
	err := server.commentItor.Drop(domain.NUUID(req.Id))
	if err != nil {
		return nil, err
	}
	return &pb.DropResponse{}, nil
}

func (server commentServer) DropByCreator(ctx context.Context, req *pb.DropByCreatorRequest) (*pb.DropByCreatorResponse, error) {
	err := server.commentItor.DropByCreator(domain.NUUID(req.Id), domain.NUUID(req.UserId))
	return &pb.DropByCreatorResponse{}, err
}

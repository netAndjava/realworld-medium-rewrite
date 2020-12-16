// Package comment provides ...
package comment

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
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
func Start(port int) {

	// handler, err := mysql.NewMysqlHandler("root@/real_world_medium?charset=utf8")
	// if err != nil {
	// 	log.Fatal("connect database err:", err)
	// }
	repo := interfaces.NewCommentRepo(handler)

	conn, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("listen to port:%d err:%v\n", port, err)
	}
	server := grpc.NewServer()
	pb.RegisterService(server, &commentServer{usecases.CommentInteractor{repo}})
	server.Serve(conn)

}

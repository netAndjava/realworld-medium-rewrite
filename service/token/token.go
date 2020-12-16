// Package token provides ...
package token

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"iohttps.com/live/realworld-medium-rewrite/infrastructure/database/mysql"
	"iohttps.com/live/realworld-medium-rewrite/interfaces"
	pb "iohttps.com/live/realworld-medium-rewrite/service/api"
	"iohttps.com/live/realworld-medium-rewrite/usecases"
)

type tokenServer struct {
	pb.UnimplementedTokenServer
	tokenItor usecases.TokenInteractor
}

func Start(port int) {
	handler, err := mysql.NewMysqlHandler("root@/real_world_medium?charset=utf8")
	if err != nil {
		log.Fatal("connect db err:", err)
	}
	tokenRepo := interfaces.NewTokenRepo(handler)
	tokenItor := usecases.TokenInteractor{TokenRepos: tokenRepo}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("listen port:%d err:%v", port, err)
	}

	s := grpc.NewServer()
	pb.RegisterTokenServer(s, &tokenServer{tokenItor: tokenItor})
	s.Serve(lis)
}

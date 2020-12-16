// Package token provides ...
package token

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"iohttps.com/live/realworld-medium-rewrite/domain"
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
	log.Println("listen to port:", port)

	s := grpc.NewServer()
	pb.RegisterTokenServer(s, &tokenServer{tokenItor: tokenItor})
	s.Serve(lis)
}

func (server *tokenServer) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginRep, error) {
	tokenID, err := server.tokenItor.Login(domain.NUUID(req.UserID), usecases.GenerateToken)
	if err != nil {
		return nil, err
	}
	return &pb.LoginRep{TokenID: string(tokenID)}, nil
}

func (server *tokenServer) IsLoggedin(ctx context.Context, req *pb.IsLoggedinReq) (*pb.Token, error) {
	token, err := server.tokenItor.IsLoggedin(req.TokenID)
	if err != nil {
		return nil, err
	}
	return &pb.Token{TokenID: string(token.ID), UserID: int64(token.UserID)}

}

func (server *tokenServer) Logout(ctx context.Context, req *pb.LogoutReq) (*pb.LoginRep, error) {
	err := server.tokenItor.Logout(usecases.SUUID(req.TokenID))
	if err != nil {
		return nil, err
	}
	return &pb.LogoutRep{}, nil
}

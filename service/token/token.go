// Package token provides ...
package token

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

type tokenServer struct {
	pb.UnimplementedTokenServer
	tokenItor usecases.TokenInteractor
}

//Start .....
func Start(address string, handler database.DbHandler) {
	tokenRepo := interfaces.NewTokenRepo(handler)

	tokenItor := usecases.TokenInteractor{TokenRepos: tokenRepo}

	conn, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("listen address:%s err:%v", conn, err)
	}
	log.Println("listen to address:", address)

	s := grpc.NewServer()
	pb.RegisterTokenServer(s, &tokenServer{tokenItor: tokenItor})
	s.Serve(conn)
}

func (server *tokenServer) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginRep, error) {
	tokenID, err := server.tokenItor.Login(domain.NUUID(req.UserID), usecases.GenerateToken)
	if err != nil {
		return nil, err
	}
	return &pb.LoginRep{TokenID: string(tokenID)}, nil
}

func (server *tokenServer) IsLoggedin(ctx context.Context, req *pb.IsLoggedinReq) (*pb.Token, error) {
	token, err := server.tokenItor.IsLoggedin(usecases.SUUID(req.TokenID))
	if err != nil {
		return nil, err
	}
	return &pb.Token{TokenID: string(token.ID), UserID: int64(token.UserID)}, nil

}

func (server *tokenServer) Logout(ctx context.Context, req *pb.LogoutReq) (*pb.LogoutRep, error) {
	err := server.tokenItor.Logout(usecases.SUUID(req.TokenID))
	if err != nil {
		return nil, err
	}
	return &pb.LogoutRep{}, nil
}

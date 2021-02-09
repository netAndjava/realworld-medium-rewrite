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
	pb.UnimplementedTokenServiceServer
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
	pb.RegisterTokenServiceServer(s, &tokenServer{tokenItor: tokenItor})
	s.Serve(conn)
}

func (server *tokenServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	tokenID, err := server.tokenItor.Login(domain.NUUID(req.UserId), usecases.GenerateToken)
	return &pb.LoginResponse{TokenId: string(tokenID)}, err
}

func (server *tokenServer) IsLoggedin(ctx context.Context, req *pb.IsLoggedinRequest) (*pb.IsLoggedinResponse, error) {
	token, err := server.tokenItor.IsLoggedin(usecases.SUUID(req.TokenId))
	if err != nil {
		return nil, err
	}
	return &pb.IsLoggedinResponse{Token: &pb.Token{TokenId: string(token.ID), UserId: int64(token.UserID)}}, nil

}

func (server *tokenServer) Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	err := server.tokenItor.Logout(usecases.SUUID(req.TokenId))
	if err != nil {
		return nil, err
	}
	return &pb.LogoutResponse{}, nil
}

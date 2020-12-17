// Package user provides ...
package user

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

type userServer struct {
	pb.UserServer
	userItor usecases.UserInteractor
}

//Start ....
func Start(address string, handler database.DbHandler) {
	repo := interfaces.NewUserRepo(handler)

	userItor := usecases.UserInteractor{UserRepo: repo}

	conn, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("listen to address:%s err:%v\n", address, err)
	}
	log.Println("start server on address:", address)

	server := grpc.NewServer()
	pb.RegisterUserServer(server, &userServer{userItor: userItor})
	server.Serve(conn)
}

func (server *userServer) Register(ctx context.Context, user *pb.User) (*pb.RegisterRep, error) {
	ID, err := server.userItor.Register(usecases.GenerateUUID, domain.User{Name: user.Name, Email: domain.Email(user.Email), Password: user.Password})
	if err != nil {
		return nil, err
	}
	return &pb.RegisterRep{Id: int64(ID)}, nil
}

func (server *userServer) CheckIdentityByEmail(ctx context.Context, req *pb.CheckIdentityByEmailReq) (*pb.User, error) {
	user, err := server.userItor.CheckIdentityByEmail(domain.Email(req.Email), req.Password)
	if err != nil {
		return nil, err
	}
	return &pb.User{Id: int64(user.ID)}, nil

}

func (server *userServer) GetUserByPhone(ctx context.Context, req *pb.GetUserByPhoneReq) (*pb.User, error) {
	user, err := server.userItor.GetUserByPhone(domain.PhoneNumber(req.Phone))
	if err != nil {
		return nil, err
	}
	return &pb.User{Id: int64(user.ID), Phone: string(user.Phone)}, nil

}

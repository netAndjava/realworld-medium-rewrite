// Package user provides ...
package user

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

type userServer struct {
	pb.UserServer
	userItor usecases.UserInteractor
}

func Start(port int) {
	handler, err := mysql.NewMysqlHandler("root@/real_world_medium?charset=utf8") // TODO: dataSourceName 作为一个value在在代码里面到处都是,变更一下,你要到处去改吗?  <17-12-20, bantana> //
	if err != nil {
		log.Fatal("connect db err:", err)
	}

	userItor := usecases.UserInteractor{UserRepo: interfaces.NewUserRepo(handler)}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("listen to port:%d err:%v\n", port, err)
	}
	log.Println("start server on port:", port)

	s := grpc.NewServer()
	pb.RegisterUserServer(s, &userServer{userItor: userItor})
	s.Serve(lis)
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

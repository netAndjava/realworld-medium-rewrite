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
	pb.UnimplementedUserServiceServer
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
	pb.RegisterUserServiceServer(server, &userServer{userItor: userItor})
	server.Serve(conn)
}

func (server *userServer) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	ID, err := server.userItor.Register(usecases.GenerateUUID, domain.User{Name: in.User.Name, Email: domain.Email(in.User.Email), Password: in.User.Password})
	return &pb.RegisterResponse{Id: int64(ID)}, err
}

func (server *userServer) LoginCheckByEmail(ctx context.Context, req *pb.LoginCheckByEmailRequest) (*pb.LoginCheckByEmailResponse, error) {
	user, err := server.userItor.CheckIdentityByEmail(domain.Email(req.Email), req.Password)
	return &pb.LoginCheckByEmailResponse{User: &pb.User{Id: int64(user.ID)}}, err

}

func (server *userServer) GetUserByPhone(ctx context.Context, req *pb.GetUserByPhoneRequest) (*pb.GetUserByPhoneResponse, error) {
	user, err := server.userItor.GetUserByPhone(domain.PhoneNumber(req.Phone))
	if err != nil {
		return nil, err
	}
	return &pb.GetUserByPhoneResponse{User: &pb.User{Id: int64(user.ID), Phone: string(user.Phone)}}, err

}

func (server *userServer) GetUserById(ctx context.Context, req *pb.GetUserByIdRequest) (*pb.GetUserByIdResponse, error) {

	user, err := server.userItor.GetUserByID(domain.NUUID(req.Id))
	return &pb.GetUserByIdResponse{User: &pb.User{Id: int64(user.ID), Name: user.Name, Email: string(user.Email), Phone: string(user.Phone)}}, err
}

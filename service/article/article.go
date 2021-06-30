// Package article provides ...
package article

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

type articleServer struct {
	pb.UnimplementedArticleServiceServer
	artInteractor usecases.ArticleInteractor
}

//Start ....
func Start(address string, handler database.DbHandler) {
	//init
	//1.1 init enviroment
	//1.2 init db
	artRepo := interfaces.NewArticleRepo(handler)

	//1.3 create interactor for usecases
	artItor := usecases.ArticleInteractor{ArticleRepo: artRepo}

	//1.3 start grpc server
	conn, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("listen adderss:%s err:%v\n", address, err)
	}
	log.Println("start server on address:", address)
	server := grpc.NewServer()

	pb.RegisterArticleServiceServer(server, &articleServer{artInteractor: artItor})
	server.Serve(conn)
}

func (server *articleServer) Write(ctxt context.Context, in *pb.WriteRequest) (*pb.WriteResponse, error) {
	article := domain.Article{ID: domain.NUUID(in.Article.Id), Title: in.Article.Title, Content: in.Article.Content, Status: domain.Draft, AuthorID: domain.NUUID(in.Article.AuthorId)}
	//1.创建文章
	if in.Article.Id == 0 {
		ID, err := server.artInteractor.Write(usecases.GenerateUUID, article)
		if err != nil {
			return nil, err
		}

		return &pb.WriteResponse{Id: int64(ID)}, err
	}

	if in.Article.Id != 0 && domain.PublicStatus(in.Article.Status) == domain.Draft {
		//2.保存编辑的文章
		err := server.artInteractor.EditDraftArticle(article)
		if err != nil {
			return nil, err
		}
	}

	//3.保存草稿文章
	if in.Article.Id != 0 && domain.PublicStatus(in.Article.Status) == domain.Public {
		err := server.artInteractor.EditPublicArticle(article)
		if err != nil {
			return nil, err
		}
	}

	// go on,按理你应该在接口修改的地方直接使用重构,原先定义丢失后就不能使用重构了,你只能search and replace
	return &pb.WriteResponse{Id: in.Article.Id}, nil
}

func (server *articleServer) ViewDraftedArticles(ctxt context.Context, req *pb.ViewDraftedArticlesRequest) (*pb.ViewDraftedArticlesResponse, error) {
	arts, err := server.artInteractor.ViewDraftArticles(domain.NUUID(req.UserId))
	if err != nil {
		return nil, err
	}

	return &pb.ViewDraftedArticlesResponse{Articles: ConvertArticles(arts)}, err
}

func (server *articleServer) View(ctxt context.Context, in *pb.ViewRequest) (*pb.ViewResponse, error) {
	art, err := server.artInteractor.View(domain.NUUID(in.Id))
	if err != nil {
		return nil, err
	}
	return &pb.ViewResponse{Article: &pb.Article{Id: int64(art.ID), Title: art.Title, Content: art.Content, Status: int32(art.Status), AuthorId: int64(art.AuthorID)}}, err

}

func (server *articleServer) Publish(ctxt context.Context, req *pb.PublishRequest) (*pb.PublishResponse, error) {
	server.Write(ctxt, &pb.WriteRequest{Article: req.Article})
	err := server.artInteractor.Publish(domain.NUUID(req.Article.Id), domain.NUUID(req.Article.AuthorId))
	return nil, err
}

func (server *articleServer) ViewPublishedArticles(ctx context.Context, req *pb.ViewPublicArticlesRequest) (*pb.ViewPublicArticlesResponse, error) {
	arts, err := server.artInteractor.ViewPublicArticles(domain.NUUID(req.UserId))

	if err != nil {
		return nil, err
	}

	return &pb.ViewPublicArticlesResponse{Articles: ConvertArticles(arts)}, err

}

func (server *articleServer) ViewDraftOfPublicArticle(ctx context.Context, req *pb.ViewDraftOfPublicArticleRequest) (*pb.ViewDraftOfPublicArticleResponse, error) {
	art, err := server.artInteractor.ViewDraftOfPublicArticle(domain.NUUID(req.Id))
	if err != nil {
		return nil, err
	}
	return &pb.ViewDraftOfPublicArticleResponse{Article: &pb.Article{Id: int64(art.ID), Title: art.Title, Content: art.Content, Status: int32(art.Status), AuthorId: int64(art.AuthorID)}}, err

}

func (server articleServer) Republish(ctx context.Context, req *pb.RepublishRequest) (*pb.RepublishResponse, error) {
	err := server.artInteractor.Republish(domain.Article{ID: domain.NUUID(req.Article.Id), Title: req.Article.Title, Content: req.Article.Content, AuthorID: domain.NUUID(req.Article.AuthorId)})
	return &pb.RepublishResponse{}, err
}

func (server articleServer) Drop(ctx context.Context, req *pb.DropArticleRequest) (*pb.DropArticleResponse, error) {
	return &pb.DropArticleResponse{}, server.artInteractor.Drop(domain.NUUID(req.GetArticleId()), domain.NUUID(req.UserId))
}

//ConvertArticles .....
func ConvertArticles(arts []domain.Article) []*pb.Article {

	articles := make([]*pb.Article, len(arts))
	for i, art := range arts {
		articles[i] = &pb.Article{Id: int64(art.ID), Title: art.Title, Content: art.Content, Status: int32(art.Status), AuthorId: int64(art.AuthorID)}
	}
	return articles
}

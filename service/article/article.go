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

func (server *articleServer) Save(ctxt context.Context, in *pb.SaveRequest) (*pb.SaveResponse, error) {
	article := domain.Article{ID: domain.NUUID(in.Article.Id), Title: in.Article.Title, Content: in.Article.Content, Status: domain.Draft, AuthorID: domain.NUUID(in.Article.AuthorId)}
	//1.创建文章
	if in.Article.Id == 0 {
		ID, err := server.artInteractor.CreateDraft(usecases.GenerateUUID, article)
		if err != nil {
			return nil, err
		}

		return &pb.SaveResponse{Id: int64(ID)}, err
	}

	if in.Article.Id != 0 && domain.PublicStatus(in.Article.Status) == domain.Draft {
		//2.保存编辑的文章
		err := server.artInteractor.SaveDraft(article)
		if err != nil {
			return nil, err
		}
	}

	//3.保存草稿文章
	if in.Article.Id != 0 && domain.PublicStatus(in.Article.Status) == domain.Public {
		err := server.artInteractor.SaveDraft(article)
		if err != nil {
			return nil, err
		}
	}

	return &pb.SaveResponse{Id: in.Article.Id}, nil
}

func (server *articleServer) ViewDraftedArticles(ctxt context.Context, req *pb.ViewDraftedArticlesRequest) (*pb.ViewDraftedArticlesResponse, error) {
	arts, err := server.artInteractor.GetAuthorDrafts(domain.NUUID(req.UserId))
	if err != nil {
		return nil, err
	}

	return &pb.ViewDraftedArticlesResponse{Articles: ConvertArticles(arts)}, err
}

func (server *articleServer) View(ctxt context.Context, in *pb.ViewRequest) (*pb.ViewResponse, error) {
	art, err := server.artInteractor.GetArticle(domain.NUUID(in.Id))
	if err != nil {
		return nil, err
	}
	return &pb.ViewResponse{Article: &pb.Article{Id: int64(art.ID), Title: art.Title, Content: art.Content, Status: int32(art.Status), AuthorId: int64(art.AuthorID)}}, err

}

func (server *articleServer) Publish(ctxt context.Context, req *pb.PublishRequest) (*pb.PublishResponse, error) {
	err := server.artInteractor.Publish(domain.NUUID(req.Id), domain.NUUID(req.UserId))
	return nil, err
}

func (server *articleServer) ViewOwnPublishedArticles(ctx context.Context, req *pb.ViewOwnPublishedArticlesRequest) (*pb.ViewOwnPublishedArticlesResponse, error) {
	arts, err := server.artInteractor.GetAuthorPublicArticles(domain.NUUID(req.UserId))

	if err != nil {
		return nil, err
	}

	return &pb.ViewOwnPublishedArticlesResponse{Articles: ConvertArticles(arts)}, err

}

func (server *articleServer) Draft(ctx context.Context, req *pb.DraftRequest) (*pb.DraftResponse, error) {
	art, err := server.artInteractor.GetPublicArticleDraft(domain.NUUID(req.Id))
	if err != nil {
		return nil, err
	}
	return &pb.DraftResponse{Article: &pb.Article{Id: int64(art.ID), Title: art.Title, Content: art.Content, Status: int32(art.Status), AuthorId: int64(art.AuthorID)}}, err

}

func (server articleServer) ViewAllArticles(ctx context.Context, req *pb.ViewAllArticlesRequest) (*pb.ViewAllArticlesResponse, error) {
	arts, err := server.artInteractor.GetAllPublicArticles()
	if err != nil {
		return nil, err
	}
	return &pb.ViewAllArticlesResponse{Articles: ConvertArticles(arts)}, nil

}

//ConvertArticles .....
func ConvertArticles(arts []domain.Article) []*pb.Article {

	articles := make([]*pb.Article, len(arts))
	for i, art := range arts {
		articles[i] = &pb.Article{Id: int64(art.ID), Title: art.Title, Content: art.Content, Status: int32(art.Status), AuthorId: int64(art.AuthorID)}
	}
	return articles
}

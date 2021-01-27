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

	pb.RegisterArticleServer(server, &articleServer{artInteractor: artItor})
	server.Serve(conn)
}

func (server *articleServer) Save(ctxt context.Context, in *pb.SaveRequest) (*pb.SaveResponse, error) {
	if art.Id == 0 {
		ID, err := server.artInteractor.CreateDraft(usecases.GenerateUUID, domain.Article{ID: domain.NUUID(art.Id), Title: art.Title, Content: art.Content, Status: domain.Draft, AuthorID: domain.NUUID(art.AuthorID)}, domain.NUUID(art.AuthorID))
		if err != nil {
			return nil, err
		}

		return &pb.SaveArticleRep{Id: int64(ID)}, err
	}

	err := server.artInteractor.SaveDraft(domain.Article{ID: domain.NUUID(art.Id), Title: art.Title, Content: art.Content, Status: domain.Draft, AuthorID: domain.NUUID(art.AuthorID)}, domain.NUUID(art.AuthorID))
	if err != nil {
		return nil, err
	}

	return &pb.SaveArticleRep{Id: art.Id}, err
}

func (server *articleServer) ViewDraftedArticlesOfAuthor(ctxt context.Context, req *pb.ViewDraftedArticlesOfAuthorReq) (*pb.ViewDraftedArticlesOfAuthorRep, error) {
	arts, err := server.artInteractor.GetAuthorDrafts(domain.NUUID(req.UserID))
	if err != nil {
		return nil, err
	}

	return &pb.ViewDraftedArticlesOfAuthorRep{Articles: ConvertArticles(arts)}, err
}

func (server *articleServer) ViewArticle(ctxt context.Context, req *pb.ViewArticleReq) (*pb.Article, error) {
	art, err := server.artInteractor.GetArticle(domain.NUUID(req.Id))
	if err != nil {
		return nil, err
	}
	return &pb.Article{Id: int64(art.ID), Title: art.Title, Content: art.Content, Status: int32(art.Status), AuthorID: int64(art.AuthorID)}, err

}

func (server *articleServer) PublishArticle(ctxt context.Context, req *pb.PublishArticleReq) (*pb.PublishArticleRep, error) {
	err := server.artInteractor.Publish(domain.NUUID(req.Id), domain.NUUID(req.UserID))
	return nil, err
}

func (server *articleServer) ViewPublishedArticlesOfAuthor(ctx context.Context, req *pb.ViewPublishedArticlesOfAuthorReq) (*pb.ViewPublishedArticlesOfAuthorRep, error) {
	arts, err := server.artInteractor.GetAuthorPublicArticles(domain.NUUID(req.UserID))

	if err != nil {
		return nil, err
	}

	return &pb.ViewPublishedArticlesOfAuthorRep{Articles: ConvertArticles(arts)}, err

}

func (server *articleServer) ViewDraft(ctx context.Context, req *pb.ViewDraftReq) (*pb.Article, error) {
	art, err := server.artInteractor.GetPublicArticleDraft(domain.NUUID(req.Id))
	if err != nil {
		return nil, err
	}
	return &pb.Article{Id: int64(art.ID), Title: art.Title, Content: art.Content, Status: int32(art.Status), AuthorID: int64(art.AuthorID)}, err

}

func (server *articleServer) SaveDraft(ctx context.Context, art *pb.Article) (*pb.SaveDraftRep, error) {
	err := server.artInteractor.SaveDraft(domain.Article{ID: domain.NUUID(art.Id), Title: art.Title, Content: art.Content, Status: domain.Draft, AuthorID: domain.NUUID(art.AuthorID)}, domain.NUUID(art.AuthorID))
	return nil, err
}

func (server articleServer) PublishDraft(ctx context.Context, art *pb.Article) (*pb.PublishDraftRep, error) {
	err := server.artInteractor.PublishPublicArticleDraft(domain.Article{ID: domain.NUUID(art.Id), Title: art.Title, Content: art.Content, Status: domain.Draft, AuthorID: domain.NUUID(art.AuthorID)}, domain.NUUID(art.AuthorID))
	return nil, err
}

func (server articleServer) ViewAllArticles(ctx context.Context, req *pb.ViewAllArticlesReq) (*pb.ViewAllArticlesRep, error) {
	arts, err := server.artInteractor.GetAllPublicArticles()
	if err != nil {
		return nil, err
	}
	return &pb.ViewAllArticlesRep{Articles: ConvertArticles(arts)}, err

}

//ConvertArticles .....
func ConvertArticles(arts []domain.Article) []*pb.Article {

	articles := make([]*pb.Article, len(arts))
	for i, art := range arts {
		articles[i] = &pb.Article{Id: int64(art.ID), Title: art.Title, Content: art.Content, Status: int32(art.Status), AuthorID: int64(art.AuthorID)}
	}
	return articles
}

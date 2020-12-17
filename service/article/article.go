// Package article provides ...
package article

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

type articleServer struct {
	pb.UnimplementedArticleServer
	artInteractor usecases.ArticleInteractor
}

//Start ....
func Start(port int) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("listen port:%d err:%v\n", port, err)
	}
	log.Println("start server on port:", port)
	s := grpc.NewServer()

	handler, err := mysql.NewMysqlHandler("root@/real_world_medium?charset=utf8")
	if err != nil {
		log.Fatalln("connect db err:", err)
	}
	artRepo := interfaces.NewArticleRepo(handler)
	artItor := usecases.ArticleInteractor{ArticleRepo: artRepo}

	pb.RegisterArticleServer(s, &articleServer{artInteractor: artItor})
	s.Serve(lis)
}

func (server *articleServer) SaveArticle(ctxt context.Context, art *pb.Article) (*pb.SaveArticleRep, error) {
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

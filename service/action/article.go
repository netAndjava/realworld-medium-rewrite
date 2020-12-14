// Package action provides ...
package action

import (
	"context"
	"flag"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "iohttps.com/live/realworld-medium-go/service/api"
	"iohttps.com/live/realworld-medium-rewrite/domain"
	"iohttps.com/live/realworld-medium-rewrite/usecases"
)

type articleServer struct {
	pb.UnimplementedArticleServer
	artInteractor usecases.ArticleInteractor
}

func Start() {
	port := flag.String("port", ":8080", "please input port")
	flag.Parse()
	lis, err := net.Listen("tcp", *port)
	if err != nil {
		log.Fatalf("listen port:%d err:%v\n", port, err)
	}
	s := grpc.NewServer()

	// TODO:  <14-12-20, nqq> //
	var artItor usecases.ArticleInteractor

	pb.RegisterArticleServer(s, &articleServer{ArtInteractor: artItor})
}

func (server *articleServer) SaveArticle(ctxt context.Context, art *pb.Article) (*pb.SaveArticle, error) {
	if art.Id == 0 {
		ID, err := server.artInteractor.CreateDraft(usecases.GenerateUUID, domain.Article{ID: art.Id, Title: art.Title, Content: art.Content, Status: domain.Draft, AuthorID: art.AuthorID}, art.AuthorID)
		if err != nil {
			return nil, err
		}

		return &pb.SaveArticle{Id: ID}, err
	}

	err := server.artInteractor.SaveDraft(domain.Article{ID: art.Id, Title: art.Title, Content: art.Content, Status: domain.Draft, AuthorID: art.AuthorID}, art.AuthorID)
	if err != nil {
		return nil, err
	}

	return &pb.SaveArticle{Id: art.Id}, err
}

func (server *articleServer) ViewDraftedArticlesOfAuthor(ctxt context.Context, req *pb.ViewDraftedArticlesOfAuthorReq) (*pb.ViewDraftedArticlesOfAuthorRep, error) {
	arts, err := server.artInteractor.GetAuthorDrafts(req.UserID)
	if err != nil {
		return nil, err
	}

	return &pb.ViewDraftedArticlesOfAuthorRep{ConvertArticles(arts)}, err
}

func (server *articleServer) ViewArticle(ctxt context.Context, req *pb.ViewArticleReq) (*pb.Article, error) {
	art, err := server.artInteractor.GetArticle(req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Article{ID: art.ID, Title: art.Title, Content: art.Content, Status: art.Status, AuthorID: art.AuthorID}, err

}

func (server *articleServer) PublishArticle(ctxt context.Context, req *pb.PublishArticleReq) (*pb.PublishDraftRep, error) {
	err := server.artInteractor.Publish(req.Id, req.UserID)
	return nil, err
}

func (server *articleServer) ViewPublishedArticlesOfAuthor(ctx context.Context, req *pb.ViewPublishedArticlesOfAuthorReq) (*pb.ViewDraftedArticlesOfAuthorRep, error) {
	arts, err := server.artInteractor.GetAuthorPublicArticles(req.UserID)

	if err != nil {
		return nil, err
	}

	return &pb.ViewDraftedArticlesOfAuthorRep{Articles: ConvertArticles(arts)}, err

}

func (server *articleServer) ViewDraft(ctx context.Context, req *pb.ViewDraftReq) (*pb.Article, error) {
	art, err := server.artInteractor.GetPublicArticleDraft(req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Article{ID: art.ID, Title: art.Title, Content: art.Content, Status: art.Status, AuthorID: art.AuthorID}, err

}

func (server *articleServer) SaveDraft(ctx context.Context, art *pb.Article) (*pb.PublishDraftRep, error) {
	err := server.artInteractor.SaveDraft(domain.Article{ID: art.Id, Title: art.Title, Content: art.Content, Status: domain.Draft, AuthorID: art.AuthorID}, art.UserID)
	return nil, err
}

func (server articleServer) PublishDraft(ctx context.Context, art *pb.Article) (*pb.PublishDraftRep, error) {
	err := server.artInteractor.PublishPublicArticleDraft(domain.Article{ID: art.Id, Title: art.Title, Content: art.Content, Status: domain.Draft, AuthorID: art.AuthorID}, art.UserID)
	return nil, err
}

func (server articleServer) ViewAllArticles(ctx context.Context, req *pb.ViewAllArticlesReq) (*pb.ViewAllArticlesRep, error) {
	arts, err := server.artInteractor.GetAllPublicArticles()
	if err != nil {
		return nil, err
	}
	return &pb.ViewAllArticlesRep{ConvertArticles(arts)}, err

}

func ConvertArticles(arts []domain.Article) []*pb.Article {

	articles := make([]pb.Article, len(arts))
	for i, art := range arts {
		articles[i] = &pb.Article{ID: art.ID, Title: art.Title, Content: art.Content, Status: art.Status, AuthorID: art.AuthorID}
	}
	return articles
}

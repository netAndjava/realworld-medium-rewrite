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
		return pb.SaveArticle{Id: ID}, err
	}

	err := server.artInteractor.SaveDraft(domain.Article{ID: art.Id, Title: art.Title, Content: art.Content, Status: domain.Draft, AuthorID: art.AuthorID}, art.AuthorID)

	return &pb.SaveArticle{Id: art.Id}, err
}

func (server *articleServer) ViewDraftedArticlesOfAuthor(ctxt context.Context, req *pb.ViewDraftedArticlesOfAuthorReq) (*pb.ViewDraftedArticlesOfAuthorRep, error) {
	arts, err := server.artInteractor.GetAuthorDrafts(req.UserID)
	articles := make([]pb.Article, len(arts))
	for i, art := range arts {
		articles[i] = &pb.Article{ID: art.ID, Title: art.Title, Content: art.Content, Status: art.Status, AuthorID: art.AuthorID}
	}
	return &pb.ViewDraftedArticlesOfAuthor{Articles: articles}, err
}

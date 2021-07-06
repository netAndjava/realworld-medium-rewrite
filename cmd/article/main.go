package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"iohttps.com/live/realworld-medium-rewrite/cmd/config"
	"iohttps.com/live/realworld-medium-rewrite/endpoints"
	"iohttps.com/live/realworld-medium-rewrite/infrastructure/database"
	"iohttps.com/live/realworld-medium-rewrite/infrastructure/database/mysql"
	"iohttps.com/live/realworld-medium-rewrite/service/api"
	"iohttps.com/live/realworld-medium-rewrite/service/article"
	"iohttps.com/live/realworld-medium-rewrite/transport"
	"iohttps.com/live/realworld-medium-rewrite/usecases"
)

var logger log.Logger

func init() {
	logger = log.NewJSONLogger(os.Stdout)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)
}

func main() {

	//init db handler
	var dbConf mysql.Config
	dbConfig := flag.String("db", "./configs/mysql.toml", "please input config file of db")
	flag.Parse()
	_, err := config.Decode(*dbConfig, &dbConf)
	if err != nil {
		level.Error(logger).Log("decode config file:%s of db err:%v", *dbConfig, err)
		os.Exit(1)
	}

	handler, err := mysql.NewMysql(dbConf)
	if err != nil {
		level.Error(logger).Log(err)
		os.Exit(1)
	}

	f := flag.String("config", "./dev.toml", "please input config file")
	flag.Parse()
	var server config.Server
	_, err = config.Decode(*f, &server)
	if err != nil {
		level.Error(logger).Log("config file:%s err:%v\n", *f, err)
		os.Exit(1)
	}
	Start(server.IP, server.Port, handler)
}

func Start(IP, port string, handler database.DbHandler) {
	//1. New Article Interactor
	articleItor := usecases.NewArticleInteractor(handler)
	//2. New Article Service
	articleService := article.NewArticleService(logger, articleItor)
	//3. Make Endpoints
	edps := endpoints.MakeEndpoints(articleService)
	//4. New Article GRPC Server
	grpcServer := transport.NewArticleGRPCServer(edps)

	errs := make(chan error)
	//5. 监控退出信号
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGABRT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", IP, port))
	if err != nil {
		logger.Log("Listen err:", err)
		os.Exit(1)
	}
	//6. Register GRPC Service Server
	go func() {
		baseServer := grpc.NewServer()
		api.RegisterArticleServiceServer(baseServer, grpcServer)
		level.Info(logger).Log("msg", fmt.Sprintf("Server start on adderss:%s%s success!", IP, port))
		baseServer.Serve(listener)
	}()

	//7. 优雅的退出
	err = <-errs

	shutDown()
	level.Error(logger).Log("exit", err)
}

func shutDown() {
	level.Info(logger).Log("server shutdown")
}

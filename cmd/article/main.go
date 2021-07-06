package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"google.golang.org/grpc"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	consulapi "github.com/hashicorp/consul/api"
	"iohttps.com/live/realworld-medium-rewrite/cmd/config"
	"iohttps.com/live/realworld-medium-rewrite/endpoints"
	"iohttps.com/live/realworld-medium-rewrite/infrastructure/database"
	"iohttps.com/live/realworld-medium-rewrite/infrastructure/database/mysql"
	"iohttps.com/live/realworld-medium-rewrite/infrastructure/register"
	"iohttps.com/live/realworld-medium-rewrite/infrastructure/register/consul"
	"iohttps.com/live/realworld-medium-rewrite/interfaces"
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

	var consulConf consul.Config
	consulCfgPath := flag.String("consul", "../configs/consul.toml", "please input config file path for consul")
	config.Decode(*consulCfgPath, &consulConf)
	if err != nil {
		level.Error(logger).Log("decode config file:%s of consul err:%v", *dbConfig, err)
		os.Exit(1)
	}

	// TODO:1.指定grpc  2.获取实际物理IP<06-07-21, bantana> //
	registrar := consul.NewConsulRegister(consulapi.Config{Address: consulConf.Address}, consulapi.AgentServiceCheck{GRPC: "", Interval: consulConf.Check.Interval, Timeout: consulConf.Check.Timeout, Notes: consulConf.Check.Notes})

	f := flag.String("config", "./dev.toml", "please input config file")
	flag.Parse()
	var server config.Server
	_, err = config.Decode(*f, &server)
	if err != nil {
		level.Error(logger).Log("config file:%s err:%v\n", *f, err)
		os.Exit(1)
	}
	Start(server, handler, registrar)
}

func Start(server config.Server, handler database.DbHandler, registrar register.Registrar) {
	//1. New Article Interactor
	articleItor := usecases.ArticleInteractor{ArticleRepo: interfaces.NewArticleRepo(handler)}
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

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", server.IP, server.Port))
	if err != nil {
		logger.Log("Listen err:", err)
		os.Exit(1)
	}

	port, _ := strconv.Atoi(server.Port)
	//regsiter self
	rgtrar, err := registrar.Register(server.IP, port, server.Name, logger)
	//6. Register GRPC Service Server
	if err != nil {
		logger.Log("Register err:", err)
		os.Exit(1)
	}
	rgtrar.Register()
	go func() {
		baseServer := grpc.NewServer()
		api.RegisterArticleServiceServer(baseServer, grpcServer)
		level.Info(logger).Log("msg", fmt.Sprintf("Server start on adderss:%s%s success!", server.IP, server.Port))
		baseServer.Serve(listener)
	}()

	//7. 优雅的退出
	err = <-errs

	shutDown()
	level.Error(logger).Log("exit", err)
}

func shutDown() {
	// TODO:在服务退出时执行程序需要退出的  <06-07-21, bantana> //
	level.Info(logger).Log("server shutdown")
}

package main

import (
	"flag"
	"fmt"
	"log"

	"iohttps.com/live/realworld-medium-rewrite/cmd/config"
	"iohttps.com/live/realworld-medium-rewrite/infrastructure/database/mysql"
	"iohttps.com/live/realworld-medium-rewrite/service/article"
)

func main() {

	//init db handler
	var dbConf mysql.Config
	dbConfig := flag.String("db", "./configs/mysql.toml", "please input config file of db")
	flag.Parse()
	_, err := config.Decode(*dbConfig, &dbConf)
	if err != nil {
		log.Fatalf("decode config file:%s of db err:%v", *dbConfig, err)
	}

	handler, err := mysql.NewMysql(dbConf)
	if err != nil {
		log.Fatalln("connect db err:", err)
	}

	f := flag.String("config", "./dev.toml", "please input config file")
	flag.Parse()
	var server config.Server
	_, err = config.Decode(*f, &server)
	if err != nil {
		log.Fatalf("config file:%s err:%v\n", *f, err)
	}
	article.Start(fmt.Sprintf("%s:%s", server.IP, server.Port), handler)
}

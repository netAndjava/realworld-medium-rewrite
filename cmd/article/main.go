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
	f := flag.String("config", "./dev.toml", "please input config file")
	flag.Parse()
	conf, err := config.Decode(*f)
	if err != nil {
		log.Fatalf("config file:%s err:%v\n", *f, err)
	}

	//init db handler
	handelr, err := mysql.NewMysqlHandler(fmt.Sprintf("%s:%s@%s(%s:%s/%s)?charset=utf-8", conf.DB.User, conf.DB.Password, conf.DB.Network, conf.DB.Host, conf.DB.Port, conf.DB.Name))
	if err != nil {
		log.Fatalln("init db err:", err)
	}
	article.Start(fmt.Sprintf("%s:%s", conf.Server.IP, conf.Server.Port), handelr)
}

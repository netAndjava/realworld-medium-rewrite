// Package main provides ...
package main

import (
	"flag"
	"fmt"
	"log"

	"iohttps.com/live/realworld-medium-rewrite/cmd/config"
	"iohttps.com/live/realworld-medium-rewrite/infrastructure/database/mysql"
	"iohttps.com/live/realworld-medium-rewrite/service/comment"
)

func main() {
	var dbConf mysql.Config
	dbConfig := flag.String("db", "./configs/mysql.toml", "please input config file of db")
	flag.Parse()
	_, err := config.Decode(*dbConfig, &dbConf)
	if err != nil {
		log.Fatalf("decode config file:%s of db err:%v", *dbConfig, err)
	}
	hander, err := mysql.NewMysql(dbConf)
	if err != nil {
		log.Fatalln("init db err:", err)
	}

	c := flag.String("config", "./dev.toml", "please input config file")
	flag.Parse()
	var serConf config.Server
	_, err = config.Decode(*c, &serConf)
	if err != nil {
		log.Fatalf("config file:%s err:%v\n", *c, err)
	}
	comment.Start(fmt.Sprintf("%s:%s", serConf.IP, serConf.Port), hander)
}

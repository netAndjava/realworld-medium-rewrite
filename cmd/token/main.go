// Package main provides ...
package main

import (
	"flag"
	"fmt"
	"log"

	"iohttps.com/live/realworld-medium-rewrite/cmd/config"
	"iohttps.com/live/realworld-medium-rewrite/infrastructure/database/mysql"
	"iohttps.com/live/realworld-medium-rewrite/service/token"
)

func main() {
	dbConfig := flag.String("db", "./configs/mysql.toml", "please input config file of db")
	flag.Parse()
	var dbConf mysql.Config
	_, err := config.Decode(*dbConfig, &dbConf)
	if err != nil {
		log.Fatalf("decode config file:%s of db err:%v\n", *dbConfig, err)
	}
	handler, err := mysql.NewMysql(dbConf)
	if err != nil {
		log.Fatalln("init db err:", err)
	}

	c := flag.String("config", "./dev.toml", "please input config file")
	flag.Parse()
	var serverConf config.Server
	_, err = config.Decode(*c, &serverConf)
	if err != nil {
		log.Fatalf("decode config file:%s err:%v\n", *c, err)
	}

	token.Start(fmt.Sprintf("%s:%s", serverConf.IP, serverConf.Port), handler)
}

// Package user provides ...
package main

import (
	"flag"
	"fmt"
	"log"

	"iohttps.com/live/realworld-medium-rewrite/cmd/config"
	"iohttps.com/live/realworld-medium-rewrite/infrastructure/database/mysql"
	"iohttps.com/live/realworld-medium-rewrite/service/user"
)

func main() {
	dbConfig := flag.String("db", "./configs/mysql.toml", "please input config of db")
	c := flag.String("config", "./dev.toml", "please input config file")

	flag.Parse()
	var dbConf mysql.Config
	_, err := config.Decode(*dbConfig, &dbConf)
	if err != nil {
		log.Fatalf("decode file:%s of db err:%v", *dbConfig, err)
	}
	handler, err := mysql.NewMysql(dbConf)
	if err != nil {
		log.Fatal("connect db err:", err)
	}

	var serverConf config.Server
	_, err = config.Decode(*c, &serverConf)
	if err != nil {
		log.Fatalf("config file:%s err:%v", *c, err)
	}
	user.Start(fmt.Sprintf("%s:%s", serverConf.IP, serverConf.Port), handler)
}

// func Start(address string, handler database.DbHandler) {

//1. address { ip : port }
//2. dbusername , dbpassword, dbname, charset

// config.address
// config.dbusername
// config.dbpassword
// ...

// toml

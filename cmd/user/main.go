// Package user provides ...
package main

import "iohttps.com/live/realworld-medium-rewrite/service/user"

var configUserListen = 3001

func main() {
	user.Start(configUserListen)
}

// func Start(address string, handler database.DbHandler) {

//1. address { ip : port }
//2. dbusername , dbpassword, dbname, charset

// config.address
// config.dbusername
// config.dbpassword
// ...

// toml

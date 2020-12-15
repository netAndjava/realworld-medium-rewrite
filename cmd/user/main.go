// Package user provides ...
package main

import "iohttps.com/live/realworld-medium-rewrite/service/user"

var ConfigUserListen = 3001

func main() {
	user.Start(ConfigUserListen)
}

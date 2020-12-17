// Package main provides ...
package main

import "iohttps.com/live/realworld-medium-rewrite/service/token"

var configTokenListen = 30002

func main() {
	token.Start(configTokenListen)
}

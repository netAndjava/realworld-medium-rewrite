// Package main provides ...
package main

import "iohttps.com/live/realworld-medium-rewrite/service/token"

var ConfigTokenListen = 30001

func main() {
	token.Start(ConfigTokenListen)
}

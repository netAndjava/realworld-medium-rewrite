// Package main provides ...
package main

import "iohttps.com/live/realworld-medium-rewrite/service/comment"

var configCommentPort = 30003

func main() {
	comment.Start(configCommentPort)
}

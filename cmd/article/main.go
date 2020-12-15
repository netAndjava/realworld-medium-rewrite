package main

import "iohttps.com/live/realworld-medium-rewrite/service/action"

var ConfigArticlelisten = 3000
var ConfigCommentlisten = 3000

func main() {
	action.Start()
	// Article.Start(ConfigArticlelisten) // 3000
	// Comment.Start(ConfigCommentlisten) // 3001
}

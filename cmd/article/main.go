package main

import "iohttps.com/live/realworld-medium-rewrite/service/article"

var ConfigArticlelisten = 3000

// var ConfigCommentlisten = 3000

func main() {
	article.Start(ConfigArticlelisten) // 3000
	// Comment.Start(ConfigCommentlisten) // 3001
}

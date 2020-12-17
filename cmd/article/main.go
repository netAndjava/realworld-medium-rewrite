package main

import "iohttps.com/live/realworld-medium-rewrite/service/article"

var configArticlelisten = 3000

// var ConfigCommentlisten = 3000

func main() {
	article.Start(configArticlelisten) // 3000
	// Comment.Start(ConfigCommentlisten) // 3001
}

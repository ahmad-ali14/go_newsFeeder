package main

import (
	"newsFeeder/httpd/handler"
	"newsFeeder/platform/newsFeed"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	feed := newsFeed.New()

	/*
	* we can't pass the function by reference here, we need to execute it, so we have the ability to pass vars.
	* we've done this bvecuase we are returning annonymos func at handler.PingGet
	 */
	r.GET("/ping", handler.PingGet())
	r.GET("/newsfeed", handler.NewsFeedGet(feed))
	r.POST("/newsfeed", handler.NewsFeedPost(feed))

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}

package handler

import (
	"net/http"
	"newsFeeder/platform/newsFeed"

	"github.com/gin-gonic/gin"
)

func NewsFeedGet(feed *newsFeed.Repo) gin.HandlerFunc {
	/*
	* we do this so we can pass dependecies to this function from main.go
	 */
	return func(c *gin.Context) {
		result := feed.GetAll()
		c.JSON(http.StatusOK, result)
	}
}

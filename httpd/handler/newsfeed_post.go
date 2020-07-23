package handler

import (
	"net/http"
	"newsFeeder/platform/newsFeed"

	"github.com/gin-gonic/gin"
)

type newsFeedPostRequest struct {
	Title string `json:"title"`
	Post  string `json:"post"`
}

func NewsFeedPost(feed *newsFeed.Repo) gin.HandlerFunc {
	/*
	* we do this so we can pass dependecies to this function from main.go
	 */
	return func(c *gin.Context) {

		requestBody := newsFeedPostRequest{}
		c.Bind(&requestBody)

		item := newsFeed.Item{
			Title: requestBody.Title,
			Post:  requestBody.Post,
		}

		feed.Add(item)

		c.Status(http.StatusNoContent)
	}
}

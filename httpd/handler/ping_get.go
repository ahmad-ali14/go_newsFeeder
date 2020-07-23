package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PingGet() gin.HandlerFunc {
	/*
	* we do this so we can pass dependecies to this function from main.go
	 */
	return func(c *gin.Context) {

		c.JSON(http.StatusOK, map[string]string{
			"hello": "Found me",
		})
	}
}

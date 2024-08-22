package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PageHandler(c *gin.Context) {

	data := gin.H{
		"Title": "Weather App",
	}
	c.HTML(http.StatusOK, "index.html", data)
}

package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func PageHandler(c *gin.Context) {
	data := gin.H{
		"Title": "Weather App",
	}
	c.HTML(http.StatusOK, "index.html", data)
}

func WeatherHandler(c *gin.Context) {
	apiKey := os.Getenv("APIKEY")

	// Create a default HTTP client
	client := &http.Client{}
	zip := c.Query("zip")
	if zip == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Zip code is required"})
		return
	}

	weather, err := getWeather(apiKey, zip, client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch weather data"})
		return
	}

	c.JSON(http.StatusOK, weather)
}

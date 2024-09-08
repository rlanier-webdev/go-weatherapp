// handlers.go
package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

const defaultLocation = "10001" // Default ZIP code

func PageHandler(c *gin.Context) {
	apiKey := os.Getenv("APIKEY")
	client := &http.Client{}

	// Fetch weather data for the default location
	weather, err := getWeather(apiKey, defaultLocation, client)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "index.html", gin.H{
			"Title": "Weather App",
			"Error": "Failed to fetch default weather data",
		})
		return
	}

	data := gin.H{
		"Title":   "Weather App",
		"Weather": weather,
	}
	c.HTML(http.StatusOK, "index.html", data)
}

func WeatherHandler(c *gin.Context) {
	apiKey := os.Getenv("APIKEY")
	client := &http.Client{}

	zip := c.Query("zip")
	if zip == "" {
		c.Redirect(http.StatusFound, "/") // Redirect to homepage if no ZIP code is provided
		return
	}

	weather, err := getWeather(apiKey, zip, client)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "index.html", gin.H{
			"Title": "Weather Dashboard",
			"Error": "Failed to fetch weather data",
		})
		return
	}

	data := gin.H{
		"Title":   "Weather App",
		"Weather": weather,
	}
	c.HTML(http.StatusOK, "index.html", data)
}

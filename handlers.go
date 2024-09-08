// handlers.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

const defaultLocation = "10001" // Default ZIP code
const title = "Weather Dashboard"

// PageHandler handles the homepage route
func PageHandler(c *gin.Context) {
	weather, err := fetchWeather(defaultLocation)
	if err != nil {
		log.Printf("Error fetching default weather data: %v", err)
		c.HTML(http.StatusInternalServerError, "index.html", gin.H{
			"Title": title,
			"Error": "Failed to fetch default weather data",
		})
		return
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"Title":   title,
		"Weather": weather,
	})
}

// WeatherHandler handles the weather request based on ZIP code
func WeatherHandler(c *gin.Context) {
	zip := c.Query("zip")
	if zip == "" {
		c.Redirect(http.StatusFound, "/")
		return
	}

	weather, err := fetchWeather(zip)
	if err != nil {
		log.Printf("Error fetching weather data for zip %s: %v", zip, err)
		c.HTML(http.StatusInternalServerError, "index.html", gin.H{
			"Title": title,
			"Error": "Failed to fetch weather data",
		})
		return
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"Title":   title,
		"Weather": weather,
	})
}

// fetchWeather abstracts the weather fetching logic
func fetchWeather(location string) (*WeatherResponse, error) {
	apiKey := os.Getenv("APIKEY")
	if apiKey == "" {
		return nil, fmt.Errorf("API key is missing")
	}

	client := &http.Client{} // Customize this client as needed
	return getWeather(apiKey, location, client)
}

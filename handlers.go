// handlers.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/gin-gonic/gin"
)

const defaultLocation = "10001" // Default ZIP code
const title = "Weather Dashboard"

// PageHandler handles the homepage route
func PageHandler(c *gin.Context) {
	weather, err := fetchWeather(defaultLocation)
	renderWeather(c, weather, err, defaultLocation)
}

// WeatherHandler handles the weather request based on ZIP code
func WeatherHandler(c *gin.Context) {
	zip := c.DefaultQuery("zip", defaultLocation)

	if !isValidZipCode(zip) {
		c.HTML(http.StatusBadRequest, "index.html", gin.H{
			"Title": title,
			"Error": "Invalid ZIP code format. Please enter a valid 5-digit ZIP code.",
		})
		return
	}

	weather, err := fetchWeather(zip)
	renderWeather(c, weather, err, zip)
}

var httpClient = &http.Client{}

// fetchWeather abstracts the weather fetching logic
func fetchWeather(location string) (*WeatherResponse, error) {
	apiKey := os.Getenv("APIKEY")
	if apiKey == "" {
		log.Println("Error: API key is missing.")
		return nil, fmt.Errorf("API key is missing")
	}

	return getWeather(apiKey, location, httpClient) // Use the global client
}

// Helper function
func renderWeather(c *gin.Context, weather *WeatherResponse, err error, zip string) {
	if err != nil {
		log.Printf("Error fetching weather data for zip %s: %v", zip, err)
		if err.Error() == "API key is missing" {
			c.HTML(http.StatusInternalServerError, "index.html", gin.H{
				"Title": title,
				"Error": "API key is missing. Please check your configuration.",
			})
		} else if err.Error() == fmt.Sprintf("weather data not found for location: %s", zip) {
			c.HTML(http.StatusNotFound, "index.html", gin.H{
				"Title": title,
				"Error": fmt.Sprintf("Weather data not found for zip code: %s", zip),
			})
		} else if err.Error() == fmt.Sprintf("invalid ZIP code. No matching location found for: %s", zip) {
			c.HTML(http.StatusBadRequest, "index.html", gin.H{
				"Title": title,
				"Error": fmt.Sprintf("Invalid ZIP code: %s. No matching location found.", zip),
			})
		} else {
			c.HTML(http.StatusInternalServerError, "index.html", gin.H{
				"Title": title,
				"Error": "Failed to fetch weather data",
			})
		}
		return
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"Title":   title,
		"Weather": weather,
	})
}

// isValidZipCode checks if a ZIP code is valid
func isValidZipCode(zip string) bool {
	// Check if the ZIP code is 5 digits long and contains only digits
	match, _ := regexp.MatchString(`^\d{5}$`, zip)
	return match
}

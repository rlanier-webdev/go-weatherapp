package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
}

func getWeather(apiKey, location string, client *http.Client) (*WeatherResponse, error) {
	baseURL := "http://api.weatherapi.com/v1/current.json"

	// Define query parameters
	params := url.Values{}
	params.Add("key", apiKey)
	params.Add("q", location)
	params.Add("aqi", "no")

	// Construct the final URL
	finalURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	// Make the GET request using the passed client
	resp, err := client.Get(finalURL)
	if err != nil {
		log.Println("Error:", err)
		return nil, err
	}
	defer resp.Body.Close()

	// Read and parse the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error:", err)
		return nil, err
	}

	// Check for HTTP errors
	if resp.StatusCode != http.StatusOK {
		// Try to parse the error response from the API
		var apiError WeatherError
		if err := json.Unmarshal(body, &apiError); err == nil {
			if apiError.Error.Code == 1006 {
				return nil, fmt.Errorf("invalid ZIP code. No matching location found for: %s", location)
			}
		}

		if resp.StatusCode == http.StatusNotFound {
			return nil, fmt.Errorf("weather data not found for location: %s", location)
		}
		return nil, fmt.Errorf("HTTP error: %s (status code: %d)", resp.Status, resp.StatusCode)
	}

	// Parse the JSON response
	var weather WeatherResponse
	err = json.Unmarshal(body, &weather)
	if err != nil {
		log.Println("Error:", err)
		return nil, err
	}
	return &weather, nil
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.Static("/static", "./static")

	// Route for the homepage
	r.GET("/", PageHandler)
	r.GET("/weather", WeatherHandler)

	port := "8080" // Default port
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	if err := r.Run(":" + port); err != nil {
		log.Println("Failed to run server: ", err)
	}
}

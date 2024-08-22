package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}

func getWeather(apiKey, location string, client *http.Client) (*WeatherResponse, error) {
	baseURL := "http://api.weatherapi.com/v1/current.json"

	// Define query parametrs
	params := url.Values{}
	params.Add("key", apiKey)
	params.Add("q", location)
	params.Add("aqi", "no")

	// Construct the final URL
	finalURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	// Make the GET request
	resp, err := http.Get(finalURL)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	defer resp.Body.Close()

	// Read and print the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	var weather WeatherResponse

	err = json.Unmarshal(body, &weather)
	if err != nil {
		fmt.Println("Error: ", err)
		return nil, err
	}
	return &weather, nil
}

func main() {
	//gin.SetMode(gin.ReleaseMode)
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
		log.Fatal("Failed to run server: ", err)
	}
}

func getUserInput() string {
	// Create a new buffered reader to read input from the standard input
	scanner := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the zip code of the location: ")

	zipCode, err := scanner.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return ""
	}
	zipCode = strings.TrimSpace(zipCode)
	return zipCode
}

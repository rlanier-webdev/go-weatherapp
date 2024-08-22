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
	apiKey := os.Getenv("APIKEY")
	location := getUserInput()

	// Create a default HTTP client
	client := &http.Client{}

	weather, err := getWeather(apiKey, location, client)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Extract and print the required weather information
	fmt.Printf("Location: %s, %s, %s\n", weather.Location.Name, weather.Location.Region, weather.Location.Country)
	fmt.Printf("Temperature: %.1f°C (%.1f°F)\n", weather.Current.TempC, weather.Current.TempF)
	fmt.Printf("Condition: %s\n", weather.Current.Condition.Text)
	fmt.Printf("Humidity: %d%%\n", weather.Current.Humidity)
	fmt.Printf("Cloud Cover: %d%%\n", weather.Current.Cloud)
	fmt.Printf("Feels Like: %.1f°C (%.1f°F)\n", weather.Current.FeelsLikeC, weather.Current.FeelsLikeF)
	fmt.Printf("Wind Speed: %.1f mph (%.1f kph)\n", weather.Current.WindMph, weather.Current.WindKph)
	fmt.Printf("Wind Chill: %.1f°C (%.1f°F)\n", weather.Current.WindchillC, weather.Current.WindchillF)
	fmt.Printf("Heat Index: %.1f°C (%.1f°F)\n", weather.Current.HeatindexC, weather.Current.HeatindexF)
	fmt.Printf("Dew Point: %.1f°C (%.1f°F)\n", weather.Current.DewpointC, weather.Current.DewpointF)
	fmt.Printf("Visibility: %.1f km (%.1f miles)\n", weather.Current.VisKm, weather.Current.VisMiles)
	fmt.Printf("UV Index: %.1f\n", weather.Current.Uv)
	fmt.Printf("Gust: %.1f mph (%.1f kph)\n", weather.Current.GustMph, weather.Current.GustKph)

	port := "8080" // Default port
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	//gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.Static("/static", "./static")
	r.GET("/", PageHandler)

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

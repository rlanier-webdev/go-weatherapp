Here’s an updated README that includes instructions for using an `.env` file to store your API key:

---

# Weather App

This Weather App is a command-line application written in Go that allows you to fetch real-time weather data for any location based on a zip code. The app leverages the [WeatherAPI.com](https://www.weatherapi.com/) to retrieve and display detailed weather information, including temperature, humidity, wind speed, and more.

## Features

- **Real-Time Weather Data**: Get up-to-date weather information for any location by entering its zip code.
- **Detailed Weather Metrics**: Displays temperature, weather condition, humidity, wind speed, cloud cover, visibility, and more.
- **User Input**: Dynamically input the zip code to retrieve specific weather data for any location.
- **Environment Variable for API Key**: Store your WeatherAPI.com API key securely in an `.env` file.

## Prerequisites

- **Go**: You must have [Go](https://golang.org/dl/) installed on your system.
- **WeatherAPI.com API Key**: You need to sign up for a free API key from [WeatherAPI.com](https://www.weatherapi.com/).

## Installation

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/rlanier-webdev/go-weatherapp.git
   cd weather-app
   ```

2. **Set Up Your API Key**:
   - Create a `.env` file in the root directory of the project.
   - Add your WeatherAPI.com API key to the `.env` file:
     ```bash
     APIKEY=your_api_key_here
     ```
   - Ensure your `.env` file is included in your `.gitignore` to keep it out of version control:
     ```bash
     echo ".env" >> .gitignore
     ```

3. **Install Go Modules**:
   - If you are using Go modules, ensure that your dependencies are up-to-date by running:
     ```bash
     go mod tidy
     ```

4. **Build the Application**:
   ```bash
   go build -o go-weatherapp
   ```

5. **Run the Application**:
   ```bash
   ./go-weatherapp
   ```

## Usage

1. When you run the application, it will prompt you to enter the zip code of the location for which you want to retrieve weather data:
   ```bash
   Enter the zip code of the location:
   ```

2. After entering the zip code, the application will fetch and display the weather information, including:
   - Location name
   - Temperature (in Celsius and Fahrenheit)
   - Weather condition (e.g., sunny, cloudy)
   - Humidity
   - Wind speed
   - Feels-like temperature
   - Visibility
   - UV Index
   - Wind gusts

Example output:
```bash
Location: Saint Louis, Missouri, USA
Temperature: 22.8°C (73.0°F)
Condition: Partly cloudy
Humidity: 46%
Cloud Cover: 75%
Feels Like: 24.2°C (75.5°F)
Wind Speed: 10.5 mph (16.9 kph)
Wind Chill: 22.4°C (72.3°F)
Heat Index: 23.9°C (75.0°F)
Dew Point: 7.3°C (45.2°F)
Visibility: 16.0 km (9.0 miles)
UV Index: 6.0
Gust: 15.0 mph (24.1 kph)
```

## Environment Variables

To securely manage your WeatherAPI.com API key:

1. **Create a `.env` File**:
   - Place your API key in a `.env` file at the root of the project:
     ```bash
     APIKEY=your_api_key_here
     ```

2. **Load the API Key in Your Code**:
   - Use a package like `godotenv` to load the environment variables in your Go application:
     ```go
     import (
         "os"
         "github.com/joho/godotenv"
     )

     func init() {
         err := godotenv.Load()
         if err != nil {
             fmt.Println("Error loading .env file")
         }
     }

     func main() {
         apiKey := os.Getenv("APIKEY")
         // Use apiKey for your API requests
     }
     ```
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Define a struct to hold the weather data
type WeatherData struct {
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
	} `json:"main"`
	Timezone int    `json:"timezone"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
}

// Define a constant for the API key
const apiKey = "1409e3baf1bb1b7f4fd6b2c0aad875cb"
const city = "6167865"

// Define a function to make a request to the OpenWeatherMap API
func getWeatherData(city string) (*WeatherData, error) {
	// Construct the API URL with the city and the API key
	apiURL := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, apiKey)

	// Make a GET request to the API URL
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body as bytes
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Parse the JSON data into a WeatherData struct
	var weatherData WeatherData
	err = json.Unmarshal(body, &weatherData)
	if err != nil {
		return nil, err
	}

	// Return the weather data
	return &weatherData, nil
}

// Define a function to format the weather data into JSON
func formatWeatherData(weatherData *WeatherData) (string, error) {
	// Marshal the weather data into JSON bytes
	jsonData, err := json.Marshal(weatherData)
	if err != nil {
		return "", err
	}

	// Convert the JSON bytes into a string
	jsonString := string(jsonData)

	// Return the JSON string
	return jsonString, nil
}

// Define a handler function for the /weather endpoint
func weatherHandler(w http.ResponseWriter, r *http.Request) {
	// Get the weather data for Toronto
	weatherData, err := getWeatherData("Toronto")
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Format the weather data into JSON
	jsonString, err := formatWeatherData(weatherData)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Set the content type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON string to the response
	w.Write([]byte(jsonString))
}

// Define the main function
func main() {
	// Create a new multiplexer
	mux := http.NewServeMux()

	// Register the /weather handler
	mux.HandleFunc("/weather", weatherHandler)

	// Start the server on port 8080
	log.Println("Starting server on port 8080")
	http.ListenAndServe(":8080", mux)
}

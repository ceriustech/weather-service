package server

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var homePage = `
<!DOCTYPE html>
<html>
	<head>
		<title>Weather Service</title>
	</head>
	<body>
    <h1>Weather Report</h1>
    <p><strong>Location:</strong> {{.Name}}</p>
    <p><strong>Latitude:</strong> {{.Coord.Lat}}</p>
    <p><strong>Longitude:</strong> {{.Coord.Lon}}</p>
		{{range .Weather}}
			<p><strong>Condition:</strong> {{.Main}} - {{.Description}}</p>
		{{end}}
	</body>
</html>
`

type WeatherData struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
			Main        string `json:"main"`
			Description string `json:"description"`
	} `json:"weather"`
	Name string `json:"name"`
}

// fetchWeatherData fetches the weather data from the OpenWeather API
func fetchWeatherData(lat, lon, apiKey string) (*WeatherData, error) {
	client := &http.Client{
			Timeout: 10 * time.Second,
	}
	requestURL := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%s&lon=%s&appid=%s", lat, lon, apiKey)

	resp, err := client.Get(requestURL)
	if err != nil {
			return nil, fmt.Errorf("error fetching weather data: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("incorrect status code: %d", resp.StatusCode)
	}

	var data WeatherData
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
			return nil, fmt.Errorf("error parsing weather data: %w", err)
	}

	return &data, nil
}

// HandleCurrentWeatherResponse processes requests for current weather data and either renders HTML or JSON data
func HandleCurrentWeatherResponse(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	lat, lon := r.URL.Query().Get("lat"), r.URL.Query().Get("lon")
	apiKey := os.Getenv("OPENWEATHER_API_KEY")
	if apiKey == "" {
		log.Fatal("API key for OpenWeather is not set")
	}

	data, err := fetchWeatherData(lat, lon, apiKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Check if the client accepts HTML
	if strings.Contains(r.Header.Get("Accept"), "text/html") {
		tmpl, err := template.New("weather").Parse(homePage)
		if err != nil {
			log.Printf("Error creating template: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		tmpl.Execute(w, data) // Render the template with the weather data
	} else {
		// Default to JSON response if not accepting HTML
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(data); err != nil {
			log.Printf("Error encoding JSON response: %v", err)
		}
	}
}




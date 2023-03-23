package weather

import (
	"log"
	"net/http"
	"strings"
	"testing"
)

func newWeatherProvider() WeatherProvider {
	return WeatherProvider{}
}

func TestWeatherAPIResponse(t *testing.T) {

	response, err := GetCurrentWeather("Kathmandu")
	if err != nil {
		t.Fatal("Weather API is down", err)
	}

	if response.StatusCode != http.StatusOK {
		t.Fatal("Weather API fetch operation failed", err)
	}
}

func TestGetWeatherCityByName(t *testing.T) {
	cityName := "kathmandu"
	weatherProvider := newWeatherProvider()
	response, err := weatherProvider.GetWeatherDataByCityName(cityName)
	if err != nil {
		t.Fatal("Get weather data unit failed", err)
	}

	if strings.ToLower(response.Name) != cityName {
		t.Fatalf("Got incorrect city details : want %v : got %v", cityName, response.Name)
	}

	log.Println(response)
}

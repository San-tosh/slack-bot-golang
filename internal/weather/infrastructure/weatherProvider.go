package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	domain "weather-bot/internal/weather/domain"
)

const (
	URL_OPEN_WEATHER_MAP_API = "http://api.openweathermap.org/data/2.5/weather"
	APP_ID                   = "b9a49f1ec0a625c5086e6b99862e5d5c"
)

type WeatherProvider struct{}

func (wp WeatherProvider) GetWeatherDataByCityName(cityName string) (*domain.Forecast, error) {
	resp, err := GetCurrentWeather(cityName)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		defer resp.Body.Close()
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var f domain.Forecast
	err = json.Unmarshal(body, &f)

	if err != nil {
		return nil, err
	}

	return &f, nil
}

func GetCurrentWeather(city string) (*http.Response, error) {

	url := fmt.Sprintf(URL_OPEN_WEATHER_MAP_API+"?q=%s&appid=%s", city, APP_ID)

	res, err := http.Get(url)

	if err != nil {
		return res, err
	}
	return res, nil
}

package weather

import (
	domain "weather-bot/internal/weather/domain"
)

type WeatherService struct {
	repo domain.IWeatherRepository
}

func NewWeatherService(r domain.IWeatherRepository) domain.IWeatherService {
	return WeatherService{
		repo: r,
	}
}

func (w WeatherService) GetWeatherDataByCityName(cityName string) (*domain.Forecast, error) {
	return w.repo.GetWeatherDataByCityName(cityName)
}

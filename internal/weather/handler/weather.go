package weather

import (
	domain "weather-bot/internal/weather/domain"
)

type Handler struct {
	Weather domain.IWeatherRepository
}

func NewHandler(weather domain.IWeatherRepository) *Handler {
	return &Handler{
		Weather: weather,
	}
}

func (h *Handler) GetWeatherDataByCityName(cityName string) (*domain.Forecast, error) {
	return h.Weather.GetWeatherDataByCityName(cityName)
}

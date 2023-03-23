package weather

type Forecast struct {
	Id              float64   `json:"id"`
	Name            string    `json:"name"`
	Info            MainBlock `json:"main"`
	Visibility      int       `json:"visibility"`
	WeatherForecast []Weather `json:"weather"`
}

type Weather struct {
	Main        string `json:"main"`
	Description string `json:"description"`
}

type MainBlock struct {
	Temp     float64 `json:"temp"`
	Pressure float64 `json:"pressure"`
	Humidity float64 `json:"humidity"`
	TempMin  float64 `json:"temp_min"`
	TempMax  float64 `json:"temp_max"`
}

type IWeatherService interface {
	GetWeatherDataByCityName(cityName string) (*Forecast, error)
}

type IWeatherRepository interface {
	GetWeatherDataByCityName(cityName string) (*Forecast, error)
}

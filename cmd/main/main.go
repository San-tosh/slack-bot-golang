package main

import (
	"context"
	"fmt"
	"log"
	app "weather-bot/internal/weather/application/services"
	domain "weather-bot/internal/weather/domain"
	handler "weather-bot/internal/weather/handler"
	infra "weather-bot/internal/weather/infrastructure"

	"github.com/shomali11/slacker"
)

type (
	Celsius float64
)

const (
	SLACK_APP_TOKEN = "xapp-1-A04UZP5S9UL-4985791710068-260c2c9cc811cd0e0f7e1c14a30bb017b8c1eca58baa53e8fd55ee919b59109a"
	SLACK_BOT_TOKEN = "xoxb-4995929414929-4976764183990-EoObXdf68o2Sxn2EaVFoznlV"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("------Command Event Details------")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println("---------------------------------")
	}
}

func main() {
	bot := slacker.NewClient(SLACK_BOT_TOKEN, SLACK_APP_TOKEN)

	go printCommandEvents(bot.CommandEvents())

	definition := &slacker.CommandDefinition{
		Description: "city kathmandu",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			word := request.Param("word")
			weatherResponse := getWeatherResponse(word)
			response.Reply(formatData(weatherResponse))
		},
	}

	bot.Command("city {word}", definition)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func getWeatherResponse(cityName string) *domain.Forecast {
	weatherService := app.NewWeatherService(infra.WeatherProvider{})
	services := handler.NewHandler(weatherService)
	response, error := services.GetWeatherDataByCityName(cityName)
	if error != nil {
		fmt.Println("Error while fetching weather data", error)
	}
	return response
}

func formatData(data *domain.Forecast) string {
	if data.Name == "" {
		return "No such city found!!"
	}
	format := fmt.Sprintf(
		"CityName : %v \nVisibility : %v km\nTemperature : %.2f °C\nPressure: %v hpa\nHumidity: %.0f%% \n"+
			"Minimum Temperature: %.2f °C \nMaximum Temperature: %.2f °C",
		data.Name, data.Visibility, kelvin2Celsius(data.Info.Temp), data.Info.Pressure, data.Info.Humidity,
		kelvin2Celsius(data.Info.TempMin), kelvin2Celsius(data.Info.TempMax))
	for _, value := range data.WeatherForecast {
		format += fmt.Sprintf("\n%v: %v", value.Main, value.Description)
	}
	return format
}

func kelvin2Celsius(k float64) Celsius {
	return Celsius(k - 273.15)
}

package application

import (
	"go-agent/internal/domain"
)

func InitWeather(weatherServicePort WeatherServicePort) WeatherPort {
	return &Weather{
		weatherServicePort: weatherServicePort,
	}
}

type Weather struct {
	weatherServicePort WeatherServicePort
}

func (c Weather) GetWeather(opts domain.CheckWeatherOptions) (string, error) {
	weather, err := c.weatherServicePort.GetWeather(opts)
	if err != nil {
		return "", err
	}
	return weather, nil
}

package weather

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"go-agent/internal/application"
	"go-agent/internal/domain"
)

func NewWeatherClient() application.WeatherServicePort {
	return &adapter{
		client: resty.New().SetBaseURL("https://wttr.in"),
	}
}

type adapter struct {
	client *resty.Client
}

func (a adapter) GetWeather(opts domain.CheckWeatherOptions) (string, error) {
	response, err := a.client.R().
		SetPathParam("city", opts.City).
		SetQueryParam("format", "3").
		Get("/{city}")

	if err != nil {
		return "", fmt.Errorf("failed to execute weather request: %w", err)
	}
	if response.IsError() {
		return "", fmt.Errorf("weather request failed with status: %s", response.Status())
	}

	return response.String(), nil
}

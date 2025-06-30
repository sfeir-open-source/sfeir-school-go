package tools

import anthropiclocal "Go/internal/infrastructure/anthropic"

type GetWeatherInput struct {
	City string `json:"city" jsonschema_description:"The name of the city to get the weather for."`
}

var GetWeatherDefinition = anthropiclocal.ToolDefinition{
	Name:        "get_weather",
	Description: "Get the current weather for a given city.",
	InputSchema: anthropiclocal.GenerateSchema[GetWeatherInput](),
	Function:    GetWeather,
}

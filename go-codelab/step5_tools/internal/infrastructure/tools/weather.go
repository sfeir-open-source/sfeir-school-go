package tools

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetWeather(input json.RawMessage) (string, error) {
	var args GetWeatherInput
	if err := json.Unmarshal(input, &args); err != nil {
		return "", err
	}

	url := fmt.Sprintf("https://wttr.in/%s?format=3", args.City)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

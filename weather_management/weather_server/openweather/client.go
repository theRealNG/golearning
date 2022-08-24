package openweather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/theRealNG/golearning/weather_management/weather_server/citydb"
)

const (
	apiKey = "522e8973cd3c34f33bfd11d3e16668b2"
	apiURL = "https://api.openweathermap.org/data/2.5/weather?"
)

type WeatherResponse struct {
	Weather []WeatherResponseWeather
	Main    struct {
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		FeelsLike float64 `json:"feels_like"`
	}
}

type WeatherResponseWeather struct {
	Description string
}

func GetWeather(city *citydb.City, lang string) WeatherResponse {
	requestURL := apiURL + "lat=" + strconv.FormatFloat(float64(city.Lat), 'f', -1, 32) + "&lon=" + strconv.FormatFloat(float64(city.Long), 'f', -1, 32) + "&appid=" + apiKey + "&lang=" + lang

	resp, err := http.Get(requestURL)
	if err != nil {
		panic(fmt.Sprintf("Failed to get weather: %v", err))
	}
	// TODO: Check http status code
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(fmt.Sprintf("Failed to read response body: %v", err))
	}
	var weather WeatherResponse
	err = json.Unmarshal(body, &weather)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal the response body: %v", err))
	}
	return weather
}

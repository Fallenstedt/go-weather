package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type forecastReponse struct {
	Properties struct {
		Periods []Forecast `json:"periods"`
	} `json:"properties"`
}

type Forecast struct {
	Number                     uint8  `json:"number"`
	Name                       string `json:"name"`
	StartTime                  string `json:"startTime"`
	EndTime                    string `json:"endTime"`
	IsDayTime                  bool   `json:"isDayTime"`
	Temperature                int16  `json:"temperature"`
	TemperatureUnit            string `json:"temperatureUnit"`
	TemperatureTrend           string `json:"temperatureTrend"`
	ProbabilityOfPrecipitation struct {
		UnitCode string `json:"unitCode"`
		Value    uint8  `json:"value"`
	} `json:"probabilityOfPrecipitation"`
	Dewpoint struct {
		UnitCode string  `json:"unitCode"`
		Value    float32 `json:"value"`
	} `json:"dewpoint"`
	RelativeHumidity struct {
		UnitCode string `json:"unitCode"`
		Value    uint8  `json:"value"`
	} `json:"RelativeHumidity"`
	WindSpeed        string `json:"windSpeed"`
	WindDirection    string `json:"windDirection"`
	Icon             string `json:"icon"`
	ShortForecast    string `json:"shortForecast"`
	DetailedForecast string `json:"detailedForecast"`
}

type IWeather interface {
	FetchForecast() ([]Forecast, error)
}

type Weather struct {
	forecastUrl     string
	activeAlertsUrl string
}

func New(args struct {
	ForecastUrl     string
	ActiveAlertsUrl string
}) Weather {
	return Weather{
		forecastUrl:     args.ForecastUrl,
		activeAlertsUrl: args.ActiveAlertsUrl,
	}
}

func (w *Weather) FetchForecast() ([]Forecast, error) {
	resp, err := http.Get(w.forecastUrl)
	if err != nil {
		return nil, fmt.Errorf("%w, %w", ErrFetchForecast, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%w, got status %d", ErrFetchForecast, resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var fr forecastReponse
	err = json.Unmarshal(body, &fr)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response, %w", err)
	}

	return fr.Properties.Periods, nil
}

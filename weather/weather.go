package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)


type IWeather interface {
	FetchForecast() ([]Forecast, error)
	FetchAlerts() ([]Alerts, error)
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


func (w *Weather) FetchAlerts() ([]Alerts, error) {
	resp, err := http.Get(w.forecastUrl)
	if err != nil {
		return nil, fmt.Errorf("%w, %w", ErrFetchAlerts, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%w, got status %d", ErrFetchAlerts, resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var ar alertsResponse
	err = json.Unmarshal(body, &ar)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response, %w", err)
	}

	return ar.Alerts, nil
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

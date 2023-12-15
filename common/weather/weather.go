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
	var ar alertsResponse

	err := w.fetch(w.activeAlertsUrl, &ar)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response, %w", err)
	}

	return ar.Alerts, nil
}

func (w *Weather) FetchForecast() ([]Forecast, error) {
	var fr forecastReponse
	 err := w.fetch(w.forecastUrl, &fr)
		if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response, %w", err)
	}

	return fr.Properties.Periods, nil
}


func (w *Weather) fetch(url string, unmarshal interface{})  error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("%w, %w", ErrFetchForecast, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%w, got status %d", ErrFetchForecast, resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, unmarshal)
	if err != nil {
		return fmt.Errorf("failed to unmarshal response, %w", err)
	}

	return nil
}

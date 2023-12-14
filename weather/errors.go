package weather

import "errors"

var (
	ErrFetchAlerts = errors.New("failed to fetch alerts")
	ErrFetchForecast = errors.New("failed to fetch forecast")
)

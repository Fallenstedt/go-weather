package weather

import "errors"

var (
	ErrFetchForecast = errors.New("failed to fetch forecast")
)
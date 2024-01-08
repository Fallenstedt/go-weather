package weather_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Fallenstedt/weather/common/weather"
)

func TestFetchForcast(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{
			"@context": [
				"https://geojson.org/geojson-ld/geojson-context.jsonld",
				{
					"@version": "1.1",
					"wx": "https://api.weather.gov/ontology#",
					"geo": "http://www.opengis.net/ont/geosparql#",
					"unit": "http://codes.wmo.int/common/unit/",
					"@vocab": "https://api.weather.gov/ontology#"
				}
			],
			"type": "Feature",
			"geometry": {
				"type": "Polygon",
				"coordinates": [
					[
						[
							-122.83065089999999,
							45.492130299999999
						],
						[
							-122.8244613,
							45.471285000000002
						],
						[
							-122.794738,
							45.475621400000001
						],
						[
							-122.80092139999999,
							45.496467000000003
						],
						[
							-122.83065089999999,
							45.492130299999999
						]
					]
				]
			},
			"properties": {
				"updated": "2023-12-22T00:41:09+00:00",
				"units": "us",
				"forecastGenerator": "BaselineForecastGenerator",
				"generatedAt": "2023-12-22T00:51:08+00:00",
				"updateTime": "2023-12-22T00:41:09+00:00",
				"validTimes": "2023-12-21T18:00:00+00:00/P7DT13H",
				"elevation": {
					"unitCode": "wmoUnit:m",
					"value": 60.960000000000001
				},
				"periods": [
					{
						"number": 1,
						"name": "This Afternoon",
						"startTime": "2023-12-21T16:00:00-08:00",
						"endTime": "2023-12-21T18:00:00-08:00",
						"isDaytime": true,
						"temperature": 47,
						"temperatureUnit": "F",
						"temperatureTrend": null,
						"probabilityOfPrecipitation": {
							"unitCode": "wmoUnit:percent",
							"value": null
						},
						"dewpoint": {
							"unitCode": "wmoUnit:degC",
							"value": 7.2222222222222223
						},
						"relativeHumidity": {
							"unitCode": "wmoUnit:percent",
							"value": 96
						},
						"windSpeed": "2 mph",
						"windDirection": "SW",
						"icon": "https://api.weather.gov/icons/land/day/bkn?size=medium",
						"shortForecast": "Mostly Cloudy",
						"detailedForecast": "Mostly cloudy, with a high near 47. Southwest wind around 2 mph."
					},
					{
						"number": 2,
						"name": "Tonight",
						"startTime": "2023-12-21T18:00:00-08:00",
						"endTime": "2023-12-22T06:00:00-08:00",
						"isDaytime": false,
						"temperature": 40,
						"temperatureUnit": "F",
						"temperatureTrend": "rising",
						"probabilityOfPrecipitation": {
							"unitCode": "wmoUnit:percent",
							"value": 80
						},
						"dewpoint": {
							"unitCode": "wmoUnit:degC",
							"value": 6.666666666666667
						},
						"relativeHumidity": {
							"unitCode": "wmoUnit:percent",
							"value": 100
						},
						"windSpeed": "2 to 10 mph",
						"windDirection": "SSE",
						"icon": "https://api.weather.gov/icons/land/night/fog/rain,80?size=medium",
						"shortForecast": "Patchy Fog then Light Rain",
						"detailedForecast": "Patchy fog between 8pm and 2am, then rain. Cloudy. Low around 40, with temperatures rising to around 43 overnight. South southeast wind 2 to 10 mph. Chance of precipitation is 80%. New rainfall amounts between a tenth and quarter of an inch possible."
					}
				]
			}
		}`))
	}))
	defer server.Close()

	w := weather.New(struct {
		ForecastUrl     string
		ActiveAlertsUrl string
	}{
		ForecastUrl:     server.URL,
		ActiveAlertsUrl: server.URL,
	})

	_, err := w.FetchForecast()

	if err != nil {
		t.Errorf("Found error: %v", err)
	}
}

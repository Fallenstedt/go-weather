package weather

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

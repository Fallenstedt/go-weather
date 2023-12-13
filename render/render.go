package render

import (
	"fmt"
	"io"

	"github.com/Fallenstedt/weather/weather"
	"github.com/jedib0t/go-pretty/v6/table"
)

type Render struct {
	Out io.Writer
}

func (r *Render) RenderForecast(forecast *[]weather.Forecast) error {

	t := table.NewWriter()
	t.SetOutputMirror(r.Out)
	t.AppendHeader(table.Row{"#", "Day", "Temp", "Wind", "Precipitation", "Forecast"})

	for _, f := range *forecast {
		temp := fmt.Sprintf("%v%s", f.Temperature, f.TemperatureUnit)
		wind := fmt.Sprintf("%s %s", f.WindDirection, f.WindSpeed)
		precipitation := fmt.Sprintf("%d%s", f.ProbabilityOfPrecipitation.Value, "%")
		t.AppendRow(table.Row{f.Number, f.Name, temp, wind, precipitation, f.ShortForecast})
		t.AppendSeparator()
	}
	t.SetStyle(table.StyleColoredBlueWhiteOnBlack)
	t.Render()
	return nil
}

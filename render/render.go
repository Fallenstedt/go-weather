package render

import (
	"context"
	"fmt"
	"io"

	"github.com/Fallenstedt/weather/util"
	"github.com/Fallenstedt/weather/weather"
	"github.com/jedib0t/go-pretty/v6/table"
)

type Render struct {
	Out io.Writer
}

func (r *Render) RenderForecast(ctx context.Context, forecast *[]weather.Forecast) error {
	flags := util.GetFlagsFromContext(ctx)
	isDetailOnly := flags.Detail > 0 && flags.Detail <= len(*forecast)

	if isDetailOnly {
		f := (*forecast)[flags.Detail-1]
    fmt.Fprintln(r.Out, f.Name, " | ", r.getTemp(&f), " | ", fmt.Sprintf("Wind %s", r.getWind(&f) ) )
    fmt.Fprintln(r.Out, f.DetailedForecast)
	} else {
		t := table.NewWriter()
		t.SetOutputMirror(r.Out)
		t.AppendHeader(table.Row{"#", "Day", "Temp", "Wind", "Precipitation", "Forecast"})
		for _, f := range *forecast {
			t.AppendRow(r.buildRow(&f))
			t.AppendSeparator()
		}

	t.SetStyle(table.StyleRounded)
	t.Render()
	}

	return nil
}

func (r *Render) buildRow(f *weather.Forecast) table.Row {
	return table.Row{f.Number, f.Name, r.getTemp(f), r.getWind(f), r.getPrecipitation(f), f.ShortForecast}
}

func (r *Render) getTemp(f *weather.Forecast) string {
	return fmt.Sprintf("%v%s", f.Temperature, f.TemperatureUnit)
}

func (r *Render) getWind(f *weather.Forecast) string {
	return fmt.Sprintf("%s %s", f.WindDirection, f.WindSpeed)
}

func (r *Render) getPrecipitation(f *weather.Forecast) string {
	return fmt.Sprintf("%d%s", f.ProbabilityOfPrecipitation.Value, "%")
}

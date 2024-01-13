package render

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/Fallenstedt/weather/common/util"
	"github.com/Fallenstedt/weather/common/weather"
	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/v6/table"
)

type Render struct {
	Out io.Writer
}

func (r *Render) RenderAlerts(ctx context.Context, alerts *[]weather.Alerts) error {

	redBg := color.New(color.FgWhite, color.Bold).Add(color.BgRed)
	// yellowBg :=  color.New(color.FgBlack).Add(color.BgHiYellow)
	greenBg := color.New(color.FgBlack, color.Bold).Add(color.BgWhite)
	bold := color.New(color.Bold)

	if len(*alerts) == 0 {
		bold.Fprintf(r.Out, "%s\n\n", "No alerts at this time")
		return nil
	}

	for _, a := range *alerts {
		var colorFn *color.Color
		switch a.Properties.Severity {
		case "Minor":
			colorFn = greenBg
		case "Severe":
			colorFn = redBg
		default:
			colorFn = redBg
		}

		colorFn.Fprintln(r.Out, fmt.Sprintf("%s - %s", a.Properties.MessageType, a.Properties.Headline))
		fmt.Println()
		bold.Fprintf(r.Out, "%s\n\n", a.Properties.Event)
		fmt.Fprintf(r.Out, "%s\n\n", a.Properties.Description)
		fmt.Fprintf(r.Out, "%s\n\n", a.Properties.Instruction)
	}
	return nil
}

func (r *Render) RenderForecast(ctx context.Context, forecast *[]weather.Forecast) error {
	flags := util.GetFlagsFromContext(ctx)
	isDetailOnly := flags.Detail > 0 && flags.Detail <= len(*forecast)

	if isDetailOnly {
		f := (*forecast)[flags.Detail-1]
		fmt.Fprintln(r.Out, f.Name, " | ", r.getTemp(&f), " | ", fmt.Sprintf("Wind %s", r.getWind(&f)))
		fmt.Fprintln(r.Out, f.DetailedForecast)
	} else {
		t := table.NewWriter()
		t.SetOutputMirror(r.Out)
		t.AppendHeader(table.Row{"#", "Day", "Temp", "Wind", "Precipitation", "Forecast"})
		for _, f := range *forecast {
			t.AppendRow(r.buildRow(&f))
			t.AppendSeparator()
		}

		t.SetStyle(table.StyleColoredGreenWhiteOnBlack)
		t.Render()
	}

	return nil
}

func (r *Render) RenderRadar(ctx context.Context, url string) error {
	_, cancel := context.WithTimeout(ctx, time.Second * 5)
	defer cancel()
	return util.OpenBrowser(url)

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

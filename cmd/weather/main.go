package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/Fallenstedt/weather/common/render"
	"github.com/Fallenstedt/weather/common/util"
	"github.com/Fallenstedt/weather/common/weather"
)

func main() {

	flag.Usage = func() {
		fmt.Fprint(flag.CommandLine.Output(), "%s tool. Developed by Alex Fallenstedt\n", os.Args[0])
		fmt.Fprintln(flag.CommandLine.Output(), "Usage information:")
		flag.PrintDefaults()
	}
	detail := flag.Int("detail", 0, "The day number to get a detailed forecast for")
	flag.Parse()

	ctx := context.WithValue(context.Background(), util.ContextKeyFlags, util.Flags{Detail: *detail})

	if err := run(os.Stdout, ctx); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}

func run(out io.Writer, ctx context.Context) error {
	r := render.Render{Out: out}

	w := weather.New(struct {
		ForecastUrl     string
		ActiveAlertsUrl string
	}{
		ForecastUrl:     "https://api.weather.gov/gridpoints/PQR/108,103/forecast",
		ActiveAlertsUrl: "https://api.weather.gov/alerts/active?zone=ORC067",
	})

	alerts, _ := w.FetchAlerts()
	forecast, err := w.FetchForecast()

	if err != nil {
		return err
	}

	r.RenderForecast(ctx, &forecast)
	r.RenderAlerts(ctx, &alerts)

	return nil
}

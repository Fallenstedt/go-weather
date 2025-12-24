package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/Fallenstedt/weather/common/config"
	"github.com/Fallenstedt/weather/common/render"
	"github.com/Fallenstedt/weather/common/weather"
)

func main() {

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "%s tool. Developed by Alex Fallenstedt\n", os.Args[0])
		fmt.Fprintln(flag.CommandLine.Output(), "Usage information:")
		fmt.Fprintln(flag.CommandLine.Output(), "\nCommands:")
		fmt.Fprintln(flag.CommandLine.Output(), "  list        List configured locations (edit ~/.weathercli.yml to add/remove)")
		fmt.Fprintln(flag.CommandLine.Output(), "\nFlags:")
		flag.PrintDefaults()
	}
	// location flag controls which configured location to use
	var location string
	flag.StringVar(&location, "location", "", "location name from config (overrides default)")

	flag.Parse()


	if err := run(os.Stdout, context.Background(), location); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}
func run(out io.Writer, ctx context.Context, locationFlag string) error {
	// If no args and no location flag, show help and exit
	args := flag.Args()
	if len(args) == 0 && locationFlag == "" {
		flag.Usage()
		return nil
	}

	// Load config and handle "list" command or selected location
	cfg, err := config.LoadConfig()
	if err != nil {
		return err
	}

	// If first positional arg is "list", show available locations
	if len(args) > 0 && args[0] == "list" {
		fmt.Fprintln(out, "Configured locations:")
		for _, name := range config.ListLocations(cfg) {
			fmt.Fprintln(out, "-", name)
		}
		return nil
	}

	// Determine which location to use
	locName := locationFlag
	if locName == "" {
		// If flag wasn't provided, try env var or default from config
		locName = cfg.Default
	}

	loc, ok := config.GetLocation(cfg, locName)
	if !ok {
		return fmt.Errorf("location '%s' not found in config", locName)
	}

	r := render.Render{Out: out}

	w := weather.New(struct {
		ForecastUrl     string
		ActiveAlertsUrl string
	}{
		ForecastUrl:     loc.ForecastUrl,
		ActiveAlertsUrl: loc.ActiveAlertsUrl,

	})

	alerts, _ := w.FetchAlerts()
	forecast, err := w.FetchForecast()

	if err != nil {
		return err
	}

	r.RenderForecast(ctx, &forecast)
	r.RenderAlerts(ctx, &alerts)
	return r.RenderRadar(ctx, loc.RadarUrl)

}

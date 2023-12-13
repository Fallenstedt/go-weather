package main

import (
	"fmt"
	"io"
	"os"

	"github.com/Fallenstedt/weather/render"
	"github.com/Fallenstedt/weather/weather"
)

func main() {

	if err := run(os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}

func run(out io.Writer) error {
	r := render.Render{Out: out}

	w := weather.New("https://api.weather.gov/gridpoints/PQR/108,103/forecast", "https://api.weather.gov/alerts/active?zone=ORZ006")

	forecast, err := w.FetchForecast()

	if err != nil {
		return err
	}

	return r.RenderForecast(&forecast)
}

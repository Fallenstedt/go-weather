package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Location struct {
	Name            string `yaml:"name"`
	ForecastUrl     string `yaml:"forecast_url"`
	ActiveAlertsUrl string `yaml:"active_alerts_url"`
	RadarUrl        string `yaml:"radar_url"`
}

type Config struct {
	Locations map[string]Location `yaml:"locations"`
	Default   string              `yaml:"default_location"`
}

// Default configuration values with an initial `beaverton` location.
var defaultConfig = Config{
	Locations: map[string]Location{
		"beaverton": {
			Name:            "Beaverton",
			ForecastUrl:     "https://api.weather.gov/gridpoints/PQR/108,103/forecast",
			ActiveAlertsUrl: "https://api.weather.gov/alerts/active?zone=ORC067",
			RadarUrl:        "https://radar.weather.gov/station/krtx/standard",
		},
		"mthood":{
			Name: "Mt Hood",
			ForecastUrl: "https://api.weather.gov/gridpoints/PQR/143,89/forecast",
			ActiveAlertsUrl: "https://api.weather.gov/alerts/active?zone=ORZ126",
			RadarUrl: "https://radar.weather.gov/station/krtx/standard",
		},
	},
	Default: "beaverton",
}

  // mthood:
  //   name: Mt Hood
  //   forecast_url: https://api.weather.gov/gridpoints/PQR/143,89/forecast
  //   active_alerts_url: https://api.weather.gov/alerts/active?zone=ORZ126
  //   radar_url: https://radar.weather.gov/station/krtx/standard

// getConfigPath returns the path for the config file in the user's home directory.
func getConfigPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDir, ".weathercli.yml"), nil
}

// LoadConfig loads the config from file or creates a default one if it doesn't exist.
func LoadConfig() (*Config, error) {
	configPath, err := getConfigPath()
	if err != nil {
		return nil, err
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		// Config file does not exist; create it with default values.
		file, err := os.Create(configPath)
		if err != nil {
			return nil, err
		}
		defer file.Close()

		data, err := yaml.Marshal(defaultConfig)
		if err != nil {
			return nil, err
		}

		if _, err := file.Write(data); err != nil {
			return nil, err
		}

		fmt.Printf("Default config created at %s\n", configPath)
	}

	// Load existing config.
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	// Ensure Locations map is not nil
	if config.Locations == nil {
		config.Locations = map[string]Location{}
	}

	return &config, nil
}

// SaveConfig writes the given config to the default config path.
func SaveConfig(cfg *Config) error {
	if cfg == nil {
		return errors.New("nil config")
	}
	configPath, err := getConfigPath()
	if err != nil {
		return err
	}
	data, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}
	return os.WriteFile(configPath, data, 0644)
}

// ListLocations returns the keys (names) of configured locations in insertion order (map iteration is not ordered).
func ListLocations(cfg *Config) []string {
	if cfg == nil {
		return nil
	}
	names := make([]string, 0, len(cfg.Locations))
	for k := range cfg.Locations {
		names = append(names, k)
	}
	return names
}

// GetLocation returns a Location by name or the default location if name is empty.
func GetLocation(cfg *Config, name string) (Location, bool) {
	if cfg == nil {
		return Location{}, false
	}
	if name == "" {
		name = cfg.Default
	}
	loc, ok := cfg.Locations[name]
	return loc, ok
}

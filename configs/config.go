package configs

import (
	"errors"
	"github.com/spf13/viper"
)

type conf struct {
	WeatherApiKey string `mapstructure:"WEATHER_API_KEY"`
}

func LoadConfig() (*conf, error) {
	v := viper.New()
	v.AutomaticEnv()

	wak := v.GetString("WEATHER_API_KEY")
	if wak == "" {
		return nil, errors.New("WEATHER_API_KEY environment variable not set")
	}
	c := &conf{WeatherApiKey: wak}

	return c, nil
}

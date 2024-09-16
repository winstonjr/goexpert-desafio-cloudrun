package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGivenAValidWeather_WhenICallNewWeatherFunc_ThenIShouldReceiveWeatherWithAllTemperatures(t *testing.T) {
	weather := NewWeather(10.0)
	assert.Equal(t, 10.0, weather.TemperatureCelsius)
	assert.Equal(t, 50.0, weather.TemperatureFahrenheit)
	assert.Equal(t, 283.0, weather.TemperatureKelvin)
}

func TestGivenAValidWeather_WhenICallCalculateFahrenheitFunc_ThenIShouldReceiveWeatherWithTemperatureFahrenheit(t *testing.T) {
	weather := Weather{TemperatureCelsius: 10.0}
	assert.Equal(t, 10.0, weather.TemperatureCelsius)
	weather.CalculateFahrenheit()
	assert.Equal(t, 50.0, weather.TemperatureFahrenheit)
	assert.Equal(t, 0.0, weather.TemperatureKelvin)
}

func TestGivenAValidWeather_WhenICallCalculateKelvinFunc_ThenIShouldReceiveWeatherWithTemperatureKelvin(t *testing.T) {
	weather := Weather{TemperatureCelsius: 10.0}
	assert.Equal(t, 10.0, weather.TemperatureCelsius)
	weather.CalculateKelvin()
	assert.Equal(t, 283.0, weather.TemperatureKelvin)
	assert.Equal(t, 0.0, weather.TemperatureFahrenheit)
}

func TestGivenAValidWeather_WhenNothingIsCalculated_ThenIShouldReceiveWeatherWithZeroesInTemperature(t *testing.T) {
	weather := Weather{}
	assert.Equal(t, 0.0, weather.TemperatureCelsius)
	assert.Equal(t, 0.0, weather.TemperatureKelvin)
	assert.Equal(t, 0.0, weather.TemperatureFahrenheit)
}

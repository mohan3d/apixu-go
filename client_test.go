package apixu

import (
	"os"
	"testing"
)

var cities = []string{
	"Cairo",
	"London",
	"Paris",
	"Berlin",
	"New York",
}

func getAPIKey() string {
	return os.Getenv("APIXU_KEY")
}

func TestInvalidAPIKey(t *testing.T) {
	client := NewClient("Invalid_Key")
	_, err := client.Current("Paris")

	if err == nil {
		t.Error("Worked with invalid key")
	}
}

func TestCurrentWeatherValidCities(t *testing.T) {
	client := NewClient(getAPIKey())

	for _, city := range cities {
		_, err := client.Current(city)

		if err != nil {
			t.Errorf("There was an error getting current weather of %s: %v", city, err)
		}
	}
}

func TestCurrentWeatherInValidCity(t *testing.T) {
	client := NewClient(getAPIKey())
	_, err := client.Current("Unknown City")

	if err == nil {
		t.Errorf("No errors getting current weather of invalid city name")
	}
}

func TestForecastWeatherValidCities(t *testing.T) {
	days := []int{1, 5, 10}
	client := NewClient(getAPIKey())

	for _, day := range days {
		for _, city := range cities {
			_, err := client.Forecast(city, day)

			if err != nil {
				t.Errorf("There was an error getting forecast weather of %s days %d: %v", city, day, err)
			}
		}
	}
}

func TestForecastWeatherInValidCities(t *testing.T) {
	days := []int{1, 5, 10}
	client := NewClient(getAPIKey())

	for _, day := range days {
		_, err := client.Forecast("Unknown City", day)

		if err == nil {
			t.Errorf("No errors getting forecast weather of invalid city name")
		}
	}
}

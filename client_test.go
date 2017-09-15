package apixu

import (
	"os"
	"testing"
	"time"
)

const InvalidCityName = "INVALID-CITY-NAME"

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
	_, err := client.Current(InvalidCityName)

	if err == nil {
		t.Error("No errors getting current weather of invalid city name")
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

func TestForecastWeatherInValidCity(t *testing.T) {
	days := []int{1, 5, 10}
	client := NewClient(getAPIKey())

	for _, day := range days {
		_, err := client.Forecast(InvalidCityName, day)

		if err == nil {
			t.Error("No errors getting forecast weather of invalid city name")
		}
	}
}

func TestHistoryWeatherValidCities(t *testing.T) {
	yesterday := time.Now().AddDate(0, 0, -1)
	client := NewClient(getAPIKey())

	for _, city := range cities {
		_, err := client.History(city, yesterday.Format("2006 01 2"))

		if err != nil {
			t.Errorf("There was an error getting current weather of %s: %v", city, err)
		}
	}
}

func TestHistoryWeatherInValidCity(t *testing.T) {
	yesterday := time.Now().AddDate(0, 0, -1)
	client := NewClient(getAPIKey())

	_, err := client.History(InvalidCityName, yesterday.Format("2006 01 2"))

	if err == nil {
		t.Error("No errors getting history weather of invalid city name")
	}

}

func TestSearchValidCities(t *testing.T) {
	client := NewClient(getAPIKey())

	for _, city := range cities {
		_, err := client.Search(city)

		if err != nil {
			t.Errorf("There was an error getting current weather of %s: %v", city, err)
		}
	}
}

func TestSearchInValidCity(t *testing.T) {
	client := NewClient(getAPIKey())

	matchedCities, err := client.Search(InvalidCityName)

	if err != nil {
		t.Error(err)
	}

	if len(*matchedCities) != 0 {
		t.Error("Non-Empty array of matched cities for invalid city name")
	}
}

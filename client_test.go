package apixu

import (
	"os"
	"testing"
)

func getAPIKey() string {
	return os.Getenv("APIXU_KEY")
}

func TestCurrentWeather(t *testing.T) {

}

package apixu

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

const apiVersion = "v1"
const apiBaseURL = "http://api.apixu.com/" + apiVersion + "/"

// OptionalParam represents optional query parameters.
type OptionalParam struct {
	name  string
	value string
}

// CurrentWeather represents json returned by current.
type CurrentWeather struct {
	Location struct {
		Name           string  `json:"name"`
		Region         string  `json:"region"`
		Country        string  `json:"country"`
		Lat            float64 `json:"lat"`
		Lon            float64 `json:"lon"`
		TzID           string  `json:"tz_id"`
		LocaltimeEpoch int     `json:"localtime_epoch"`
		Localtime      string  `json:"localtime"`
	} `json:"location"`
	Current struct {
		LastUpdatedEpoch int     `json:"last_updated_epoch"`
		LastUpdated      string  `json:"last_updated"`
		TempC            float64 `json:"temp_c"`
		TempF            float64 `json:"temp_f"`
		IsDay            int     `json:"is_day"`
		Condition        struct {
			Text string `json:"text"`
			Icon string `json:"icon"`
			Code int    `json:"code"`
		} `json:"condition"`
		WindMph    float64 `json:"wind_mph"`
		WindKph    float64 `json:"wind_kph"`
		WindDegree int     `json:"wind_degree"`
		WindDir    string  `json:"wind_dir"`
		PressureMb float64 `json:"pressure_mb"`
		PressureIn float64 `json:"pressure_in"`
		PrecipMm   float64 `json:"precip_mm"`
		PrecipIn   float64 `json:"precip_in"`
		Humidity   int     `json:"humidity"`
		Cloud      int     `json:"cloud"`
		FeelslikeC float64 `json:"feelslike_c"`
		FeelslikeF float64 `json:"feelslike_f"`
		VisKm      float64 `json:"vis_km"`
		VisMiles   float64 `json:"vis_miles"`
	} `json:"current"`
}

// ForecastWeather represents json returned by forecast.
type ForecastWeather struct {
	Location struct {
		Name           string  `json:"name"`
		Region         string  `json:"region"`
		Country        string  `json:"country"`
		Lat            float64 `json:"lat"`
		Lon            float64 `json:"lon"`
		TzID           string  `json:"tz_id"`
		LocaltimeEpoch int     `json:"localtime_epoch"`
		Localtime      string  `json:"localtime"`
	} `json:"location"`
	Current struct {
		LastUpdatedEpoch int     `json:"last_updated_epoch"`
		LastUpdated      string  `json:"last_updated"`
		TempC            float64 `json:"temp_c"`
		TempF            float64 `json:"temp_f"`
		IsDay            int     `json:"is_day"`
		Condition        struct {
			Text string `json:"text"`
			Icon string `json:"icon"`
			Code int    `json:"code"`
		} `json:"condition"`
		WindMph    float64 `json:"wind_mph"`
		WindKph    float64 `json:"wind_kph"`
		WindDegree int     `json:"wind_degree"`
		WindDir    string  `json:"wind_dir"`
		PressureMb float64 `json:"pressure_mb"`
		PressureIn float64 `json:"pressure_in"`
		PrecipMm   float64 `json:"precip_mm"`
		PrecipIn   float64 `json:"precip_in"`
		Humidity   int     `json:"humidity"`
		Cloud      int     `json:"cloud"`
		FeelslikeC float64 `json:"feelslike_c"`
		FeelslikeF float64 `json:"feelslike_f"`
		VisKm      float64 `json:"vis_km"`
		VisMiles   float64 `json:"vis_miles"`
	} `json:"current"`
	Forecast struct {
		Forecastday []struct {
			Date      string `json:"date"`
			DateEpoch int    `json:"date_epoch"`
			Day       struct {
				MaxtempC      float64 `json:"maxtemp_c"`
				MaxtempF      float64 `json:"maxtemp_f"`
				MintempC      float64 `json:"mintemp_c"`
				MintempF      float64 `json:"mintemp_f"`
				AvgtempC      float64 `json:"avgtemp_c"`
				AvgtempF      float64 `json:"avgtemp_f"`
				MaxwindMph    float64 `json:"maxwind_mph"`
				MaxwindKph    float64 `json:"maxwind_kph"`
				TotalprecipMm float64 `json:"totalprecip_mm"`
				TotalprecipIn float64 `json:"totalprecip_in"`
				AvgvisKm      float64 `json:"avgvis_km"`
				AvgvisMiles   float64 `json:"avgvis_miles"`
				Avghumidity   float64 `json:"avghumidity"`
				Condition     struct {
					Text string `json:"text"`
					Icon string `json:"icon"`
					Code int    `json:"code"`
				} `json:"condition"`
				Uv float64 `json:"uv"`
			} `json:"day"`
			Astro struct {
				Sunrise  string `json:"sunrise"`
				Sunset   string `json:"sunset"`
				Moonrise string `json:"moonrise"`
				Moonset  string `json:"moonset"`
			} `json:"astro"`
			Hour []struct {
				TimeEpoch int     `json:"time_epoch"`
				Time      string  `json:"time"`
				TempC     float64 `json:"temp_c"`
				TempF     float64 `json:"temp_f"`
				IsDay     int     `json:"is_day"`
				Condition struct {
					Text string `json:"text"`
					Icon string `json:"icon"`
					Code int    `json:"code"`
				} `json:"condition"`
				WindMph      float64 `json:"wind_mph"`
				WindKph      float64 `json:"wind_kph"`
				WindDegree   int     `json:"wind_degree"`
				WindDir      string  `json:"wind_dir"`
				PressureMb   float64 `json:"pressure_mb"`
				PressureIn   float64 `json:"pressure_in"`
				PrecipMm     float64 `json:"precip_mm"`
				PrecipIn     float64 `json:"precip_in"`
				Humidity     int     `json:"humidity"`
				Cloud        int     `json:"cloud"`
				FeelslikeC   float64 `json:"feelslike_c"`
				FeelslikeF   float64 `json:"feelslike_f"`
				WindchillC   float64 `json:"windchill_c"`
				WindchillF   float64 `json:"windchill_f"`
				HeatindexC   float64 `json:"heatindex_c"`
				HeatindexF   float64 `json:"heatindex_f"`
				DewpointC    float64 `json:"dewpoint_c"`
				DewpointF    float64 `json:"dewpoint_f"`
				WillItRain   int     `json:"will_it_rain"`
				ChanceOfRain string  `json:"chance_of_rain"`
				WillItSnow   int     `json:"will_it_snow"`
				ChanceOfSnow string  `json:"chance_of_snow"`
				VisKm        float64 `json:"vis_km"`
				VisMiles     float64 `json:"vis_miles"`
			} `json:"hour"`
		} `json:"forecastday"`
	} `json:"forecast"`
}

// HistoryWeather represents json returned by history.
type HistoryWeather struct {
	Location struct {
		Name           string  `json:"name"`
		Region         string  `json:"region"`
		Country        string  `json:"country"`
		Lat            float64 `json:"lat"`
		Lon            float64 `json:"lon"`
		TzID           string  `json:"tz_id"`
		LocaltimeEpoch int     `json:"localtime_epoch"`
		Localtime      string  `json:"localtime"`
	} `json:"location"`
	Forecast struct {
		Forecastday []struct {
			Date      string `json:"date"`
			DateEpoch int    `json:"date_epoch"`
			Day       struct {
				MaxtempC      float64 `json:"maxtemp_c"`
				MaxtempF      float64 `json:"maxtemp_f"`
				MintempC      float64 `json:"mintemp_c"`
				MintempF      float64 `json:"mintemp_f"`
				AvgtempC      float64 `json:"avgtemp_c"`
				AvgtempF      float64 `json:"avgtemp_f"`
				MaxwindMph    float64 `json:"maxwind_mph"`
				MaxwindKph    float64 `json:"maxwind_kph"`
				TotalprecipMm float64 `json:"totalprecip_mm"`
				TotalprecipIn float64 `json:"totalprecip_in"`
				AvgvisKm      float64 `json:"avgvis_km"`
				AvgvisMiles   float64 `json:"avgvis_miles"`
				Avghumidity   float64 `json:"avghumidity"`
				Condition     struct {
					Text string `json:"text"`
					Icon string `json:"icon"`
					Code int    `json:"code"`
				} `json:"condition"`
				Uv float64 `json:"uv"`
			} `json:"day"`
			Astro struct {
				Sunrise  string `json:"sunrise"`
				Sunset   string `json:"sunset"`
				Moonrise string `json:"moonrise"`
				Moonset  string `json:"moonset"`
			} `json:"astro"`
			Hour []struct {
				TimeEpoch int     `json:"time_epoch"`
				Time      string  `json:"time"`
				TempC     float64 `json:"temp_c"`
				TempF     float64 `json:"temp_f"`
				IsDay     int     `json:"is_day"`
				Condition struct {
					Text string `json:"text"`
					Icon string `json:"icon"`
					Code int    `json:"code"`
				} `json:"condition"`
				WindMph      float64 `json:"wind_mph"`
				WindKph      float64 `json:"wind_kph"`
				WindDegree   int     `json:"wind_degree"`
				WindDir      string  `json:"wind_dir"`
				PressureMb   float64 `json:"pressure_mb"`
				PressureIn   float64 `json:"pressure_in"`
				PrecipMm     float64 `json:"precip_mm"`
				PrecipIn     float64 `json:"precip_in"`
				Humidity     int     `json:"humidity"`
				Cloud        int     `json:"cloud"`
				FeelslikeC   float64 `json:"feelslike_c"`
				FeelslikeF   float64 `json:"feelslike_f"`
				WindchillC   float64 `json:"windchill_c"`
				WindchillF   float64 `json:"windchill_f"`
				HeatindexC   float64 `json:"heatindex_c"`
				HeatindexF   float64 `json:"heatindex_f"`
				DewpointC    float64 `json:"dewpoint_c"`
				DewpointF    float64 `json:"dewpoint_f"`
				WillItRain   int     `json:"will_it_rain"`
				ChanceOfRain string  `json:"chance_of_rain"`
				WillItSnow   int     `json:"will_it_snow"`
				ChanceOfSnow string  `json:"chance_of_snow"`
				VisKm        float64 `json:"vis_km"`
				VisMiles     float64 `json:"vis_miles"`
			} `json:"hour"`
		} `json:"forecastday"`
	} `json:"forecast"`
}

// MatchingCities represents json returned by search.
type MatchingCities []struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Region  string  `json:"region"`
	Country string  `json:"country"`
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
	URL     string  `json:"url"`
}

type errorResponse struct {
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

// Client represents apixu client.
type Client struct {
	apiKey string
}

// Current returns CurrentWeather obj representing current weather status.
func (client *Client) Current(q string, optionalParams ...OptionalParam) (*CurrentWeather, error) {
	url, err := client.getURL("current.json", q, optionalParams...)

	if err != nil {
		return nil, err
	}

	body, err := request(url)

	if err != nil {
		return nil, err
	}

	var currentWeather CurrentWeather

	if err := json.Unmarshal(body, &currentWeather); err != nil {
		return nil, err
	}

	return &currentWeather, nil
}

// Forecast returns ForecastWeather obj representing Forecast status.
func (client *Client) Forecast(q string, days int, optionalParams ...OptionalParam) (*ForecastWeather, error) {
	optionalParams = append(optionalParams, OptionalParam{"days", string(days)})
	url, err := client.getURL("forecast.json", q, optionalParams...)

	if err != nil {
		return nil, err
	}

	body, err := request(url)

	if err != nil {
		return nil, err
	}

	var forecastWeather ForecastWeather

	if err := json.Unmarshal(body, &forecastWeather); err != nil {
		return nil, err
	}

	return &forecastWeather, nil
}

// History returns HistoryWeather obj representing History status.
func (client *Client) History(q string, dt string, optionalParams ...OptionalParam) (*HistoryWeather, error) {
	optionalParams = append(optionalParams, OptionalParam{"dt", dt})
	url, err := client.getURL("history.json", q)

	if err != nil {
		return nil, err
	}

	body, err := request(url)

	if err != nil {
		return nil, err
	}

	var historyWeather HistoryWeather

	if err := json.Unmarshal(body, &historyWeather); err != nil {
		return nil, err
	}

	return &historyWeather, nil
}

// Search returns MatchingCities obj representing a list of matched cities.
func (client *Client) Search(q string) (*MatchingCities, error) {
	url, err := client.getURL("search.json", q)

	if err != nil {
		return nil, err
	}

	body, err := request(url)

	if err != nil {
		return nil, err
	}

	var matchingCities MatchingCities

	if err := json.Unmarshal(body, &matchingCities); err != nil {
		return nil, err
	}

	return &matchingCities, nil
}

func (client *Client) getURL(path string, q string, optionalParams ...OptionalParam) (string, error) {
	baseURL, err := url.Parse(apiBaseURL)

	if err != nil {
		return "", err
	}

	pathURL, err := url.Parse(path)

	if err != nil {
		return "", err
	}

	URL := baseURL.ResolveReference(pathURL)
	query := URL.Query()
	query.Set("key", client.apiKey)
	query.Set("q", q)

	URL.RawQuery = query.Encode()

	return URL.String(), nil
}

// NewClient Creates new client and returns a ref.
func NewClient(apiKey string) *Client {
	client := &Client{apiKey: apiKey}
	return client
}

func request(url string) ([]byte, error) {
	response, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	// Validate response.
	if response.StatusCode != 200 {
		var errorJSON errorResponse

		if err := json.Unmarshal(body, &errorJSON); err != nil {
			return nil, err
		}

		return nil, errors.New(errorJSON.Error.Message)
	}

	return body, err
}

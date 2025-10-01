package weather_test

import (
	"go-demo-5/geo"
	"go-demo-5/weather"
	"strings"
	"testing"
)

func TestGetWeather(t *testing.T) {
	var geo *geo.GeoData = &geo.GeoData{
		City: "london",
	}
	var format int = 3
	expexted := "london"

	got := weather.GetWeather(*geo, format)

	if !strings.Contains(got, expexted) {
		t.Errorf("Ожидалось %v, получено %v", expexted, got)
	}

}

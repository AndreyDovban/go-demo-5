package weather

import (
	"demo/weather-app/geo"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

func GetWeather(geo geo.GeoData, format int) string {
	log.Println(geo.City, format)
	baseUrl, err := url.Parse("https://wttr.in/" + geo.City)
	if err != nil {
		log.Println(err.Error())
		return ""
	}
	params := url.Values{}
	params.Add("format", fmt.Sprint(format))
	baseUrl.RawQuery = params.Encode()

	req, err := http.NewRequest("GET", baseUrl.String(), nil)
	if err != nil {
		log.Println(err.Error())
		return ""
	}
	req.Header.Set("User-Agent", "Mozilla/5.0")

	resp, err := new(http.Client).Do(req)
	if err != nil {
		log.Println(err.Error())
		return ""
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Println(resp.StatusCode)
		return ""
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
		return ""
	}

	return string(body)
}

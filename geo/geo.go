package geo

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

type GeoData struct {
	City string `json:"city"`
}

func GetMyLocation(city string) (*GeoData, error) {
	if city != "" {
		return &GeoData{
			City: city,
		}, nil
	}

	req, err := http.NewRequest("GET", "https://ipapi.co/json/", nil)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0")

	resp, err := new(http.Client).Do(req)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New("NOT_200")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var geo GeoData
	err = json.Unmarshal(body, &geo)
	if err != nil {
		return nil, err
	}

	return &geo, nil
}

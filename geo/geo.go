package geo

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

type GeoData struct {
	City string `json:"city"`
}

type CityPopulationResponse struct {
	Error bool `json:"error"`
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

func ChekCity(city string) bool {
	postBody, _ := json.Marshal(map[string]string{
		"city": city,
	})
	req, err := http.NewRequest("POST", "https://countriesnow.space/api/v0.1/countries/population/cities", bytes.NewBuffer(postBody))
	if err != nil {
		log.Println(err.Error())
		return false
	}
	req.Header.Set("User-Agent", "Mozilla/5.0")

	resp, err := new(http.Client).Do(req)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	// if errr != nil {
	// 	return false
	// }

	var populationResponse CityPopulationResponse
	json.Unmarshal(body, &populationResponse)

	return true
}

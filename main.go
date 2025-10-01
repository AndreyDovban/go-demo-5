package main

import (
	"flag"
	"fmt"
	"go-demo-5/geo"
	"go-demo-5/weather"
)

func main() {
	city := flag.String("city", "", "Город пользователя")

	flag.Parse()

	// fmt.Println(*city)

	geoData, err := geo.GetMyLocation(*city)

	if err != nil {
		fmt.Println(err.Error())
	}

	// fmt.Println(geoData)

	weather := weather.GetWeather(*geoData, 3)
	fmt.Println(weather)
}

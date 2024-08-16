package main

import (
	"demo/weather/geo"
	"demo/weather/weather"
	"flag"
	"log"
)

func main() {
	log.Println("start")
	city := flag.String("city", "", "Город пользователя")
	format := flag.Int("format", 1, "Формат ыввода данных")

	flag.Parse()

	log.Println(*city)

	geoData, err := geo.GetMyLocation(*city)
	if err != nil {
		log.Println(err.Error())
	}

	log.Println(geoData)

	weathrData := weather.GetWeather(*geoData, *format)

	log.Println(weathrData)
}

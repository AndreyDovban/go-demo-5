package geo_test

import (
	"demo/weather-app/geo"
	"testing"
)

func TestGetMyLocation(t *testing.T) {
	// Arrange - подготовка, expected результата, данные функции
	city := "London"
	expected := geo.GeoData{
		City: "London",
	}

	// Act - выполняем функцию
	got, err := geo.GetMyLocation(city)

	// Assert - проверка результат с expect
	if err != nil {
		t.Error("Ошибка получения города")
	}
	if got.City != expected.City {
		t.Errorf("Ожидальсь %v, получено %v", expected, got.City)
	}

}

func TestGetMyLocation_negative(t *testing.T) {
	// Arrange - подготовка, expected результата, данные функции
	city := "London"
	expected := geo.GeoData{
		City: "London",
	}

	// Act - выполняем функцию
	got, err := geo.GetMyLocation(city)

	// Assert - проверка результат с expect
	if err != nil {
		t.Error("Ошибка получения города")
	}
	if got.City != expected.City {
		t.Errorf("Ожидальсь %v, получено %v", expected, got.City)
	}

}

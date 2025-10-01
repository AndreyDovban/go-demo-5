package geo_test

import (
	"go-demo-5/geo"
	"testing"
)

func TestGetMyLocation(t *testing.T) {
	// Arrange - подготовка, expected результат
	city := "London"
	expected := geo.GeoData{
		City: "London",
	}

	// Act - выполняем функцию
	got, err := geo.GetMyLocation(city)

	// Assert - провека результата с expected
	if err != nil {
		t.Error("Ошибка получения города")
	}
	if got.City != expected.City {
		t.Errorf("Ожидалось %v, получено %v", expected, got)
	}

}

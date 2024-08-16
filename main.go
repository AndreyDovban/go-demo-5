package main

import (
	"flag"
	"log"
)

func main() {
	log.Println("start")
	name := flag.String("name", "Anton", "Имя пользователя")
	age := flag.Int("age", 18, "Возраст пользователя")

	flag.Parse()

	log.Println(*name, *age)
}

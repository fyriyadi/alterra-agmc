package main

import (
	"day2/config"
	"day2/route"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}

	config.InitDB()
	e := route.New()
	e.Logger.Fatal(e.Start(":8080"))
}

package main

import (
	"booking-system/lvl1/application"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	app := application.New()
	app.Run(":8080")
}

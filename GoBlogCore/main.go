package main

import (
	"GoBlogCore/configs"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	configs.ConnectDB()

	//routes.UserRoute(app)
	err := app.Listen(":" + configs.HttpPort)
	if err != nil {
		panic(err)
	}
}

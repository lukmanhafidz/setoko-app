package main

import (
	"fmt"
	"log"
	"setokoapp/domain/model"
	"setokoapp/infrastructure/persistence"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/jinzhu/configor"
)

func init() {
	log.SetFlags(log.Ldate | log.Lshortfile) //set up logger to show date and file location

	err := configor.New(&configor.Config{}).Load(&model.Config, "config.yml") //load config from config.yml
	if err != nil {
		log.Print("Load config error: ", err)
		return
	}
}

func main() {
	app := fiber.New(fiber.Config{ //set up fiber
		BodyLimit: 5 * 1024 * 1024, //limit 5 MB
	})

	app.Use(cors.New(cors.ConfigDefault)) //using default cors config to fiber app

	db, err := persistence.ConnectDb()
	if err != nil {
		log.Printf("Error when try connect to db: %v", err)
		return
	}

	setoko := app.Group("setoko-app")
	setoko.Get("healthcheck", func(ctx *fiber.Ctx) error { //to check if api is available
		return model.ResponseOk(ctx, nil)
	})

	// setoko.Get("/v1/generate-receipt/:orderId", middleware.AuthJwt)
	if err := app.Listen(fmt.Sprintf(":%d", model.Config.Port)); err != nil {
		log.Print("Load config error: ", err)
		return
	}
}

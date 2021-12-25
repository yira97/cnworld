package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/yira97/cnworld/app/assitant/middleware/auth"
	"github.com/yira97/cnworld/app/assitant/options"
)

var app *fiber.App

func Register() *fiber.App {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New())

	app.Use(auth.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	api := app.Group("/api")

	v1 := api.Group("/v1")

	v1.Post("/users/actions/register", CreateUser)
	v1.Post("/users/actions/login", LoginUser)
	v1.Get("/user/user::value", GetUser)
	v1.Get("/user/uid::uid/dashboard", GetUserDashboard)
	v1.Put("/user/uid::uid", UpdateUser)
	v1.Post("/applications/actions/create", CreateApplication)

	return app
}

func Setup() {
	app = Register()
}

func Start() {

	app.Listen(fmt.Sprintf(":%s", options.GetPort()))
}

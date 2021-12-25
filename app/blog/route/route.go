package route

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/yira97/cnworld/app/blog/config"
	"github.com/yira97/cnworld/app/blog/route/api/content"
	"github.com/yira97/cnworld/app/blog/route/api/content/comment"
	"github.com/yira97/cnworld/app/blog/route/api/message"
	"github.com/yira97/cnworld/app/blog/route/api/user"
)

var app *fiber.App

func Setup() {

	app = fiber.New()

	app.Use(logger.New())

	// Default config
	// https://docs.gofiber.io/api/middleware/cors
	app.Use(cors.New())

	api := app.Group("/api")

	v1 := api.Group("/v1")

	v1.Post("/temp/users", user.CreateTempUser)

	v1.Post("/messages", message.CreateMessage)
	v1.Get("/messages", message.ListMessages)
	v1.Get("/message/uid::uid", message.GetMessage)
	v1.Post("/message/uid::uid/op-read", message.ReadMessage)

	v1.Get("/contents", content.ListContents)
	v1.Post("/contents", content.CreateContent)
	v1.Get("/content/uid::uid", content.GetContent)
	v1.Get("/content/uid::uid/comments", comment.ListComments)
	// v1.Get("/content/uid::uid/comments", comment.CreateComment)

	// v1.Get("/video-course-reviews")
	// v1.Get("/video-course-review/uid::uid")
	// v1.Post("/video-course-reviews")
}

func Start() {
	app.Listen(fmt.Sprintf(":%s", config.GetPort()))
}

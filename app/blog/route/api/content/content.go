package content

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yira97/cnworld/app/blog/service/content"
	"github.com/yira97/cnworld/lib/base/short/s_fiber"
)

func ListContents(c *fiber.Ctx) error {
	count := s_fiber.QueryInt(c, "count", 100)
	offset := s_fiber.QueryInt(c, "offset", 0)

	res, err := content.ListContents(c.Context(), count, offset)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(res)
}

func CreateContent(c *fiber.Ctx) error {
	m := new(content.ContentView)

	if err := c.BodyParser(m); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	m.Clean4Create()
	created, err := content.CreateContent(c.Context(), *m)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(created)
}

func GetContent(c *fiber.Ctx) error {
	contentUID := c.Params("uid", "")

	res := content.GetContent(c.Context(), contentUID)
	if res == nil {
		return c.Status(401).SendString("no content_uid")
	}
	return c.JSON(res)
}

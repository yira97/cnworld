package comment

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yira97/cnworld/app/blog/service/content"
	"github.com/yira97/cnworld/lib/base/short/s_fiber"
)

func ListComments(c *fiber.Ctx) error {
	count := s_fiber.QueryInt(c, "count", 100)
	offset := s_fiber.QueryInt(c, "offset", 0)
	contentUID := c.Params("uid", "")
	if contentUID == "" {
		return c.Status(401).SendString("no content_uid")
	}

	res, err := content.ListComment(c.Context(), contentUID, count, offset)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(res)
}

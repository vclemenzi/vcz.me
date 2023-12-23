package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vclemenzi/vcz.me/utils"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("https://vclemenzi.dev", 301)
	})

	app.Get("/:redirect", func(c *fiber.Ctx) error {
		redirect := c.Params("redirect")
		url, err := utils.GetRedirect(redirect)

		if err != nil {
			c.Status(404)
			return c.JSON(fiber.Map{
				"error": "Redirect not found",
			})
		}

		return c.Redirect(url, 301)
	})

	app.Listen(":3000")
}

package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("นี่คือหัวหน้าเฟิร์ส")
	})

	app.Get("/dev-pug", func(c *fiber.Ctx) error {
		return c.SendString("นี่คือสมาชิก ปั๊ก")
	})

	app.Get("/dev-brand", func(c *fiber.Ctx) error {
		return c.SendString("นี่คือ กีกี้ หมายเลข 1 (แบรนด์)")
	})

	app.Get("/dev-fa", func(c *fiber.Ctx) error {
		return c.SendString("fa: นี้คือสมาชิก")
	})
	app.Get("/dev-real", func(c *fiber.Ctx) error {
		return c.SendString("นี่คือมือขวาของ inw005")
	})

	app.Listen(":3000")
}

package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shahriarsohan/go-blog-practise/utils"
)

func IsAuthenticated(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	_, err := utils.ParseJwt(cookie)

	if err != nil {
		c.Status(401)
		return c.JSON(fiber.Map{
			"msg": " Unauthorized Request",
		})
	}

	return c.Next()
}

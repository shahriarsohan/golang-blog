package controllers

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/shahriarsohan/go-blog-practise/database"
	"github.com/shahriarsohan/go-blog-practise/models"
)

func validateEmail(email string) bool {
	Re := regexp.MustCompile(`[a-z0-9. %+/-]+@[a-z0-9. %+/-]+\.[a-z0-9. %+/-]`)

	return Re.MatchString(email)
}

func Register(c *fiber.Ctx) error {
	var data map[string]interface{}
	var userData models.User

	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Unavle to pass body")
	}

	// Check password length
	if len(data["password"].(string)) <= 6 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"msg": "Password must be greater than 6 character",
		})
	}

	if !validateEmail(strings.TrimSpace(data["email"].(string))) {
		c.Status(400)
		return c.JSON(fiber.Map{
			"msg": "Invalid Email Addr",
		})
	}

	// Check if email already exists in DB
	database.DB.Where("email=?", strings.TrimSpace(data["email"].(string))).First(&userData)
	if userData.ID != 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"msg": "Email already exists",
		})
	}

	user := models.User{
		Fname: data["first_name"].(string),
		Lname: data["last_name"].(string),
		Email: strings.TrimSpace(data["email"].(string)),
		Phone: data["phone"].(string),
	}

	user.SetPassword(data["password"].(string))

	err := database.DB.Create(&user)
	if err != nil {
		log.Println(err)

	}

	c.Status(200)
	return c.JSON(fiber.Map{
		"user": user,
		"msg":  "account created successfully",
	})
}

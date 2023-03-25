package controllers

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/shahriarsohan/go-blog-practise/database"
	"github.com/shahriarsohan/go-blog-practise/models"
	"github.com/shahriarsohan/go-blog-practise/utils"
	"gorm.io/gorm"
)

func CreatePost(c *fiber.Ctx) error {
	var blogpost models.Post

	if err := c.BodyParser(&blogpost); err != nil {
		fmt.Println("Unable to parse body")
	}

	if err := database.DB.Create(&blogpost).Error; err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"msg": "Invalid payload",
		})

	}
	c.Status(200)
	return c.JSON(fiber.Map{
		"msg": "post added :)",
	})
}

func AllPost(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1")) // Get ?page= , 1 is the default value here
	limit := 5
	offset := (page - 1) * limit

	var total int64
	var getBlog []models.Post

	database.DB.Preload("User").Offset(offset).Limit(limit).Find(&getBlog)

	database.DB.Model(&models.Post{}).Count(&total)

	return c.JSON(fiber.Map{
		"data": getBlog,
		"meta": fiber.Map{
			"total": total,
			"page":  page,
			// "lastPage": math.Ceil(float64(int(total) / limit)),
		},
	})
}

func PostDetails(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var post models.Post

	database.DB.Where("id=?", id).Preload("User").First(&post)

	return c.JSON(fiber.Map{
		"data": post,
	})
}

func UpdatePost(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	blog := models.Post{
		ID: uint(id),
	}

	if err := c.BodyParser(&blog); err != nil {
		fmt.Println("Unable to parse body")
	}

	database.DB.Model(&blog).Updates(blog)

	return c.JSON(blog)
}

func GetUserPost(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	id, _ := utils.ParseJwt(cookie)
	var blog []models.Post

	database.DB.Model(&blog).Where("author_id=?", id).Preload("User").Find(&blog)

	return c.JSON(blog)
}

func DeletePost(c *fiber.Ctx) error {

	id, _ := strconv.Atoi(c.Params("id"))

	blog := models.Post{
		ID: uint(id),
	}

	deleteQuery := database.DB.Delete(&blog)
	if errors.Is(deleteQuery.Error, gorm.ErrRecordNotFound) {
		c.Status(404)
		c.JSON(fiber.Map{
			"msg": "not found",
		})
	}

	return c.JSON(fiber.Map{
		"msg": "Post deleted",
	})
}

package controllers

import (
	"math/rand"

	"github.com/gofiber/fiber/v2"
)

var letters = []rune("abcedfghijklmopqrstuvwxyz")

func randLetters(n int) string {
	b := make([]rune, n)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func Upload(ctx *fiber.Ctx) error {
	form, err := ctx.MultipartForm()

	if err != nil {
		return err
	}

	files := form.File["image"]
	fileName := ""

	for _, file := range files {
		fileName = randLetters(5) + " " + file.Filename
		if err := ctx.SaveFile(file, "./uploads/"+fileName); err != nil {
			return nil
		}
	}
	return ctx.JSON(fiber.Map{
		"url": "http://localhost:3000/api/uploads/" + fileName,
	})
}

package controllers

import (

	// "github.com/binesh/gomvc/helpers"
	"github.com/binesh/gomvc/initalizer"
	"github.com/binesh/gomvc/models"
	"github.com/gofiber/fiber/v2"
)

// function first letter must be capital
func PostIndex(c *fiber.Ctx) error {

	var post []models.Post

	initalizer.DB.Find(&post)
	// var a = helpers.GetData(post)
	// fmt.Println(a)
	return c.Status(200).JSON(post)
}

func PostCreate(c *fiber.Ctx) error {
	var post = models.Post{
		Title:       "test",
		Body:        "test",
		Description: "test",
		Slug:        "test	",
	}
	initalizer.DB.Create(&post)
	return c.Status(200).JSON(post)
}

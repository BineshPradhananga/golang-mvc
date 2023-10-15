package controllers

import (
	"fmt"

	"github.com/binesh/gomvc/helpers"
	"github.com/binesh/gomvc/initalizer"
	"github.com/binesh/gomvc/models"
	"github.com/gofiber/fiber/v2"
)

// function first letter must be capital
func PostIndex(c *fiber.Ctx) error {

	var post []models.Post

	initalizer.DB.Find(&post)
	var response = helpers.GetData(post)
	statusCode := response.(map[string]interface{})["statusCode"]
	return c.Status(statusCode.(int)).JSON(response)
}

func PostCreate(c *fiber.Ctx) error {
	var post = models.Post{
		Title:       "test",
		Body:        "test",
		Description: "test",
		Slug:        "test	",
	}
	initalizer.DB.Create(&post)
	var response = helpers.SuccessMessage("Post Created")
	fmt.Println(response)

	statusCode := response.(map[string]interface{})["statusCode"]
	return c.Status(statusCode.(int)).JSON(response)
	// statusCode := response.statusCode
	// return c.Status(statusCode).JSON(response)
	// return c.Status(200).JSON(response)
}

package controllers

import (
	"strconv"

	"github.com/binesh/gomvc/helpers"
	"github.com/binesh/gomvc/initalizer"
	"github.com/binesh/gomvc/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

// function first letter must be capital
func PostIndex(c *fiber.Ctx) error {

	var post []models.Post

	initalizer.DB.Find(&post)
	var response = helpers.GetData(post, "")
	statusCode := response.(map[string]interface{})["statusCode"]
	return c.Status(statusCode.(int)).JSON(response)
}

func PostCreate(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil //using the SecretKey which was generated in th Login function
	})

	if err != nil {

		var response = helpers.UnauthorizedMessage("Error while Authenticating!!")
		statusCode := response.(map[string]interface{})["statusCode"]
		return c.Status(statusCode.(int)).JSON(response)
	}

	claims := token.Claims.(*jwt.StandardClaims)

	// Convert the claims.Issuer (string) to a uint
	userID, err := strconv.ParseUint(claims.Issuer, 10, 64)
	if err != nil {
		// Handle the error, e.g., log it and return an error response
		var response = helpers.ErrorMessage("Error while converting user ID.")
		statusCode := response.(map[string]interface{})["statusCode"]
		return c.Status(statusCode.(int)).JSON(response)
	}
	var post = models.Post{
		Title:       "test",
		Body:        "test",
		Description: "test",
		Slug:        "test",
		UserID:      uint(userID),
	}
	initalizer.DB.Create(&post)
	var response = helpers.SuccessMessage("Post Created")
	statusCode := response.(map[string]interface{})["statusCode"]
	return c.Status(statusCode.(int)).JSON(response)
}

type Params struct {
	Id int `json:"id" form:"id" validate:"required"`
}

func PostShow(c *fiber.Ctx) error {

	var data = Params{}
	err := c.BodyParser(&data)
	if err != nil {
		var response = helpers.ErrorMessage("Error occured!!")
		statusCode := response.(map[string]interface{})["statusCode"]
		return c.Status(statusCode.(int)).JSON(response)

	}
	// Validate the struct
	v := validator.New()
	if err := v.Struct(data); err != nil {
		// Handle validation errors
		for _, err := range err.(validator.ValidationErrors) {
			var validationMes string = "Field:" + err.Field() + ", Error: " + err.Tag()

			var response = helpers.ValidationMessage(validationMes)
			statusCode := response.(map[string]interface{})["statusCode"]
			return c.Status(statusCode.(int)).JSON(response)
		}
	}
	var post models.Post
	initalizer.DB.Preload("User").First(&post, data.Id)
	var response = helpers.GetData(post, "")
	statusCode := response.(map[string]interface{})["statusCode"]
	return c.Status(statusCode.(int)).JSON(response)
}

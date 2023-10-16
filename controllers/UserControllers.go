package controllers

import (
	"github.com/binesh/gomvc/helpers"

	"strconv"
	"time"

	"github.com/binesh/gomvc/initalizer"
	"github.com/binesh/gomvc/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string
	err := c.BodyParser(&data)
	if err != nil {
		var response = helpers.ErrorMessage("Error occured!!")
		statusCode := response.(map[string]interface{})["statusCode"]
		return c.Status(statusCode.(int)).JSON(response)

	}
	password, err := bcrypt.GenerateFromPassword([]byte(data["password"]), 14) //GenerateFromPassword returns the bcrypt hash of the password at the given cost i.e. (14 in our case).
	if err != nil {
		var response = helpers.ErrorMessage("Error occured!!")
		statusCode := response.(map[string]interface{})["statusCode"]
		return c.Status(statusCode.(int)).JSON(response)

	}
	user := models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
	}

	initalizer.DB.Create(&user) //Adds the data to the DB
	var response = helpers.GetData(user, "Registered Successfully!!")
	statusCode := response.(map[string]interface{})["statusCode"]
	return c.Status(statusCode.(int)).JSON(response)

}

const SecretKey = "secret"

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {

		var response = helpers.ErrorMessage("Invalid Request!!")
		statusCode := response.(map[string]interface{})["statusCode"]
		return c.Status(statusCode.(int)).JSON(response)
	}

	var user models.User

	initalizer.DB.Where("email = ?", data["email"]).First(&user) //Check the email is present in the DB

	if user.ID == 0 { //If the ID return is '0' then there is no such email present in the DB
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	} // If the email is present in the DB then compare the Passwords and if incorrect password then return error.

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.ID)),            //issuer contains the ID of the user.
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), //Adds time to the token i.e. 24 hours.
	})

	token, err := claims.SignedString([]byte(SecretKey))
	if err != nil {

		var response = helpers.ErrorMessage("Error occured!!")
		statusCode := response.(map[string]interface{})["statusCode"]
		return c.Status(statusCode.(int)).JSON(response)
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	} //Creates the cookie to be passed.

	c.Cookie(&cookie)

	c.Status(200)

	var response = helpers.GetData(user, "Registered Successfully!!")
	statusCode := response.(map[string]interface{})["statusCode"]
	return c.Status(statusCode.(int)).JSON(response) // If Login is Successfully done return the User data.

}

func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil //using the SecretKey which was generated in th Login function
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User
	initalizer.DB.Where("id = ?", claims.Issuer).First(&user)
	var response = helpers.GetData(user, "User Data found!!")
	statusCode := response.(map[string]interface{})["statusCode"]
	return c.Status(statusCode.(int)).JSON(response)

}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour), //Sets the expiry time an hour ago in the past.
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	var response = helpers.SuccessMessage("User Logout Successfully!!")
	statusCode := response.(map[string]interface{})["statusCode"]
	return c.Status(statusCode.(int)).JSON(response)

}

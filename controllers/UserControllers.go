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

	"github.com/go-playground/validator/v10"
)

type RegisterParams struct {
	Name     string `json:"name" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required" "email"`
	Password string `json:"password" form:"password" validate:"required"`
}

func Register(c *fiber.Ctx) error {
	var data = RegisterParams{}
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
	password, err := bcrypt.GenerateFromPassword([]byte(data.Password), 14) //GenerateFromPassword returns the bcrypt hash of the password at the given cost i.e. (14 in our case).
	if err != nil {
		var response = helpers.ErrorMessage("Error occured!!")
		statusCode := response.(map[string]interface{})["statusCode"]
		return c.Status(statusCode.(int)).JSON(response)

	}
	user := models.User{
		Name:     data.Name,
		Email:    data.Email,
		Password: password,
	}

	initalizer.DB.Create(&user) //Adds the data to the DB
	var response = helpers.GetData(user, "Registered Successfully!!")
	statusCode := response.(map[string]interface{})["statusCode"]
	return c.Status(statusCode.(int)).JSON(response)

}

const SecretKey = "secret"

type LoginParams struct {
	Email    string `json:"email" form:"email" validate:"required" "email"`
	Password string `json:"password" form:"password" validate:"required"`
}

func Login(c *fiber.Ctx) error {
	var data = LoginParams{}
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
	var user models.User
	initalizer.DB.Where("email = ?", data.Email).First(&user) //Check the email is present in the DB

	if user.ID == 0 { //If the ID return is '0' then there is no such email present in the DB

		var response = helpers.ErrorMessage("User not found!!")
		statusCode := response.(map[string]interface{})["statusCode"]
		return c.Status(statusCode.(int)).JSON(response)
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data.Password)); err != nil {

		var response = helpers.ErrorMessage("Incorrect Password!!")
		statusCode := response.(map[string]interface{})["statusCode"]
		return c.Status(statusCode.(int)).JSON(response)
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
	var returns = make(map[string]interface{})
	returns["userdetails"] = user
	returns["token"] = token
	var response = helpers.GetData(returns, "Login Successfully!!")
	statusCode := response.(map[string]interface{})["statusCode"]
	return c.Status(statusCode.(int)).JSON(response) // If Login is Successfully done return the User data.

}

func User(c *fiber.Ctx) error {
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

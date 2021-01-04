package auth

import (
	"fmt"
	"log"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/gofiber/fiber/v2"
)

// AccountModel - type for signup account
type AccountModel struct {
	Username             string `json:"username" xml:"username" form:"username"`
	Email                string `json:"email" xml:"email" form:"email"`
	Password             string `json:"password" xml:"password" form:"password"`
	PasswordConfirmation string `json:"passwordConfirmation" xml:"passwordConfirmation" form:"passwordConfirmation"`
}

func (am AccountModel) Validate() error {
	return validation.ValidateStruct(&am,
		validation.Field(&am.Username, validation.Required),
		validation.Field(&am.Email, validation.Required),
		validation.Field(&am.Password, validation.Required),
		validation.Field(&am.PasswordConfirmation, validation.Required),
	)
}

// JSONResponse - type for signup reponse
type JSONResponse struct {
	StatusCode    int    `json:"statusCode"`
	StatusMessage string `json:"statusMessage"`
}

// SignUp - create a new user
func SignUp(c *fiber.Ctx) error {
	c.SendString("signup")
	am := new(AccountModel)

	if err := c.BodyParser(am); err != nil {
		return err
	}

	log.Println(am.Username)
	log.Println(am.Email)
	log.Println(am.Password)
	log.Println(am.PasswordConfirmation)

	err := am.Validate()
	fmt.Println(err)
	response := JSONResponse{StatusCode: 200, StatusMessage: "create successfully"}

	c.JSON(response.StatusMessage)
	c.SendStatus(response.StatusCode)

	return nil
}

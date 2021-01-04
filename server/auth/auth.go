package auth

import (
	"errors"
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

func passwordMustMatch(str *string) validation.RuleFunc {
	return func(value interface{}) error {
		s, _ := value.(*string)
		if s != str {
			return errors.New("password must match")
		}
		return nil
	}
}

// Validate - validate for signup
func (am AccountModel) Validate() error {
	return validation.ValidateStruct(&am,
		validation.Field(&am.Username, validation.Required),
		validation.Field(&am.Email, validation.Required),
		validation.Field(&am.Password, validation.Required),
		validation.Field(&am.PasswordConfirmation, validation.Required,
			validation.By(passwordMustMatch(&am.Password))),
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
	if err != nil {
		response := JSONResponse{StatusCode: 400, StatusMessage: "account not created"}
		c.JSON(response.StatusMessage)
		c.SendStatus(response.StatusCode)
		return nil
	}
	response := JSONResponse{StatusCode: 200, StatusMessage: "create successfully"}

	c.JSON(response.StatusMessage)
	c.SendStatus(response.StatusCode)

	return nil
}

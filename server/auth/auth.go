package auth

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

// AccountModel - type for signup account
type AccountModel struct {
	Username             string `json:"username" xml:"username" form:"username"`
	Email                string `json:"email" xml:"email" form:"email"`
	Password             string `json:"password" xml:"password" form:"password"`
	PasswordConfirmation string `json:"passwordConfirmation" xml:"passwordConfirmation" form:"passwordConfirmation"`
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

	response := JSONResponse{StatusCode: 200, StatusMessage: "create successfully"}

	c.JSON(response.StatusMessage)
	c.SendStatus(response.StatusCode)

	return nil
}

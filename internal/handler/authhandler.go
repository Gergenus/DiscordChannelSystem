package handler

import (
	"log"

	"github.com/Gergenus/internal/models"
	"github.com/Gergenus/internal/service"
	"github.com/labstack/echo/v4"
)

type Messenger struct {
	Message interface{}
}

type authHandler interface {
	SignIn(ctx echo.Context) error
	SignUp(ctx echo.Context) error
}

type EchoAuthHandler struct {
	authService service.Auth
}

func NewEchoAuthHandler(authService service.Auth) EchoAuthHandler {
	return EchoAuthHandler{authService: authService}
}

func (e *EchoAuthHandler) SignIn(ctx echo.Context) error {
	var input models.AuthInput
	err := ctx.Bind(&input)
	if err != nil {
		log.Println("SignUp", err)
		return echo.NewHTTPError(400, "Bad req")
	}
	token, err := e.authService.SignIn(input.Name, input.Password)
	if err != nil {
		log.Println("SignIn", err)
	}
	return ctx.JSON(200, Messenger{
		Message: token,
	})
}

func (e *EchoAuthHandler) SignUp(ctx echo.Context) error {
	var input models.AuthInput
	err := ctx.Bind(&input)
	if err != nil {
		log.Println("SignUp", err)
		return echo.NewHTTPError(400, "Bad req")
	}

	id, err := e.authService.SignUp(input.Name, input.Password)
	if err != nil {
		log.Println("SignUp", err)
		return echo.NewHTTPError(500, "Internal error")
	}
	return ctx.JSON(200, Messenger{Message: id})
}

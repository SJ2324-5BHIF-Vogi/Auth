package controller

import (
	"bytes"
	"context"
	"crypto/sha256"

	"github.com/SJ2324-5BHIF-Vogi/Auth/pkg/dto"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type UserReader interface {
	Read(ctx context.Context, id uuid.UUID) (*dto.UserDTO, error)

	ReadByName(ctx context.Context, username string) (*dto.UserDTO, error)
}

type UserPasswordReader interface {
	Read(ctx context.Context, id uuid.UUID) ([]byte, error)

	ReadByName(ctx context.Context, username string) ([]byte, error)
}

type AuthController struct {
	reader UserPasswordReader

	userReader UserReader
}

func NewAuthController(reader UserPasswordReader, userReader UserReader) *AuthController {
	return &AuthController{reader: reader, userReader: userReader}
}

// Authenticate handles the authentication process for a user.
// It takes an echo.Context as input and returns an error.
// The function retrieves the username and password from the request form values,
// hashes the password using SHA256, and compares it with the stored password.
// If the passwords match, it generates a JWT token with the user ID and returns it as a string.
// If there is an error during the authentication process, it logs the error and returns it.
// If the username or password is incorrect, it returns the corresponding HTTP status code.
func (ac *AuthController) Authenticate(e echo.Context) error {
	ctx := e.Request().Context()

	username := e.FormValue("username")
	password := e.FormValue("password")

	hahs := sha256.New()
	hahs.Write([]byte(password))
	hashed := hahs.Sum(nil)

	pass, err := ac.reader.ReadByName(ctx, username)

	if err != nil {
		e.Logger().Error(err)
		return err
	}

	if pass == nil {
		return e.NoContent(404)
	}

	if !bytes.Equal(pass, hashed) {
		return e.NoContent(401)
	}

	usr, err := ac.userReader.ReadByName(ctx, username)

	if err != nil {
		e.Logger().Error(err)
		return err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user": usr.ID,
		})

	tokenString, err := token.SignedString([]byte("SECRET")) // TODO: move secret to config

	if err != nil {
		e.Logger().Error(err)
		return err
	}

	return e.JSON(200, echo.Map{
		"token": tokenString,
	})
}

// Validate validates the token received from the client.
// It parses the token using the provided secret key and checks if it is valid.
// If the token is valid, it returns a 200 status code.
// If the token is invalid, it returns a 401 status code.
func (ac *AuthController) Validate(e echo.Context) error {
	tokenString := e.FormValue("token")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("SECRET"), nil // TODO: move secret to config
	})

	if err != nil {
		e.Logger().Error(err)
		return err
	}

	if !token.Valid {
		return e.NoContent(401)
	}

	return e.NoContent(200)
}

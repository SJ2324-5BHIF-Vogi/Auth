package controller

import (
	"context"

	"github.com/SJ2324-5BHIF-Vogi/Auth/pkg/dto"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type userCreator interface {
	Create(ctx context.Context, username string, password []byte) error
}

type userDeleter interface {
	Delete(ctx context.Context, id uuid.UUID) error
}

type userReader interface {
	Read(ctx context.Context, id uuid.UUID) (*dto.UserDTO, error)

	ReadByName(ctx context.Context, username string) (*dto.UserDTO, error)
}

type userUpdater interface {
	Update(ctx context.Context, id uuid.UUID, username string, password []byte) error
}

type UserController struct {
	creator userCreator

	deleter userDeleter

	reader userReader

	updater userUpdater
}

func NewUserController(creator userCreator, deleter userDeleter, reader userReader, updater userUpdater) *UserController {
	return &UserController{creator: creator, deleter: deleter, reader: reader, updater: updater}
}

// GetUser retrieves a user by their ID.
// It takes an echo.Context as a parameter and returns an error.
// The function first parses the ID parameter from the request URL.
// If the ID is not a valid UUID, it logs an error and returns the error.
// Otherwise, it calls the Read method of the UserController's reader field to retrieve the user.
// If an error occurs during the retrieval, it logs the error and returns it.
// Finally, it returns the user as a JSON response with status code 200.
func (uc *UserController) GetUser(c echo.Context) error {
	ctx := c.Request().Context()

	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.Logger().Error(err)
		return err
	}

	usr, err := uc.reader.Read(ctx, id)

	if err != nil {
		c.Logger().Error(err)
		return err
	}

	return c.JSON(200, usr)
}

// CreateUser creates a new user.
// It takes an echo.Context as a parameter and returns an error.
// The function retrieves the username and password from the request form values.
// It then calls the Create method of the UserController's creator field to create the user.
// If an error occurs during the creation process, it logs the error using the echo.Context's logger.
// After that, it returns a 201 status code with no content.
func (uc *UserController) CreateUser(c echo.Context) error {
	ctx := c.Request().Context()

	username := c.FormValue("username")
	password := c.FormValue("password")

	if err := uc.creator.Create(ctx, username, []byte(password)); err != nil {
		c.Logger().Error(err)
	}

	// TODO: push RabbitMQ user profile creation message to queue

	return c.NoContent(201)
}

// DeleteUser deletes a user with the specified ID.
// It takes an echo.Context as a parameter and returns an error.
// The function first parses the ID parameter from the context and then calls the Delete method of the UserController's deleter field to delete the user.
// If any error occurs during the parsing or deletion process, it logs the error and returns it.
// Otherwise, it returns a 204 No Content response.
func (uc *UserController) DeleteUser(c echo.Context) error {
	ctx := c.Request().Context()

	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.Logger().Error(err)
		return err
	}

	if err := uc.deleter.Delete(ctx, id); err != nil {
		c.Logger().Error(err)
		return err
	}

	return c.NoContent(204)
}

// UpdateUser updates a user's information.
// It takes an echo.Context as a parameter and returns an error.
// The function parses the user ID from the request parameter, retrieves the username and password from the form values,
// hashes the password, and then calls the updater.Update method to update the user's information in the database.
// If any error occurs during the process, it logs the error and returns it.
// Otherwise, it returns a 204 No Content response.
func (uc *UserController) UpdateUser(c echo.Context) error {
	ctx := c.Request().Context()

	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.Logger().Error(err)
		return err
	}

	username := c.FormValue("username")
	password := c.FormValue("password")

	// TODO: hash password

	if err := uc.updater.Update(ctx, id, username, []byte(password)); err != nil {
		c.Logger().Error(err)
		return err
	}

	return c.NoContent(204)
}

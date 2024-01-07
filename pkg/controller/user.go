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

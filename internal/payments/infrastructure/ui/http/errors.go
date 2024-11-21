package http

import (
	"github.com/labstack/echo/v4"
)

type HTTPError struct {
	Message string `json:"message"`
}

func (e HTTPError) InternalServerError() error {
	err := echo.ErrInternalServerError
	err.Message = e.Message

	return err
}

func (e HTTPError) BadRequest() error {
	err := echo.ErrBadRequest
	err.Message = e.Message

	return err
}

func (e HTTPError) NotFound() error {
	err := echo.ErrNotFound
	err.Message = e.Message

	return err
}

func (e HTTPError) Unauthorized() error {
	err := echo.ErrUnauthorized
	err.Message = e.Message

	return err
}

func (e HTTPError) Forbidden() error {
	err := echo.ErrForbidden
	err.Message = e.Message

	return err
}

func (e HTTPError) Conflict() error {
	err := echo.ErrConflict
	err.Message = e.Message

	return err
}

func (e HTTPError) ErrUnprocessableEntity() error {
	err := echo.ErrUnprocessableEntity
	err.Message = e.Message

	return err
}

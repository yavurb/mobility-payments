package middlewares

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/yavurb/mobility-payments/internal/auth/domain"
	usersDomain "github.com/yavurb/mobility-payments/internal/users/domain"
)

func Authenticate(tokenManager domain.TokenManager, headerKey string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := c.Request().Header.Get(headerKey)
			if token == "" {
				err := echo.ErrUnauthorized
				err.Message = "token not found"

				return err
			}

			payload, err := tokenManager.Verify(token)
			if err != nil {
				fmt.Printf("%v", err)
				err := echo.ErrUnauthorized
				err.Message = "invalid token"

				return err
			}

			c.Set("user", payload)

			return next(c)
		}
	}
}

func Authorize(scopes []usersDomain.UserType) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			payload, ok := c.Get("user").(*domain.TokenPayload)
			if !ok {
				err := echo.ErrUnauthorized
				err.Message = "user not authenticated"

				return err
			}

			for _, scope := range scopes {
				if payload.Type == scope {
					return next(c)
				}
			}

			err := echo.ErrForbidden
			err.Message = "user not authorized"

			return err
		}
	}
}

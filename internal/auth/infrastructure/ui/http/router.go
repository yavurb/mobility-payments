package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yavurb/mobility-payments/internal/auth/domain"
	usersDomain "github.com/yavurb/mobility-payments/internal/users/domain"
)

type authRouter struct {
	echo        *echo.Echo
	authUsecase domain.Usecase
}

func NewAuthRouter(e *echo.Echo, authUsecase domain.Usecase) {
	routerGroup := e.Group("/auth")
	routerCtx := authRouter{
		echo:        e,
		authUsecase: authUsecase,
	}

	routerGroup.POST("/signup", routerCtx.signup)
	routerGroup.POST("/signin", routerCtx.signin)
}

func (r *authRouter) signup(c echo.Context) error {
	var data SignUpData

	if err := c.Bind(&data); err != nil {
		return HTTPError{
			Message: "Invalid request data",
		}.BadRequest()
	}

	if err := c.Validate(data); err != nil {
		// TODO: Return the list of errors
		return HTTPError{
			Message: "Invalid payload",
		}.ErrUnprocessableEntity()
	}

	token, id, err := r.authUsecase.SignUp(c.Request().Context(), usersDomain.UserType(data.AccountType), data.Name, data.Email, data.Password)
	if err != nil {
		return handleErr(err)
	}

	return c.JSON(http.StatusCreated, AuthPayload{Token: token, ID: id})
}

func (r *authRouter) signin(c echo.Context) error {
	var data SignInData

	if err := c.Bind(&data); err != nil {
		return HTTPError{
			Message: "Invalid request data",
		}.BadRequest()
	}

	if err := c.Validate(data); err != nil {
		// TODO: Return the list of errors
		return HTTPError{
			Message: "Invalid payload",
		}.ErrUnprocessableEntity()
	}

	token, id, err := r.authUsecase.SignIn(c.Request().Context(), data.Email, data.Password)
	if err != nil {
		return handleErr(err)
	}

	return c.JSON(http.StatusOK, AuthPayload{Token: token, ID: id})
}

func handleErr(err error) error {
	switch err {
	case domain.ErrUserNotFound:
		return HTTPError{
			Message: "User not found",
		}.NotFound()
	case domain.ErrInvalidCredentials:
		return HTTPError{
			Message: "Invalid credentials",
		}.Unauthorized()
	default:
		return HTTPError{
			Message: "Internal server error",
		}.InternalServerError()
	}
}

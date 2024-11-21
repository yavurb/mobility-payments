package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	authDomain "github.com/yavurb/mobility-payments/internal/auth/domain"
	"github.com/yavurb/mobility-payments/internal/common/middlewares"
	"github.com/yavurb/mobility-payments/internal/payments/domain"
	usersDomain "github.com/yavurb/mobility-payments/internal/users/domain"
)

type authRouter struct {
	echo            *echo.Echo
	paymentsUsecase domain.Usecase
}

func NewPaymentsRouter(e *echo.Echo, paymentsUsecase domain.Usecase, tokenManager authDomain.TokenManager, headerKey string) {
	routerGroup := e.Group("/payments", middlewares.Authenticate(tokenManager, headerKey))
	routerCtx := authRouter{
		echo:            e,
		paymentsUsecase: paymentsUsecase,
	}

	routerGroup.GET("", routerCtx.getTransactions, middlewares.Authorize([]usersDomain.UserType{usersDomain.Customer, usersDomain.Merchant}))
	routerGroup.GET("/:id", routerCtx.getTransaction, middlewares.Authorize([]usersDomain.UserType{usersDomain.Customer, usersDomain.Merchant}))
	routerGroup.PATCH("/:id/verify", routerCtx.verify, middlewares.Authorize([]usersDomain.UserType{usersDomain.Merchant}))
	routerGroup.POST("/pay", routerCtx.pay, middlewares.Authorize([]usersDomain.UserType{usersDomain.Customer}))
}

func (r *authRouter) pay(c echo.Context) error {
	var data PaymentData

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

	payload, _ := c.Get("user").(*authDomain.TokenPayload)

	transaction, err := r.paymentsUsecase.Pay(c.Request().Context(), payload.ID, data.Merchant, data.Description, data.Method, data.Amount)
	if err != nil {
		return handleErr(err)
	}

	return c.JSON(http.StatusCreated, TransactionCreated{ID: transaction.PublicID, Status: transaction.Status})
}

func (r *authRouter) getTransaction(c echo.Context) error {
	var data GetTransactionParam

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

	payload, _ := c.Get("user").(*authDomain.TokenPayload)

	transaction, err := r.paymentsUsecase.GetTransaction(c.Request().Context(), data.ID, payload.ID)
	if err != nil {
		return handleErr(err)
	}

	return c.JSON(http.StatusOK, Transaction{
		ID:          transaction.PublicID,
		ReceiverID:  transaction.ReceiverPublicID,
		SenderID:    transaction.SenderPublicID,
		Status:      transaction.Status,
		Description: transaction.Description,
		Amount:      transaction.Amount,
		Method:      transaction.Method,
		CreatedAt:   transaction.CreatedAt,
	})
}

func (r *authRouter) getTransactions(c echo.Context) error {
	payload, _ := c.Get("user").(*authDomain.TokenPayload)

	transactions, err := r.paymentsUsecase.GetTransactions(c.Request().Context(), payload.ID)
	if err != nil {
		return handleErr(err)
	}

	var transactionsList []Transaction
	for _, transaction := range transactions {
		transactionsList = append(transactionsList, Transaction{
			ID:          transaction.PublicID,
			ReceiverID:  transaction.ReceiverPublicID,
			SenderID:    transaction.SenderPublicID,
			Status:      transaction.Status,
			Description: transaction.Description,
			Amount:      transaction.Amount,
			Method:      transaction.Method,
			CreatedAt:   transaction.CreatedAt,
		})
	}

	return c.JSON(http.StatusOK, TransactionList{Data: transactionsList})
}

func (r *authRouter) verify(c echo.Context) error {
	var data VerifyTransactionData

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

	status := domain.Declined
	if data.Confirmation == ConfirmConfirmation {
		status = domain.Succeeded
	}

	payload, _ := c.Get("user").(*authDomain.TokenPayload)

	transactions, err := r.paymentsUsecase.Verify(c.Request().Context(), data.ID, status, payload.ID)
	if err != nil {
		return handleErr(err)
	}

	return c.JSON(http.StatusOK, TransactionVeriyfied{ID: transactions.PublicID, Status: transactions.Status})
}

func handleErr(err error) error {
	switch err {
	case domain.ErrTransactionNotFound:
		return HTTPError{
			Message: "Transaction not found",
		}.NotFound()
	case domain.ErrInsufficientBalance:
		return HTTPError{
			Message: "User has insufficient balance",
		}.ErrUnprocessableEntity()
	case domain.ErrTransactionNotFromUser:
		return HTTPError{
			Message: "User is not allowed to access this transaction",
		}.Forbidden()
	default:
		return HTTPError{
			Message: "Internal server error",
		}.InternalServerError()
	}
}

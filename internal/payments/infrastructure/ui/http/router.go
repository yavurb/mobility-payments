package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yavurb/mobility-payments/internal/payments/domain"
)

type authRouter struct {
	echo            *echo.Echo
	paymentsUsecase domain.Usecase
}

func NewPaymentsRouter(e *echo.Echo, paymentsUsecase domain.Usecase) {
	routerGroup := e.Group("/payments")
	routerCtx := authRouter{
		echo:            e,
		paymentsUsecase: paymentsUsecase,
	}

	routerGroup.GET("/", routerCtx.getTransactions)
	routerGroup.GET("/:id", routerCtx.getTransaction)
	routerGroup.POST("/:id/verify", routerCtx.verify)
	routerGroup.POST("/pay", routerCtx.pay)
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

	transaction, err := r.paymentsUsecase.Pay(c.Request().Context(), "", data.Merchant, data.Description, data.Method, data.Amount)
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

	transaction, err := r.paymentsUsecase.GetTransaction(c.Request().Context(), data.ID, "")
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
	transactions, err := r.paymentsUsecase.GetTransactions(c.Request().Context(), "")
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

	transactions, err := r.paymentsUsecase.Verify(c.Request().Context(), data.ID, status, "")
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

package http

type SignUpData struct {
	Name        string `json:"name" validate:"required"`
	Email       string `json:"email" validate:"required"`
	AccountType string `json:"account_type" validate:"required,oneof=customer merchant"`
	Password    string `json:"password" validate:"required"`
}

type SignInData struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type AuthPayload struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}

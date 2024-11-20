package adapters

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/yavurb/mobility-payments/internal/auth/domain"
	userDomain "github.com/yavurb/mobility-payments/internal/users/domain"
)

type authTokenManager struct {
	JWTSecret string
}

func NewAuthTokenManager(secret string) domain.TokenManager {
	return &authTokenManager{
		JWTSecret: secret,
	}
}

func (tm *authTokenManager) Generate(payload domain.TokenPayload, duration time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  payload.ID,
		"type": payload.Type,
		"exp":  time.Now().UTC().Add(duration).Unix(),
	})

	signedToken, err := token.SignedString([]byte(tm.JWTSecret))
	if err != nil {
		fmt.Printf("Error signing token: %v\n", err)
		return "", err
	}

	return signedToken, nil
}

func (tm *authTokenManager) Verify(tokenString string) (*domain.TokenPayload, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return tm.JWTSecret, nil
	})
	if err != nil {
		fmt.Printf("Error parsing token: %v\n", err)
		return nil, err
	}

	// Extract and validate claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Printf("Token is valid. Claims: %v\n", claims)

		id, err := claims.GetSubject()
		if err != nil {
			return nil, err
		}

		user_type := claims["type"].(userDomain.UserType)

		return &domain.TokenPayload{
			ID:   id,
			Type: user_type,
		}, nil
	} else {
		return nil, domain.ErrInvalidCredentials
	}
}

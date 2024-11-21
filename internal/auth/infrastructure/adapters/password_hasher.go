package adapters

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"github.com/yavurb/mobility-payments/internal/auth/domain"
	"golang.org/x/crypto/argon2"
)

type authPasswordHasher struct{}

func NewAuthPasswordHasher() domain.PasswordHasher {
	return &authPasswordHasher{}
}

const (
	memory      = 64 * 1024
	iterations  = 3
	parallelism = 2
	saltLength  = 16
	keyLength   = 32
)

func (ph *authPasswordHasher) Hash(password string) (string, error) {
	salt := make([]byte, saltLength)

	_, err := rand.Read(salt)
	if err != nil {
		return "", fmt.Errorf("failed to generate salt: %w", err)
	}

	hashedPassword := argon2.IDKey([]byte(password), salt, iterations, memory, parallelism, keyLength)

	saltBase64 := base64.RawStdEncoding.EncodeToString(salt)
	hashBase64 := base64.RawStdEncoding.EncodeToString(hashedPassword)

	return fmt.Sprintf("%s:%s", saltBase64, hashBase64), nil
}

func (ph *authPasswordHasher) Verify(password, hashedPassword string) (bool, error) {
	parts := strings.Split(hashedPassword, ":")
	if len(parts) != 2 {
		return false, errors.New("invalid hashed value format")
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[0])
	if err != nil {
		return false, fmt.Errorf("failed to decode salt: %w", err)
	}

	hash, err := base64.RawStdEncoding.DecodeString(parts[1])
	if err != nil {
		return false, fmt.Errorf("failed to decode hash: %w", err)
	}

	computedHash := argon2.IDKey([]byte(password), salt, iterations, memory, parallelism, keyLength)

	return string(hash) == string(computedHash), nil
}

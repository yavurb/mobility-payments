package ids

import (
	"strings"

	nanoid "github.com/matoous/go-nanoid/v2"
)

const (
	alphabet = "0123456789abcdefghijklmnopqrstuvwxyz"
	length   = 12
)

func NewPublicID(prefix string) (string, error) {
	id, err := nanoid.Generate(alphabet, length)
	if err != nil {
		return "", err
	}

	return strings.Join([]string{prefix, id}, "_"), nil
}

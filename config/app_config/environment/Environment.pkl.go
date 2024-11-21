// Code generated from Pkl module `Config`. DO NOT EDIT.
package environment

import (
	"encoding"
	"fmt"
)

type Environment string

const (
	Development Environment = "development"
	Production  Environment = "production"
)

// String returns the string representation of Environment
func (rcv Environment) String() string {
	return string(rcv)
}

var _ encoding.BinaryUnmarshaler = new(Environment)

// UnmarshalBinary implements encoding.BinaryUnmarshaler for Environment.
func (rcv *Environment) UnmarshalBinary(data []byte) error {
	switch str := string(data); str {
	case "development":
		*rcv = Development
	case "production":
		*rcv = Production
	default:
		return fmt.Errorf(`illegal: "%s" is not a valid Environment`, str)
	}
	return nil
}

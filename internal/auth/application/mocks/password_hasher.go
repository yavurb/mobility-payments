package mocks

type PasswordHasherMock struct {
	HashFn   func(password string) (string, error)
	VerifyFn func(password, hashedPassword string) bool
}

func (ph *PasswordHasherMock) Hash(password string) (string, error) {
	return ph.HashFn(password)
}

func (ph *PasswordHasherMock) Verify(password, hashedPassword string) bool {
	return ph.VerifyFn(password, hashedPassword)
}

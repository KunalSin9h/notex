package user

import (
	"github.com/google/uuid"
	"github.com/kunalsin9h/notex/password"
)

type User struct {
	ID           uuid.UUID    `json:"id"`
	Username     Username     `json:"username"`
	Email        Email        `json:"email"`
	PasswordHash PasswordHash `json:"passwordHash"`
}

type Username string
type Email string
type PasswordHash string

func ParseUsername(row string) (Username, error) {
	return Username(row), nil
}

func ParseEmail(row string) (Email, error) {
	return Email(row), nil
}

func ParsePassword(row string) (PasswordHash, error) {
	return PasswordHash(row), nil
}

func (u *User) HashPassword(row string) error {
	argon2Param := password.Params{
		Memory:      64 * 1024,
		Iterations:  3,
		Parallelism: 2,
		SaltLength:  16,
		KeyLength:   32,
	}

	hash, err := password.GenerateFromPassword(row, &argon2Param)
	if err != nil {
		return err
	}

	u.PasswordHash = PasswordHash(hash)
	return nil
}

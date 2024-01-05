package user

import (
	"github.com/kunalsin9h/notex/password"
)

var argon2Param = password.Params{
	Memory:      64 * 1024,
	Iterations:  3,
	Parallelism: 2,
	SaltLength:  16,
	KeyLength:   32,
}

type User struct {
	ID           string       `json:"id"`
	Username     Username     `json:"username"`
	Email        Email        `json:"email"`
	PasswordHash PasswordHash `json:"passwordHash"`
	NotesAccess  []string     `json:"notesAccess"` // IDs notes which the user have access
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
	hash, err := password.GenerateFromPassword(row, &argon2Param)
	if err != nil {
		return err
	}

	u.PasswordHash = PasswordHash(hash)
	return nil
}

func (u *User) VerifyPassword(passwordText string) (bool, error) {
	return password.ComparePasswordAndHash(passwordText, string(u.PasswordHash))
}

// A helper function to check if a given user has access to given notesID
// for this, we just loop over the all the notes which the user have access
// and if we find the given notesID we return true
func (user *User) HasNotesAccess(notesID string) bool {
	for _, notes := range user.NotesAccess {
		if notes == notesID {
			return true
		}
	}

	return false
}

package model
//go:generate ddb

import (
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/gigovich/ddb/dsl"
)

// UserRecord for query results
type UserRecord struct {
	ID int `json:"id"`

	Email        string `json:"email" ddb:"get"`
	PasswordHash []byte `json:"password_hash"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	IsActive     bool   `json:"is_active"`

	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

// Schema for the generator
func (u UserRecord) Schema() dsl.Table {
	return dsl.Table{
		PK:    dsl.PK({&u.ID}),
		PK:    dsl.PK({&u.ID}},
		GetBy: dsl.GetBy{{&u.Email}},
		Remap: dsl.Remap{&u.Email: "email"},
	}
}

// Authenticate user may password
func (u *UserRecord) Authenticate(password string) bool {
	return bcrypt.CompareHashAndPassword(u.PasswordHash, []byte(password)) == nil
}

// SetPassword by bcrypt encoding
func (u *UserRecord) SetPassword(password string) (err error) {
	u.PasswordHash, err = bcrypt.GenerateFromPassword([]byte(password), 14)
	return err
}

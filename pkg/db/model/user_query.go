package model

import (
	"context"

	"github.com/jackc/pgx"
)

// UserQuery to query user records
type UserQuery struct {
	ctx context.Context
	tx  *pgx.Tx
}

// Create user record and save it in th database
func (uq *UserQuery) Create(ur *UserRecord, password string) (err error) {
	_, err = uq.tx.ExecEx(
		uq.ctx,
		`INSERT INTO "user"(email, password_hash, first_name, last_name, is_active) VALUES ($1, $2, $3, $4, $5)`,
		&pgx.QueryExOptions{},
		ur.Email, ur.PasswordHash, ur.FirstName, ur.LastName, ur.IsActive,
	)
	if err != nil {
		return err
	}

	return err
}

// Get user by it's email
func (uq *UserQuery) Get(email string) (*UserRecord, error) {
	stmt := uq.tx.QueryRowEx(
		uq.ctx,
		`SELECT id, email, password_hash, first_name, last_name, is_active, updated_at, create_at `+
			`FROM "user" WHERE email = $1 LIMIT 1`,
		&pgx.QueryExOptions{},
		email,
	)

	ur := UserRecord{}
	err := stmt.Scan(
		&ur.ID, &ur.Email, &ur.PasswordHash, &ur.FirstName, &ur.LastName, &ur.IsActive, &ur.UpdatedAt, &ur.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &ur, nil
}

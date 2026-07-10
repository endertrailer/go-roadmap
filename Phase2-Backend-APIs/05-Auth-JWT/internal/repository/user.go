package repository

import (
	"context"
)

type User struct {
	Hashedpassword string `json:"password"`
	Username       string `json:"username"`
	ID             string `json:"id"`
}

func (r *PostgresRepo) CreateUser(ctx context.Context, username, passwordHash string) error {
	query := `INSERT INTO users (username, password_hash) VALUES ($1, $2)`
	_, err := r.db.Exec(ctx, query, username, passwordHash)
	return err
}

func (r *PostgresRepo) GetUserByUsername(ctx context.Context, username string) (*User, error) {
	user := User{}
	query := `SELECT id, username, password_hash from users where username = $1`
	err := r.db.QueryRow(ctx, query, username).Scan(&user.ID, &user.Username, &user.Hashedpassword)
	if err != nil {
		return nil, err
	}
	return &user, err
}




func (r *PostgresRepo) DeleteUser(ctx context.Context, id string) error {
	// TODO: Implement the SQL query to delete a user by ID (e.g., "DELETE FROM users WHERE id = $1") using r.db.Exec
	return ctx.Err()
}

package store

import (
	"src/myFirstGoProject/chapter003/http-rest-api/internal/app/model"
)

// UserRepo ...
type UserRepo struct {
	store *Store
}

// Create ...
func (r *UserRepo) Create(u *model.User) (*model.User, error) {
	if err := r.store.db.QueryRow(
		"INSERT INTO users (email, encrypted_password) VALUES ($1, $2) RETURNING id",
		u.Email,
		u.EncryptedPassword,
	).Scan(&u.ID); err != nil {
		return nil, err
	}
	return u, nil
}

// FindByEmail ...
func (r *UserRepo) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow(
		"SELECT id, email, encrypted_password FROM users WHERE email = $1",
		email,
	).Scan(
		&u.ID,
		&u.Email,
		&u.EncryptedPassword); err != nil {
		return nil, err
	}
	return u, nil
}

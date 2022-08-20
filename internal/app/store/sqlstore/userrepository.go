package sqlstore

import (
	"database/sql"

	"github.com/VladKuch/go_practic/internal/app/model"
	"github.com/VladKuch/go_practic/internal/app/store"
)

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) error {

	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	if err := r.store.db.QueryRow(
		"INSERT INTO users (first_name, last_name, email, password) VALUES ($1, $2, $3, $4) RETURNING user_id",
		u.FirstName,
		u.LastName,
		u.Email,
		u.PasswordEncrypted,
	).Scan(&u.ID); err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) Find(id int) (*model.User, error) {

	u := &model.User{}

	if err := r.store.db.QueryRow("SELECT user_id, first_name, last_name, email, password FROM users WHERE user_id = $1", id).Scan(
		&u.ID,
		&u.FirstName,
		&u.LastName,
		&u.Email,
		&u.PasswordEncrypted,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}

	return u, nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {

	u := &model.User{}

	if err := r.store.db.QueryRow("SELECT user_id, first_name, last_name, email, password FROM users WHERE email = $1", email).Scan(
		&u.ID,
		&u.FirstName,
		&u.LastName,
		&u.Email,
		&u.PasswordEncrypted,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}

	return u, nil
}

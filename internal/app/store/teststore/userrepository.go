package teststore

import (
	"github.com/root5427/http-rest-api/internal/app/model"
	"github.com/root5427/http-rest-api/internal/app/store"
)

type UserRepository struct {
	store *Store
	users map[string]*model.User
}

// Create ...
func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return  err
	}
	if err := u.BeforeCreate(); err != nil {
		return err
	}

	r.users[u.Email] = u
	u.ID = len(r.users)

	return nil
}

// FindByEmail ...
func (r *UserRepository) FindByEmail(email string) (u *model.User, err error) {
	u, ok := r.users[email]
	if !ok {
		return nil, store.RecordNotFound
	}

	return u, nil
}
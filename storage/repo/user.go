package repo

import (
	"context"
	"time"
)

type UserStorageI interface {
	Create(ctx context.Context, req *User) (*User, error)
	Get(ctx context.Context, id string) (*User, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, req *UpdateUser) error
}

type User struct {
	ID        string
	FirstName string
	LastName  string
	Email     string
	Password  string
	CreatedAt *time.Time `db:"created_at"`
}

type UpdateUser struct {
	ID        string
	FirstName string
	LastName  string
}

package user

import "context"

// Repository handle the CRUD operation with user
type Repository interface {
	// Here more information about the use of `context` https://blog.golang.org/context
	GetAll(ctx context.Context) ([]User, error)
	GetOne(ctx context.Context, id uint) (User, error)
	GetByUsername(ctx context.Context, username string) (User, error)
	Create(ctx context.Context, user *User) error
	Update(ctx context.Context, id uint, user *User) error
	Delete(ctx context.Context, id uint) error
}

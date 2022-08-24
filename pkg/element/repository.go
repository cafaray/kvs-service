package element

import "context"

type Repository interface {
	GetAll(ctx context.Context) ([]Element, error)
	GetOne(ctx context.Context, id uint) (Element, error)
	GetByUser(ctx context.Context, userID uint) ([]Element, error)
	Create(ctx context.Context, element *Element) error
	Update(ctx context.Context, id uint, element Element) error
	Delete(ctx context.Context, id uint) error
}

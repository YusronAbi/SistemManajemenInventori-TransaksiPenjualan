package repo

import "context"

type PropertyRepository interface {
	Create(ctx context.Context, property *Property) error
	GetByID(ctx context.Context, id int64) (*Property, error)
	Update(ctx context.Context, property *Property) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context, filter PropertyFilter) ([]Property, error)
}

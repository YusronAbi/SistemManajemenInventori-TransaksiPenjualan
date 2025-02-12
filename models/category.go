package models

type Category struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type CategoryRepository interface {
	GetByID(id int64) (*Category, error)
	Store(category *Category) error
	Update(category *Category) error
	Delete(id int64) error
	Fetch() ([]Category, error)
}

type CategoryUsecase interface {
	GetByID(id int64) (*Category, error)
	Store(category *Category) error
	Update(category *Category) error
	Delete(id int64) error
	Fetch() ([]Category, error)
}

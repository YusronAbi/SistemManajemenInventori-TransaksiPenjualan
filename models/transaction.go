package models

type Transaction struct {
	ID        int64 `json:"id"`
	UserID    int64 `json:"user_id"`
	ProductID int64 `json:"product_id"`
	Quantity  int   `json:"quantity"`
}

type TransactionRepository interface {
	GetByID(id int64) (*Transaction, error)
	Store(transaction *Transaction) error
	Update(transaction *Transaction) error
	Delete(id int64) error
	Fetch() ([]Transaction, error)
}

type TransactionUsecase interface {
	GetByID(id int64) (*Transaction, error)
	Store(transaction *Transaction) error
	Update(transaction *Transaction) error
	Delete(id int64) error
	Fetch() ([]Transaction, error)
}

package repository

import (
	"ECOMERCE-API/models"
	"database/sql"
	"time"
)

type ProductRepositoryImpl struct {
	DB *sql.DB
}

func NewProductRepository(db *sql.DB) models.ProductRepository {
	return &ProductRepositoryImpl{DB: db}
}

func (r *ProductRepositoryImpl) GetByID(id int64) (*models.Product, error) {
	var product models.Product
	err := r.DB.QueryRow("SELECT id, name, description, price, stock, category_id, created_at, updated_at FROM products WHERE id = ?", id).Scan(
		&product.ID, &product.Name, &product.Description, &product.Price, &product.Stock, &product.CategoryID, &product.CreatedAt, &product.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepositoryImpl) Store(product *models.Product) error {
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()
	_, err := r.DB.Exec("INSERT INTO products (name, description, price, stock, category_id, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)",
		product.Name, product.Description, product.Price, product.Stock, product.CategoryID, product.CreatedAt, product.UpdatedAt)
	return err
}

func (r *ProductRepositoryImpl) Update(product *models.Product) error {
	product.UpdatedAt = time.Now()
	_, err := r.DB.Exec("UPDATE products SET name = ?, description = ?, price = ?, stock = ?, category_id = ?, updated_at = ? WHERE id = ?",
		product.Name, product.Description, product.Price, product.Stock, product.CategoryID, product.UpdatedAt, product.ID)
	return err
}

func (r *ProductRepositoryImpl) Delete(id int64) error {
	_, err := r.DB.Exec("DELETE FROM products WHERE id = ?", id)
	return err
}

func (r *ProductRepositoryImpl) Fetch(cursor string, num int64) ([]models.Product, string, error) {
	return nil, "", nil
}

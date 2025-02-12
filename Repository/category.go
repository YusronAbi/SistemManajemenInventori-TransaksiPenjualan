package repository

import (
	"ECOMERCE-API/models"
	"database/sql"
)

type CategoryRepositoryImpl struct {
	DB *sql.DB
}

func NewCategoryRepository(db *sql.DB) models.CategoryRepository {
	return &CategoryRepositoryImpl{DB: db}
}

func (r *CategoryRepositoryImpl) GetByID(id int64) (*models.Category, error) {
	var category models.Category
	err := r.DB.QueryRow("SELECT id, name FROM categories WHERE id = ?", id).Scan(&category.ID, &category.Name)
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *CategoryRepositoryImpl) Store(category *models.Category) error {
	_, err := r.DB.Exec("INSERT INTO categories (name) VALUES (?)", category.Name)
	return err
}

func (r *CategoryRepositoryImpl) Update(category *models.Category) error {
	_, err := r.DB.Exec("UPDATE categories SET name = ? WHERE id = ?", category.Name, category.ID)
	return err
}

func (r *CategoryRepositoryImpl) Delete(id int64) error {
	_, err := r.DB.Exec("DELETE FROM categories WHERE id = ?", id)
	return err
}

func (r *CategoryRepositoryImpl) Fetch() ([]models.Category, error) {
	rows, err := r.DB.Query("SELECT id, name FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var category models.Category
		if err := rows.Scan(&category.ID, &category.Name); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}

package repository

import (
    "database/sql"
    "ecommerce-api/models"
)

type UserRepositoryImpl struct {
    DB *sql.DB
}

func NewUserRepository(db *sql.DB) models.UserRepository {
    return &User RepositoryImpl{DB: db}
}

// GetByID retrieves a user by  ID
func (r *User RepositoryImpl) GetByID(id int64) (*models.User, error) {
    var user models.User
    err := r.DB.QueryRow("SELECT id, username, password FROM users WHERE id = ?", id).Scan(
        &user.ID, &user.Username, &user.Password,
    )
    if err != nil {
        return nil, err
    }
    return &user, nil
}

func (r *User RepositoryImpl) Store(user *models.User) error {
    _, err := r.DB.Exec("INSERT INTO users (username, password) VALUES (?, ?)", user.Username, user.Password)
    return err
}

func (r *User RepositoryImpl) Update(user *models.User) error {
    _, err := r.DB.Exec("UPDATE users SET username = ?, password = ? WHERE id = ?", user.Username, user.Password, user.ID)
    return err
}

// Delete removes a user by its ID
func (r *User RepositoryImpl) Delete(id int64) error {
    _, err := r.DB.Exec("DELETE FROM users WHERE id = ?", id)
    return err
}

// Fetch retrieves all users
func (r *User RepositoryImpl) Fetch() ([]models.User, error) {
    rows, err := r.DB.Query("SELECT id, username, password FROM users")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var users []models.User
    for rows.Next() {
        var user models.User
        if err := rows.Scan(&user.ID, &user.Username, &user.Password); err != nil {
            return nil, err
        }
        users = append(users, user)
    }
    return users, nil
}
package repository

import (
    "database/sql"
    "ecommerce-api/models"
)

type TransactionRepositoryImpl struct {
    DB *sql.DB
}

func NewTransactionRepository 
(db *sql.DB) models.TransactionRepository {
    return &TransactionRepositoryImpl{DB: db}
}

func (r *TransactionRepositoryImpl) GetByID(id int64) (*models.Transaction, error) {
    var transaction models.Transaction
    err := r.DB.QueryRow("SELECT id, user_id, product_id, quantity FROM transactions WHERE id = ?", id).Scan(
        &transaction.ID, &transaction.UserID, &transaction.ProductID, &transaction.Quantity,
    )
    if err != nil {
        return nil, err
    }
    return &transaction, nil
}

func (r *TransactionRepositoryImpl) Store(transaction *models.Transaction) error {
    _, err := r.DB.Exec("INSERT INTO transactions (user_id, product_id, quantity) VALUES (?, ?, ?)",
        transaction.UserID, transaction.ProductID, transaction.Quantity)
    return err
}

func (r *TransactionRepositoryImpl) Update(transaction *models.Transaction) error {
    _, err := r.DB.Exec("UPDATE transactions SET user_id = ?, product_id = ?, quantity = ? WHERE id = ?",
        transaction.UserID, transaction.ProductID, transaction.Quantity, transaction.ID)
    return err
}

func (r *TransactionRepositoryImpl) Delete(id int64) error {
    _, err := r.DB.Exec("DELETE FROM transactions WHERE id = ?", id)
    return err
}

func (r *TransactionRepositoryImpl) Fetch() ([]models.Transaction, error) {
    rows, err := r.DB.Query("SELECT id, user_id, product_id, quantity FROM transactions")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var transactions []models.Transaction
    for rows.Next() {
        var transaction models.Transaction
        if err := rows.Scan(&transaction.ID, &transaction.UserID, &transaction.ProductID, &transaction.Quantity); err != nil {
            return nil, err
        }
        transactions = append(transactions, transaction)
    }
    return transactions, nil
}
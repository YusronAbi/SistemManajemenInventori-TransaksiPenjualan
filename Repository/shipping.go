package repository

import (
	"database/sql"
)

type ShippingRepositoryImpl struct {
	DB *sql.DB
}

func NewShippingRepository(db *sql.DB) models.ShippingRepository {
	return &ShippingRepositoryImpl{DB: db}
}

func (r *ShippingRepositoryImpl) GetByID(id int64) (*models.Shipping, error) {
	var shipping models.Shipping
	err := r.DB.QueryRow("SELECT id, transaction_id, address, status FROM shippings WHERE id = ?", id).Scan(
		&shipping.ID, &shipping.TransactionID, &shipping.Address, &shipping.Status,
	)
	if err != nil {
		return nil, err
	}
	return &shipping, nil
}

func (r *ShippingRepositoryImpl) Store(shipping *models.Shipping) error {
	_, err := r.DB.Exec("INSERT INTO shippings (transaction_id, address, status) VALUES (?, ?, ?)",
		shipping.TransactionID, shipping.Address, shipping.Status)
	return err
}

func (r *ShippingRepositoryImpl) Update(shipping *models.Shipping) error {
	_, err := r.DB.Exec("UPDATE shippings SET transaction_id = ?, address = ?, status = ? WHERE id = ?",
		shipping.TransactionID, shipping.Address, shipping.Status, shipping.ID)
	return err
}

func (r *ShippingRepositoryImpl) Delete(id int64) error {
	_, err := r.DB.Exec("DELETE FROM shippings WHERE id = ?", id)
	return err
}

func (r *ShippingRepositoryImpl) Fetch() ([]models.Shipping, error) {
	rows, err := r.DB.Query("SELECT id, transaction_id, address, status FROM shippings")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var shippings []models.Shipping
	for rows.Next() {
		var shipping models.Shipping
		if err := rows.Scan(&shipping.ID, &shipping.TransactionID, &shipping.Address, &shipping.Status); err != nil {
			return nil, err
		}
		shippings = append(shippings, shipping)
	}
	return shippings, nil
}

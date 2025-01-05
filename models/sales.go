package models

type Sale struct {
	ID            uint    `json:"id"`
	CustomerID    uint    `json:"customer_id"`
	SaleDate      string  `json:"sale_date"`
	TotalAmount   float64 `json:"total_amount"`
	PaymentStatus string  `json:"payment_status"`
}

func (Sale) TableName() string {
	return "sales"
}

func CreateSale(sale *Sale) error {
	if err := db.Create(&sale).Error; err != nil {
		return err
	}
	return nil
}

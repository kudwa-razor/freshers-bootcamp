package models

type Product struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	ProductName string `json:"product_name"`
	Price       int    `json:"price"`
	Quantity    int    `json:"quantity"`
}

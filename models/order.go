package models

type Order struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	ProductID uint   `json:"product_id"`
	Quantity  int    `json:"quantity"`
	Status    string `json:"status" gorm:"default:'order placed'"`
}

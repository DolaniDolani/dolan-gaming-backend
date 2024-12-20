package models

import "time"

type Game struct {
	ID            int64      `gorm:"primaryKey" json:"id"`
	Name          string     `json:"name"`
	Condition     string     `json:"condition"`
	Console       string     `json:"console"`
	Language      string     `json:"language"`
	Notes         string     `json:"notes"`
	PurchaseDate  *time.Time `json:"purchase_date"`
	PurchasePrice *float64   `json:"purchase_price"`
	SaleDate      *time.Time `json:"sale_date"`
	SalePrice     *float64   `json:"sale_price"`
	PurchaseID    int64      `json:"purchase_id`
}

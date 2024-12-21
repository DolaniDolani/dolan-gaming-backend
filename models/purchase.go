package models

import "time"

type Purchase struct {
	ID    int64     `gorm:"primaryKey" json:"id"`
	Name  string    `gorm:"not null" json:"name"`
	Date  time.Time `gorm:"not null" json:"date"`
	Cost  float64   `gorm:"not null" json:"cost"`
	Notes string    `json:"notes"`
	Games []Game    `gorm:"foreignKey:PurchaseID; not null" json:"games"`
}

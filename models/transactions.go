package models

import "time"

type Transaction struct {
	ID         uint `gorm:"primaryKey"`
	UserID     uint
	MerchantID uint
	Amount     float64
	Status     string `gorm:"size:50"`
	CreatedAt  time.Time
}

package models

import "time"

type Entrada struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Valor     float64   `json:"valor"`
	CreatedAt time.Time `json:"created_at"`
}

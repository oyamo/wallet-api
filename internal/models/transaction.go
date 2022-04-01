package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Transaction struct {
	TransactionID uuid.UUID `json:"transaction_id" gorm:"primaryKey" db:"transaction_id"`
	Type          string    `json:"type"`
	Reason        string    `json:"reason"`
	Amount        float64
	WalletID      uuid.UUID `json:"wallet_id"`
	Date          time.Time `json:"date"`
}

func (tx *Transaction) BeforeCreate(db *gorm.DB) (err error) {
	tx.TransactionID = uuid.New()
	tx.Date = time.Now()
	return nil
}

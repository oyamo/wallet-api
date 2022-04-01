package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Wallet struct {
	WalletID  uuid.UUID `json:"wallet_id" gorm:"primaryKey" db:"wallet_id" redis:"wallet_id"`
	UserID    uuid.UUID `json:"user_id" db:"user_id" redis:"user_id" validate:"required"`
	Balance   float64   `json:"balance" db:"balance" redis:"balance"`
	CreatedAt time.Time `json:"created_at,omitempty" db:"created_at" redis:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" db:"updated_at" redis:"updated_at"`
}

func (u *Wallet) BeforeCreate(tx *gorm.DB) (err error) {
	u.WalletID = uuid.New()
	u.Balance = 0
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return
}

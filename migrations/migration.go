package migrations

import (
	"github.com/oyamo/wallet-api/internal/models"
	"gorm.io/gorm"
)

func AutoMigrate(mysql *gorm.DB) error {
	return mysql.AutoMigrate(&models.User{}, &models.Wallet{}, &models.Transaction{})
}

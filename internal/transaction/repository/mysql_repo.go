package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/oyamo/wallet-api/internal/models"
	"github.com/oyamo/wallet-api/internal/transaction"
	"github.com/oyamo/wallet-api/pkg/util"
	"gorm.io/gorm"
	"net/http"
)

type mySQLRepo struct {
	db *gorm.DB
}

func (m mySQLRepo) Create(ctx context.Context, transaction *models.Transaction) (*models.Transaction, error) {
	res := m.db.Create(transaction)
	if res.Error != nil {
		return transaction, util.NewError(http.StatusInternalServerError, "internal server error: "+res.Error.Error())
	}
	return transaction, nil
}

func (m mySQLRepo) Update(ctx context.Context, transaction *models.Transaction) (*models.Transaction, error) {
	res := m.db.Save(transaction)
	return transaction, res.Error
}

func (m mySQLRepo) Delete(ctx context.Context, transactionID uuid.UUID) error {
	return m.db.Where("transaction_id = ?", transactionID.String()).Delete(&models.Transaction{}).Error
}

func (m mySQLRepo) GetByID(ctx context.Context, transactionID uuid.UUID) (*models.Transaction, error) {
	var tx models.Transaction
	res := m.db.Where("transaction_id = ?", transactionID.String()).First(&tx)

	if res.Error != nil {
		return nil, util.NewError(http.StatusNotFound, "transaction_id "+transactionID.String()+" not available")
	}

	return &tx, nil
}

func (m mySQLRepo) GetByWalletId(ctx context.Context, walletID uuid.UUID) (*[]models.Transaction, error) {
	var tx []models.Transaction
	res := m.db.Where("wallet_id = ?", walletID.String()).Find(&tx)

	if res.Error != nil {
		return nil, util.NewError(http.StatusNotFound, "wallet_id "+walletID.String()+" not available")
	}

	return &tx, nil
}

func NewTXRepo(database *gorm.DB) transaction.Repository {
	return mySQLRepo{database}
}

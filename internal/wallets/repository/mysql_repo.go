package repository

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/oyamo/wallet-api/internal/models"
	"github.com/oyamo/wallet-api/internal/wallets"
	"github.com/oyamo/wallet-api/pkg/util"
	"gorm.io/gorm"
	"net/http"
)

type walletsRepo struct {
	db *gorm.DB
}

func (w *walletsRepo) Create(ctx context.Context, wallet *models.Wallet) (*models.Wallet, error) {
	if _, err := w.GetByUserId(ctx, wallet.UserID); err != nil {
		return wallet, util.NewError(http.StatusUnauthorized, "wallet with the user id exists")
	}

	res := w.db.Create(wallet)
	if res.Error != nil {
		return wallet, util.NewError(http.StatusInternalServerError, "internal server error")
	}

	return wallet, nil
}

func (w *walletsRepo) Update(ctx context.Context, wallet *models.Wallet) (*models.Wallet, error) {
	res := w.db.Save(wallet)
	return wallet, res.Error
}

func (w *walletsRepo) Delete(ctx context.Context, walletID uuid.UUID) error {
	return w.db.Where("wallet_id = ?", walletID.String()).Delete(&models.Wallet{}).Error
}

func (w *walletsRepo) GetByID(ctx context.Context, walletID uuid.UUID) (*models.Wallet, error) {
	var wallet models.Wallet
	res := w.db.Where("wallet_id = ?", walletID.String()).First(&wallet)

	if res.Error != nil {
		return nil, errors.New("wallet_id " + walletID.String() + " not available")
	}

	return &wallet, nil
}

func (w *walletsRepo) GetByUserId(ctx context.Context, userId uuid.UUID) (*models.Wallet, error) {
	var wallet models.Wallet
	res := w.db.Where("user_id = ?", userId.String()).First(&wallet)

	if res.Error != nil {
		return nil, errors.New("No transactions.go for " + userId.String())
	}

	return &wallet, nil
}

func NewWalletRepo(db *gorm.DB) wallets.Repository {
	return &walletsRepo{db}
}

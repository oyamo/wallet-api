//go:generate mockgen -source usecase.go -destination mock/usecase_mock.go -package mock

package wallets

import (
	"context"
	"github.com/google/uuid"
	"github.com/oyamo/wallet-api/internal/models"
)

type Usecase interface {
	Create(ctx context.Context, wallet *models.Wallet) (*models.Wallet, error)
	Update(ctx context.Context, wallet *models.Wallet) (*models.Wallet, error)
	Delete(ctx context.Context, walletID uuid.UUID) error
	GetByWalletID(ctx context.Context, walletID uuid.UUID) (*models.Wallet, error)
	GetByUserID(ctx context.Context, userID uuid.UUID) (*models.Wallet, error)
	GetBalanceByID(ctx context.Context, walletID uuid.UUID) (*models.Wallet, error)
	Debit(ctx context.Context, walletID uuid.UUID, amount float64) (*models.Wallet, error)
	Credit(ctx context.Context, walletID uuid.UUID, amount float64) (*models.Wallet, error)
}

//go:generate mockgen -source mysql_repo.go -destination mock/mysql_repo_mock.go -package mock

package wallets

import (
	"context"
	"github.com/google/uuid"
	"github.com/oyamo/wallet-api/internal/models"
)

type Repository interface {
	Create(ctx context.Context, wallet *models.Wallet) (*models.Wallet, error)
	Update(ctx context.Context, wallet *models.Wallet) (*models.Wallet, error)
	Delete(ctx context.Context, walletID uuid.UUID) error
	GetByID(ctx context.Context, walletID uuid.UUID) (*models.Wallet, error)
	GetByUserId(ctx context.Context, userId uuid.UUID) (*models.Wallet, error)
}

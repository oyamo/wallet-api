//go:generate mockgen -source mysql_repo.go -destination mock/mysql_repo_mock.go -package mock

package transaction

import (
	"context"
	"github.com/google/uuid"
	"github.com/oyamo/wallet-api/internal/models"
)

type Repository interface {
	Create(ctx context.Context, transaction *models.Transaction) (*models.Transaction, error)
	Update(ctx context.Context, transaction *models.Transaction) (*models.Transaction, error)
	Delete(ctx context.Context, transactionID uuid.UUID) error
	GetByID(ctx context.Context, transactionID uuid.UUID) (*models.Transaction, error)
	GetByWalletId(ctx context.Context, walletID uuid.UUID) (*[]models.Transaction, error)
}

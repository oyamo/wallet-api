//go:generate mockgen -source usecase.go -destination mock/usecase_mock.go -package mock

package transaction

import (
	"context"
	"github.com/google/uuid"
	"github.com/oyamo/wallet-api/internal/models"
)

type Usecase interface {
	Create(ctx context.Context, transaction *models.Transaction) (*models.Transaction, error)
	Update(ctx context.Context, transaction *models.Transaction) (*models.Transaction, error)
	Delete(ctx context.Context, txID uuid.UUID) error
	GetByID(ctx context.Context, txID uuid.UUID) (*models.Transaction, error)
	GetByWalletID(ctx context.Context, walletID uuid.UUID) (*[]models.Transaction, error)
}

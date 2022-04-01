//go:generate mockgen -source redis_repo.go -destination mock/redis_repo_mock.go -package mock

package wallets

import (
	"context"
	"github.com/oyamo/wallet-api/internal/models"
)

// Auth Redis repository interface
type RedisRepository interface {
	GetByIDCtx(ctx context.Context, key string) (*models.Wallet, error)
	SetWalletCtx(ctx context.Context, key string, seconds int, user *models.Wallet) error
	DeleteWalletCtx(ctx context.Context, key string) error
}

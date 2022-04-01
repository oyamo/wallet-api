//go:generate mockgen -source redis_repo.go -destination mock/redis_repo_mock.go -package mock

package transaction

import (
	"context"
	"github.com/oyamo/wallet-api/internal/models"
)

// RedisRepository Auth Redis repository interface
type RedisRepository interface {
	GetByIDCtx(ctx context.Context, key string) (*models.Transaction, error)
	SetTxCtx(ctx context.Context, key string, seconds int, user *models.Transaction) error
	DeleteTxCtx(ctx context.Context, key string) error
}

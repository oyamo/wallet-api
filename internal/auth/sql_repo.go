//go:generate mockgen -source sql_repo.go -destination mock/sql_repo_mock.go -package mock
package auth

import (
	"context"
	"github.com/google/uuid"
	"github.com/oyamo/wallet-api/internal/models"
)

// Auth repository interface
type Repository interface {
	Register(ctx context.Context, user *models.User) (*models.User, error)
	Update(ctx context.Context, user *models.User) (*models.User, error)
	Delete(ctx context.Context, userID uuid.UUID) error
	GetByID(ctx context.Context, userID uuid.UUID) (*models.User, error)
	FindByEmail(ctx context.Context, user *models.User) (*models.User, error)
}

package usecase

import (
	"context"
	"github.com/google/uuid"
	"github.com/oyamo/wallet-api/config"
	"github.com/oyamo/wallet-api/internal/models"
	"github.com/oyamo/wallet-api/internal/transaction"
)

const (
	TX_LIFE = 3600
)

// Auth UseCase
type txUC struct {
	cfg       *config.Config
	txRepo    transaction.Repository
	redisRepo transaction.RedisRepository
}

func (t txUC) Create(ctx context.Context, transaction *models.Transaction) (*models.Transaction, error) {
	createdTx, err := t.txRepo.Create(ctx, transaction)

	if err != nil {
		return nil, err
	}

	return createdTx, nil
}

func (t txUC) Update(ctx context.Context, transaction *models.Transaction) (*models.Transaction, error) {
	updatedTx, err := t.txRepo.Update(ctx, transaction)

	if err != nil {
		return nil, err
	}

	return updatedTx, nil
}

func (t txUC) Delete(ctx context.Context, txID uuid.UUID) error {
	err := t.txRepo.Delete(ctx, txID)

	if err != nil {
		return err
	}

	return nil
}

func (t txUC) GetByID(ctx context.Context, txID uuid.UUID) (*models.Transaction, error) {
	txFound, err := t.txRepo.GetByID(ctx, txID)

	if err != nil {
		return nil, err
	}

	return txFound, nil
}

func (t txUC) GetByWalletID(ctx context.Context, walletID uuid.UUID) (*[]models.Transaction, error) {
	foundTx, err := t.txRepo.GetByWalletId(ctx, walletID)

	if err != nil {
		return nil, err
	}

	return foundTx, nil
}

func NewTxUC(cfg *config.Config, txRepo transaction.Repository, redisRepo transaction.RedisRepository) transaction.Usecase {
	return txUC{
		cfg:       cfg,
		txRepo:    txRepo,
		redisRepo: redisRepo,
	}
}

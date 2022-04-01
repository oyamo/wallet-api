package repository

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/opentracing/opentracing-go"
	"github.com/oyamo/wallet-api/internal/models"
	"github.com/oyamo/wallet-api/internal/transaction"
	"github.com/pkg/errors"
	"time"
)

type redisRepo struct {
	db *redis.Client
}

func (r redisRepo) GetByIDCtx(ctx context.Context, key string) (*models.Transaction, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "transactionRedisRepo.GetByIDCtx")
	defer span.Finish()

	transactionBytes, err := r.db.Get(ctx, key).Bytes()
	if err != nil {
		return nil, errors.Wrap(err, "transactionRedisRepo.GetByIDCtx.redisClient.Get")
	}
	txFound := &models.Transaction{}
	if err = json.Unmarshal(transactionBytes, txFound); err != nil {
		return nil, errors.Wrap(err, "transactionRedisRepo.GetByIDCtx.json.Unmarshal")
	}
	return txFound, nil
}

func (r redisRepo) SetTxCtx(ctx context.Context, key string, seconds int, transaction *models.Transaction) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "transactionRedisRepo.SetTransactionCtx")
	defer span.Finish()

	transactionBytes, err := json.Marshal(transaction)
	if err != nil {
		return errors.Wrap(err, "transactionRedisRepo.SetTransactionCtx.json.Unmarshal")
	}
	if err = r.db.Set(ctx, key, transactionBytes, time.Second*time.Duration(seconds)).Err(); err != nil {
		return errors.Wrap(err, "transactionRedisRepo.SetTransactionCtx.redisClient.Set")
	}
	return nil
}

func (r redisRepo) DeleteTxCtx(ctx context.Context, key string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "transactionRedisRepo.DeleteTransactionCtx")
	defer span.Finish()

	if err := r.db.Del(ctx, key).Err(); err != nil {
		return errors.Wrap(err, "transactionRedisRepo.DeleteTransactionCtx.redisClient.Del")
	}
	return nil
}

func NewRedisRepo(db *redis.Client) transaction.RedisRepository {
	return &redisRepo{db: db}
}

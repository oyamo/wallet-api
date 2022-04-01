package repository

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/opentracing/opentracing-go"
	"github.com/oyamo/wallet-api/internal/models"
	"github.com/oyamo/wallet-api/internal/wallets"
	"github.com/pkg/errors"
	"time"
)

type redisRepo struct {
	db *redis.Client
}

func (r redisRepo) GetByIDCtx(ctx context.Context, key string) (*models.Wallet, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "walletRedisRepo.GetByIDCtx")
	defer span.Finish()

	walletBytes, err := r.db.Get(ctx, key).Bytes()
	if err != nil {
		return nil, errors.Wrap(err, "walletRedisRepo.GetByIDCtx.redisClient.Get")
	}
	wallet := &models.Wallet{}
	if err = json.Unmarshal(walletBytes, wallet); err != nil {
		return nil, errors.Wrap(err, "walletRedisRepo.GetByIDCtx.json.Unmarshal")
	}
	return wallet, nil
}

func (r redisRepo) SetWalletCtx(ctx context.Context, key string, seconds int, wallet *models.Wallet) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "walletRedisRepo.SetWalletCtx")
	defer span.Finish()

	walletBytes, err := json.Marshal(wallet)
	if err != nil {
		return errors.Wrap(err, "walletRedisRepo.SetWalletCtx.json.Unmarshal")
	}
	if err = r.db.Set(ctx, key, walletBytes, time.Second*time.Duration(seconds)).Err(); err != nil {
		return errors.Wrap(err, "walletRedisRepo.SetWalletCtx.redisClient.Set")
	}
	return nil
}

func (r redisRepo) DeleteWalletCtx(ctx context.Context, key string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "walletRedisRepo.DeleteWalletCtx")
	defer span.Finish()

	if err := r.db.Del(ctx, key).Err(); err != nil {
		return errors.Wrap(err, "walletRedisRepo.DeleteWalletCtx.redisClient.Del")
	}
	return nil
}

func NewRedisRepo(db *redis.Client) wallets.RedisRepository {
	return &redisRepo{db: db}
}

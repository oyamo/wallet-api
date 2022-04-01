package usecase

import (
	"context"
	"github.com/google/uuid"
	"github.com/oyamo/wallet-api/config"
	"github.com/oyamo/wallet-api/internal/models"
	"github.com/oyamo/wallet-api/internal/wallets"
	"github.com/oyamo/wallet-api/pkg/util"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// Auth UseCase
type walletUC struct {
	cfg        *config.Config
	walletRepo wallets.Repository
	redisRepo  wallets.RedisRepository
}

func (w walletUC) GetByUserID(ctx context.Context, userID uuid.UUID) (*models.Wallet, error) {

	walletFound, err := w.walletRepo.GetByUserId(ctx, userID)
	if err != nil || walletFound == nil {
		return nil, util.NewError(http.StatusNotFound, "not found")
	}

	return walletFound, nil
}

const (
	WALLET_LIFE = 3600
)

func (w walletUC) Create(ctx context.Context, wallet *models.Wallet) (*models.Wallet, error) {
	createdWallet, err := w.walletRepo.Create(ctx, wallet)
	if err != nil {
		return nil, err
	}

	err = w.redisRepo.SetWalletCtx(ctx, createdWallet.WalletID.String(), WALLET_LIFE, createdWallet)

	if err != nil {
		log.Errorf("cannot save wallet to redis: %v", err.Error())
		return nil, util.NewError(http.StatusInternalServerError, "internal sever error")
	}

	return createdWallet, nil
}

func (w walletUC) Update(ctx context.Context, wallet *models.Wallet) (*models.Wallet, error) {
	createdWallet, err := w.walletRepo.Update(ctx, wallet)
	if err != nil {
		return nil, err
	}

	err = w.redisRepo.SetWalletCtx(ctx, createdWallet.WalletID.String(), WALLET_LIFE, createdWallet)

	if err != nil {
		log.Errorf("cannot save wallet to redis: %v", err.Error())
		return nil, util.NewError(http.StatusInternalServerError, "internal sever error")
	}

	return createdWallet, nil
}

func (w walletUC) Delete(ctx context.Context, walletID uuid.UUID) error {
	err := w.walletRepo.Delete(ctx, walletID)
	if err != nil {
		log.Errorf("cannot delete wallet: %v", err.Error())
		return util.NewError(http.StatusInternalServerError, "internal server error")
	}

	err = w.redisRepo.DeleteWalletCtx(ctx, walletID.String())

	if err != nil {
		log.Errorf("cannot delete wallet in redis: %v", err.Error())
		return util.NewError(http.StatusInternalServerError, "internal sever error")
	}

	return nil
}

func (w walletUC) GetByWalletID(ctx context.Context, walletID uuid.UUID) (*models.Wallet, error) {
	walletFound, err := w.redisRepo.GetByIDCtx(ctx, walletID.String())
	if err == nil && walletFound != nil {
		return walletFound, nil
	}

	walletFound, err = w.walletRepo.GetByID(ctx, walletID)
	if err != nil || walletFound == nil {
		return nil, util.NewError(http.StatusNotFound, "not found")
	}

	return walletFound, nil
}

func (w walletUC) GetBalanceByID(ctx context.Context, walletID uuid.UUID) (*models.Wallet, error) {
	wallet, err := w.GetByWalletID(ctx, walletID)

	if err != nil {
		return nil, err
	}

	return wallet, nil
}

func (w walletUC) Debit(ctx context.Context, walletID uuid.UUID, amount float64) (*models.Wallet, error) {
	wallet, err := w.GetByWalletID(ctx, walletID)

	if err != nil {
		return nil, err
	}

	if amount <= 0 {
		return nil, util.NewError(http.StatusBadRequest, "amount must be gte 0")
	}

	newBalance := wallet.Balance - amount

	if newBalance < 0 {
		return nil, util.NewError(http.StatusExpectationFailed, "amount exceeds balance")
	}

	wallet.Balance = newBalance

	update, err := w.Update(ctx, wallet)
	if err != nil {
		log.Errorf("error debiting: %v", err)
		return nil, util.NewError(http.StatusInternalServerError, "internal server error")
	}

	return update, nil
}

func (w walletUC) Credit(ctx context.Context, walletID uuid.UUID, amount float64) (*models.Wallet, error) {
	wallet, err := w.GetByWalletID(ctx, walletID)

	if err != nil {
		return nil, err
	}

	if amount <= 0 {
		return nil, util.NewError(http.StatusBadRequest, "amount must be gte 0")
	}

	newBalance := wallet.Balance + amount
	wallet.Balance = newBalance

	update, err := w.Update(ctx, wallet)
	if err != nil {
		log.Errorf("error debiting: %v", err)
		return nil, util.NewError(http.StatusInternalServerError, "internal server error")
	}

	return update, nil
}

// Auth UseCase constructor
func NewWalletUseCase(cfg *config.Config, walletsRepo wallets.Repository, redisRepo wallets.RedisRepository) wallets.Usecase {
	return &walletUC{cfg: cfg, walletRepo: walletsRepo, redisRepo: redisRepo}
}

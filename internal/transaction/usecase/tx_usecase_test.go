package usecase

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/oyamo/wallet-api/config"
	"github.com/oyamo/wallet-api/internal/models"
	"github.com/oyamo/wallet-api/internal/transaction"
	"github.com/oyamo/wallet-api/internal/transaction/mock"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
	"time"
)

func Test_txUC_Create(t1 *testing.T) {
	t1.Parallel()

	ctrl := gomock.NewController(t1)
	defer ctrl.Finish()

	ctx := context.Background()

	mockTxRepo := mock.NewMockRepository(ctrl)
	mockRedis := mock.NewMockRedisRepository(ctrl)

	mockTxUC := NewTxUC(nil, mockTxRepo, mockRedis)

	tx := models.Transaction{
		TransactionID: uuid.New(),
		Type:          "debit",
		Reason:        "buy assets",
		Amount:        100,
		WalletID:      uuid.New(),
		Date:          time.Now(),
	}

	var createdTx models.Transaction

	mockTxRepo.EXPECT().Create(ctx, &tx).Return(&createdTx, nil)
	create, err := mockTxUC.Create(ctx, &tx)

	require.NoError(t1, err)
	require.NotNil(t1, create)

}

func Test_txUC_Delete(t1 *testing.T) {
	t1.Parallel()

	ctrl := gomock.NewController(t1)
	defer ctrl.Finish()

	ctx := context.Background()

	mockTxRepo := mock.NewMockRepository(ctrl)
	mockRedis := mock.NewMockRedisRepository(ctrl)

	mockTxUC := NewTxUC(nil, mockTxRepo, mockRedis)

	tx := models.Transaction{
		TransactionID: uuid.New(),
		Type:          "debit",
		Reason:        "buy assets",
		Amount:        100,
		WalletID:      uuid.New(),
		Date:          time.Now(),
	}

	var createdTx models.Transaction

	mockTxRepo.EXPECT().Create(ctx, &tx).Return(&createdTx, nil)
	create, err := mockTxUC.Create(ctx, &tx)

	require.NoError(t1, err)
	require.NotNil(t1, create)
	require.NotNil(t1, create.TransactionID)

}

func Test_txUC_GetByID(t1 *testing.T) {
	type fields struct {
		cfg       *config.Config
		txRepo    transaction.Repository
		redisRepo transaction.RedisRepository
	}
	type args struct {
		ctx  context.Context
		txID uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Transaction
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := txUC{
				cfg:       tt.fields.cfg,
				txRepo:    tt.fields.txRepo,
				redisRepo: tt.fields.redisRepo,
			}
			got, err := t.GetByID(tt.args.ctx, tt.args.txID)
			if (err != nil) != tt.wantErr {
				t1.Errorf("GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("GetByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_txUC_GetByWalletID(t1 *testing.T) {
	type fields struct {
		cfg       *config.Config
		txRepo    transaction.Repository
		redisRepo transaction.RedisRepository
	}
	type args struct {
		ctx      context.Context
		walletID uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *[]models.Transaction
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := txUC{
				cfg:       tt.fields.cfg,
				txRepo:    tt.fields.txRepo,
				redisRepo: tt.fields.redisRepo,
			}
			got, err := t.GetByWalletID(tt.args.ctx, tt.args.walletID)
			if (err != nil) != tt.wantErr {
				t1.Errorf("GetByWalletID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("GetByWalletID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_txUC_Update(t1 *testing.T) {
	type fields struct {
		cfg       *config.Config
		txRepo    transaction.Repository
		redisRepo transaction.RedisRepository
	}
	type args struct {
		ctx         context.Context
		transaction *models.Transaction
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Transaction
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := txUC{
				cfg:       tt.fields.cfg,
				txRepo:    tt.fields.txRepo,
				redisRepo: tt.fields.redisRepo,
			}
			got, err := t.Update(tt.args.ctx, tt.args.transaction)
			if (err != nil) != tt.wantErr {
				t1.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("Update() got = %v, want %v", got, tt.want)
			}
		})
	}
}

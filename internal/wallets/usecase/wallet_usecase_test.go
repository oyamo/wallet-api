package usecase

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/oyamo/wallet-api/internal/models"
	"github.com/oyamo/wallet-api/internal/wallets/mock"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func Test_walletUC_Create(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	mockUc := mock.NewMockUsecase(ctrl)
	createdWallet := models.Wallet{
		WalletID:  uuid.New(),
		UserID:    uuid.New(),
		Balance:   0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mockUc.EXPECT().Create(ctx, &createdWallet).Return(&createdWallet, nil)
	returnedWallet, err := mockUc.Create(ctx, &createdWallet)
	require.NoError(t, err)
	require.Equal(t, createdWallet, *returnedWallet)
}

func Test_walletUC_Credit(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	mockUc := mock.NewMockUsecase(ctrl)
	createdWallet := models.Wallet{
		WalletID:  uuid.New(),
		UserID:    uuid.New(),
		Balance:   0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mockUc.EXPECT().Create(ctx, &createdWallet).Return(&createdWallet, nil)
	returnedWallet, err := mockUc.Create(ctx, &createdWallet)
	require.NoError(t, err)
	require.Equal(t, createdWallet, *returnedWallet)

	mockUc.EXPECT().Credit(ctx, createdWallet.WalletID, 20.0).Return(&createdWallet, nil).Times(1)
	returnedWallet, err = mockUc.Credit(ctx, createdWallet.WalletID, 20.0)
	require.NoError(t, err)
	require.Equal(t, createdWallet, *returnedWallet)
}

func Test_walletUC_Debit(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	mockUc := mock.NewMockUsecase(ctrl)
	createdWallet := models.Wallet{
		WalletID:  uuid.New(),
		UserID:    uuid.New(),
		Balance:   0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mockUc.EXPECT().Create(ctx, &createdWallet).Return(&createdWallet, nil)
	returnedWallet, err := mockUc.Create(ctx, &createdWallet)
	require.NoError(t, err)
	require.Equal(t, createdWallet, *returnedWallet)

	mockUc.EXPECT().Credit(ctx, createdWallet.WalletID, 20.0).Return(&createdWallet, nil).Times(1)
	returnedWallet, err = mockUc.Credit(ctx, createdWallet.WalletID, 20.0)
	require.NoError(t, err)
	require.Equal(t, createdWallet, *returnedWallet)

	mockUc.EXPECT().Debit(ctx, createdWallet.WalletID, 10.0).Return(&createdWallet, nil).Times(1)
	returnedWallet, err = mockUc.Debit(ctx, createdWallet.WalletID, 10.0)
	require.NoError(t, err)
	require.Equal(t, createdWallet, *returnedWallet)
	//require.Equal(t, float64(10.0), returnedWallet.Balance)
}

func Test_walletUC_Delete(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	mockUc := mock.NewMockUsecase(ctrl)
	createdWallet := models.Wallet{
		WalletID:  uuid.New(),
		UserID:    uuid.New(),
		Balance:   0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mockUc.EXPECT().Create(ctx, &createdWallet).Return(&createdWallet, nil)
	returnedWallet, err := mockUc.Create(ctx, &createdWallet)
	require.NoError(t, err)
	require.Equal(t, createdWallet, *returnedWallet)

	mockUc.EXPECT().Delete(ctx, createdWallet.WalletID).Return(nil)
	err = mockUc.Delete(ctx, createdWallet.WalletID)
	require.NoError(t, err)
}

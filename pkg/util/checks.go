package util

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/oyamo/wallet-api/internal/models"
	"github.com/oyamo/wallet-api/internal/wallets"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func WalletVerification(ctx *gin.Context, walletId uuid.UUID, walletUC wallets.Usecase) *HTTPError {
	userStruct, found := ctx.Get("user")

	if !found || userStruct == nil {
		log.Errorf("pkg/util/checks::WalletVerification: user not in context")
		return NewError(http.StatusNotFound, "wallet not found")
	}

	user := userStruct.(*models.User)
	userWallet, err := walletUC.GetByUserID(context.Background(), user.UserID)

	if err != nil {
		log.Errorf("pkg/util/checks::WalletVerification: userId: %v: %v", user.UserID, err)
		return NewError(http.StatusInternalServerError, "internal server error: ")
	}

	_ = userWallet

	if walletId.String() != userWallet.WalletID.String() {
		log.Errorf("pkg/util/checks::WalletVerification: trying to get non owned wallet")
		return NewError(http.StatusNotFound, "wallet not found")
	}

	return nil

}

package http

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/oyamo/wallet-api/config"
	"github.com/oyamo/wallet-api/internal/auth"
	"github.com/oyamo/wallet-api/internal/models"
	"github.com/oyamo/wallet-api/internal/session"
	"github.com/oyamo/wallet-api/internal/transaction"
	"github.com/oyamo/wallet-api/internal/wallets"
	"github.com/oyamo/wallet-api/pkg/util"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type txHandlers struct {
	cfg      *config.Config
	authUC   auth.UseCase
	sessUC   session.UCSession
	walletUC wallets.Usecase
	txUC     transaction.Usecase
}

func (t txHandlers) GetByWalletID() gin.HandlerFunc {
	response := models.Response[interface{}]{}

	return func(ctx *gin.Context) {
		walletID, err := uuid.Parse(ctx.Param("wallet_id"))

		if err != nil {
			util.NewError(http.StatusBadRequest, "wallet_id parameter needed").WriteJSONToCtx(ctx)
			return
		}

		verificationErr := util.WalletVerification(ctx, walletID, t.walletUC)

		if err != nil {
			verificationErr.WriteJSONToCtx(ctx)
			return
		}

		var transactions *[]models.Transaction
		transactions, err = t.txUC.GetByWalletID(context.Background(), walletID)

		if err != nil {
			log.Errorf("transaction/delivery/http/handlers::GetByWalletID: %v", err)
			util.NewError(http.StatusNotFound, "tran").WriteJSONToCtx(ctx)
			return
		}

		response.ResponseCode = http.StatusOK
		response.Message = "transactions found"
		response.Data = transactions
		ctx.JSON(http.StatusOK, response)
	}
}

func NewTxHandler(cfg *config.Config, authUC auth.UseCase, walletUc wallets.Usecase, txUC transaction.Usecase, sessUC session.UCSession) transaction.Handler {
	return txHandlers{
		cfg:      cfg,
		authUC:   authUC,
		sessUC:   sessUC,
		walletUC: walletUc,
		txUC:     txUC,
	}
}

package http

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/oyamo/wallet-api/config"
	"github.com/oyamo/wallet-api/internal/models"
	"github.com/oyamo/wallet-api/internal/transaction"
	"github.com/oyamo/wallet-api/internal/wallets"
	"github.com/oyamo/wallet-api/pkg/util"
	"net/http"
)

type walletHandler struct {
	cfg      *config.Config
	WalletUC wallets.Usecase
	txUc     transaction.Usecase
}

func (w *walletHandler) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var wallet models.Wallet
		err := util.ReadRequest(ctx, &wallet)
		response := &models.Response[interface{}]{}

		if err != nil {
			response.Data = nil
			response.Message = err.Error()
			response.ResponseCode = http.StatusBadRequest
			ctx.JSON(http.StatusBadRequest, response)
			return
		}

		currentUser, exists := ctx.Get("user")

		if !exists {
			response.Data = nil
			response.Message = "Authentication failed"
			response.ResponseCode = http.StatusUnauthorized
			ctx.JSON(http.StatusUnauthorized, response)
			return
		}

		if currentUser.(*models.User).UserID != wallet.UserID {
			response.Data = nil
			response.Message = "Forbidden access"
			response.ResponseCode = http.StatusForbidden
			ctx.JSON(http.StatusForbidden, response)
			return
		}

		savedWallet, err := w.WalletUC.Create(context.Background(), &wallet)
		if err != nil {
			response.Data = nil
			response.Message = err.Error()
			response.ResponseCode = http.StatusInternalServerError
			ctx.JSON(http.StatusInternalServerError, response)
			return
		}

		response.Data = savedWallet
		response.Message = "success"
		ctx.JSON(http.StatusOK, response)

	}
}

func (w *walletHandler) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var wallet models.Wallet
		err := util.ReadRequest(ctx, &wallet)

		response := &models.Response[interface{}]{}

		if err != nil {
			response.Data = nil
			response.Message = err.Error()
			response.ResponseCode = http.StatusBadRequest
			ctx.JSON(http.StatusBadRequest, response)
			return
		}

		currentUser, exists := ctx.Get("user")

		if !exists {
			response.Data = nil
			response.Message = "Authentication failed"
			response.ResponseCode = http.StatusUnauthorized
			ctx.JSON(http.StatusUnauthorized, response)
			return
		}

		if currentUser.(*models.User).UserID != wallet.UserID {
			response.Data = nil
			response.Message = "Forbidden access"
			response.ResponseCode = http.StatusForbidden
			ctx.JSON(http.StatusForbidden, response)
			return
		}

		savedWallet, err := w.WalletUC.Update(context.Background(), &wallet)
		if err != nil {
			response.Data = nil
			response.Message = err.Error()
			response.ResponseCode = http.StatusInternalServerError
			ctx.JSON(http.StatusInternalServerError, response)
			return
		}

		response.Data = savedWallet
		response.Message = "success"
		ctx.JSON(http.StatusOK, response)
	}
}

func (w *walletHandler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		walletId, err := uuid.Parse(ctx.Param("wallet_id"))
		response := models.Response[any]{}

		if err != nil {
			response.ResponseCode = http.StatusBadRequest
			response.Data = nil
			response.Message = "missing parameter wallet_id: " + err.Error()
			ctx.JSON(response.ResponseCode, response)
			return
		}

		verificationErr := util.WalletVerification(ctx, walletId, w.WalletUC)

		if err != nil {
			verificationErr.WriteJSONToCtx(ctx)
			return
		}

		err = w.WalletUC.Delete(context.Background(), walletId)
		if err != nil {
			response.Data = nil
			response.ResponseCode = http.StatusInternalServerError
			response.Message = err.Error()
			ctx.JSON(response.ResponseCode, response)
			return
		}

		response.ResponseCode = http.StatusOK
		response.Data = nil
		response.Message = "deletion successful"
		ctx.JSON(response.ResponseCode, response)
	}
}

func (w *walletHandler) GetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		walletId, err := uuid.Parse(ctx.Param("wallet_id"))
		response := models.Response[interface{}]{}

		if err != nil {
			response.ResponseCode = http.StatusBadRequest
			response.Data = nil
			response.Message = "missing parameter wallet_id"
			ctx.JSON(response.ResponseCode, response)
			return
		}

		// verify wallet if it belongs to current user
		verificationErr := util.WalletVerification(ctx, walletId, w.WalletUC)

		if verificationErr != nil {
			verificationErr.WriteJSONToCtx(ctx)
			return
		}

		// Search for wallet
		foundWallet, err := w.WalletUC.GetByWalletID(context.Background(), walletId)

		if err != nil {
			response.Data = nil
			response.ResponseCode = http.StatusNotFound
			response.Message = err.Error()
			ctx.JSON(response.ResponseCode, response)
			return
		}

		response.ResponseCode = http.StatusOK
		response.Data = foundWallet
		response.Message = "search successful"
		ctx.JSON(response.ResponseCode, response)

	}
}

func (w *walletHandler) GetBalanceByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		walletUUID, err := uuid.Parse(c.Param("wallet_id"))

		if err != nil {
			util.NewError(http.StatusBadRequest, "wallet_id parameter not in request").WriteJSONToCtx(c)
			return
		}

		walletFound, err := w.WalletUC.GetBalanceByID(context.Background(), walletUUID)
		if err != nil {
			util.NewError(http.StatusBadRequest, "wallet_id not found").WriteJSONToCtx(c)
			return
		}

		c.JSON(http.StatusOK, models.Response[*models.Wallet]{
			ResponseCode: http.StatusOK,
			Message:      "wallet found successfully",
			Data:         walletFound,
		})
	}
}

func (w *walletHandler) Debit() gin.HandlerFunc {
	type tx struct {
		Amount float64 `json:"amount" validate:"required,gte=1"`
		Reason string  `json:"reason" validate:"required"`
	}
	return func(ctx *gin.Context) {
		var walletTX tx
		response := &models.Response[interface{}]{}

		walletId, err := uuid.Parse(ctx.Param("wallet_id"))

		if err != nil {
			response.Data = nil
			response.Message = err.Error()
			response.ResponseCode = http.StatusBadRequest
			ctx.JSON(http.StatusBadRequest, response)
			return
		}

		verificationErr := util.WalletVerification(ctx, walletId, w.WalletUC)

		if verificationErr != nil {
			verificationErr.WriteJSONToCtx(ctx)
			return
		}

		err = util.ReadRequest(ctx, &walletTX)
		if err != nil {
			response.Data = nil
			response.Message = err.Error()
			response.ResponseCode = http.StatusBadRequest
			ctx.JSON(http.StatusBadRequest, response)
			return
		}

		savedWallet, err := w.WalletUC.Debit(context.Background(), walletId, walletTX.Amount)
		if err != nil {
			err.(*util.HTTPError).WriteJSONToCtx(ctx)
			return
		}

		savedTx, err := w.txUc.Create(
			context.Background(),
			&models.Transaction{WalletID: savedWallet.WalletID, Amount: walletTX.Amount, Type: "debit", Reason: walletTX.Reason},
		)

		response.Data = struct {
			*models.Wallet
			transactionID uuid.UUID
		}{savedWallet, savedTx.TransactionID}

		response.Data = savedWallet
		response.Message = "success"
		response.ResponseCode = http.StatusOK
		ctx.JSON(http.StatusOK, response)
	}
}

func (w *walletHandler) Credit() gin.HandlerFunc {

	type tx struct {
		Amount float64 `json:"amount" validate:"required,gte=1"`
		Reason string  `json:"reason" validate:"required"`
	}

	return func(ctx *gin.Context) {
		var walletTX tx
		response := &models.Response[interface{}]{}
		walletId, err := uuid.Parse(ctx.Param("wallet_id"))

		if err != nil {
			response.Data = nil
			response.Message = err.Error()
			response.ResponseCode = http.StatusBadRequest
			ctx.JSON(http.StatusBadRequest, response)
			return
		}

		verificationErr := util.WalletVerification(ctx, walletId, w.WalletUC)

		if verificationErr != nil {
			verificationErr.WriteJSONToCtx(ctx)
			return
		}

		err = util.ReadRequest(ctx, &walletTX)
		if err != nil {
			response.Data = nil
			response.Message = err.Error()
			response.ResponseCode = http.StatusBadRequest
			ctx.JSON(http.StatusBadRequest, response)
			return
		}

		savedWallet, err := w.WalletUC.Credit(context.Background(), walletId, walletTX.Amount)
		if err != nil {
			err.(*util.HTTPError).WriteJSONToCtx(ctx)
			return
		}

		savedTx, err := w.txUc.Create(
			context.Background(),
			&models.Transaction{WalletID: savedWallet.WalletID, Amount: walletTX.Amount, Type: "credit", Reason: walletTX.Reason},
		)

		response.Data = struct {
			*models.Wallet
			transactionID uuid.UUID
		}{savedWallet, savedTx.TransactionID}

		response.Message = "success"
		response.ResponseCode = http.StatusOK
		ctx.JSON(http.StatusOK, response)
	}
}

func NewWalletHandler(cfg *config.Config, walletUC *wallets.Usecase, txUc transaction.Usecase) wallets.Handler {
	return &walletHandler{
		cfg:      cfg,
		WalletUC: *walletUC,
		txUc:     txUc,
	}
}

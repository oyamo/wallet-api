package http

import (
	"github.com/gin-gonic/gin"
	"github.com/oyamo/wallet-api/internal/middleware"
	"github.com/oyamo/wallet-api/internal/transaction"
)

func MapTxRoutes(group *gin.RouterGroup, h transaction.Handler, mw *middleware.MiddlewareManager) {
	group.GET("/bywallet/:wallet_id", mw.AuthSessionMiddleware, mw.CSRF, h.GetByWalletID())
}

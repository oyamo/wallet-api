package http

import (
	"github.com/gin-gonic/gin"
	"github.com/oyamo/wallet-api/internal/middleware"
	"github.com/oyamo/wallet-api/internal/wallets"
)

// Map wallets routes
func MapWalletRoutes(group *gin.RouterGroup, h wallets.Handler, mw *middleware.MiddlewareManager) {
	group.POST("", mw.AuthSessionMiddleware, mw.CSRF, h.Create())
	group.DELETE("/:wallet_id", mw.AuthSessionMiddleware, mw.CSRF, h.Delete())
	group.PUT("/:wallet_id", mw.AuthSessionMiddleware, mw.CSRF, h.Update())
	group.GET("/:wallet_id", mw.AuthSessionMiddleware, mw.CSRF, h.GetByID())
	group.GET("/:wallet_id/balance", mw.AuthSessionMiddleware, mw.CSRF, h.GetBalanceByID())
	group.POST("/:wallet_id/credit/", mw.AuthSessionMiddleware, mw.CSRF, h.Credit())
	group.POST("/:wallet_id/debit/", mw.AuthSessionMiddleware, mw.CSRF, h.Debit())
}

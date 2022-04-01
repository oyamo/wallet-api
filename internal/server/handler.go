package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	authHandler "github.com/oyamo/wallet-api/internal/auth/delivery/http"
	authRepo "github.com/oyamo/wallet-api/internal/auth/repository"
	"github.com/oyamo/wallet-api/internal/auth/usecase"
	"github.com/oyamo/wallet-api/internal/middleware"
	"github.com/oyamo/wallet-api/internal/session/repository"
	sessionUC "github.com/oyamo/wallet-api/internal/session/usecase"
	transactionHandlers "github.com/oyamo/wallet-api/internal/transaction/delivery/http"
	txRepo "github.com/oyamo/wallet-api/internal/transaction/repository"
	txUc "github.com/oyamo/wallet-api/internal/transaction/usecase"
	walletHandler "github.com/oyamo/wallet-api/internal/wallets/delivery/http"
	walletRepo "github.com/oyamo/wallet-api/internal/wallets/repository"
	walletUseCase "github.com/oyamo/wallet-api/internal/wallets/usecase"
	"net/http"
)

func (s *Server) MapHandlers(e *gin.Engine) error {

	v1 := e.Group("/api/v1")
	wallet := v1.Group("/wallets")
	auth := v1.Group("/auth")
	status := v1.Group("/status")
	transactions := v1.Group("/transactions")

	// init repositories
	walletRepository := walletRepo.NewWalletRepo(s.mysqlClient)
	walletRedisRepo := walletRepo.NewRedisRepo(s.redisClient)
	authRepository := authRepo.NewAuthRepository(s.mysqlClient)
	authRedisRepo := authRepo.NewAuthRedisRepo(s.redisClient)
	sessionRepo := repository.NewSessionRepository(s.redisClient, s.cfg)
	txRepository := txRepo.NewTXRepo(s.mysqlClient)
	txRedisRepo := txRepo.NewRedisRepo(s.redisClient)

	// init usecases
	walletUC := walletUseCase.NewWalletUseCase(s.cfg, walletRepository, walletRedisRepo)
	transactionUC := txUc.NewTxUC(s.cfg, txRepository, txRedisRepo)
	authUC := usecase.NewAuthUseCase(s.cfg, authRepository, authRedisRepo)
	sessUC := sessionUC.NewSessionUseCase(sessionRepo, s.cfg)

	//Init handlers
	walletHandlers := walletHandler.NewWalletHandler(s.cfg, &walletUC, transactionUC)
	authHandlers := authHandler.NewAuthHandlers(s.cfg, authUC, sessUC)
	txHandlers := transactionHandlers.NewTxHandler(s.cfg, authUC, walletUC, transactionUC, sessUC)

	// middlewares
	mw := middleware.NewMiddlewareManager(sessUC, authUC, s.cfg, []string{"*"})
	e.Use(requestid.New())
	e.Use(cors.Default())

	walletHandler.MapWalletRoutes(wallet, walletHandlers, mw)
	authHandler.MapAuthRoutes(auth, authHandlers, mw)
	transactionHandlers.MapTxRoutes(transactions, txHandlers, mw)

	// healthcheck
	status.GET("", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	_ = wallet
	return nil
}

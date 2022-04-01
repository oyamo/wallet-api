package wallets

import (
	"github.com/gin-gonic/gin"
)

type Handler interface {
	Create() gin.HandlerFunc
	Update() gin.HandlerFunc
	Delete() gin.HandlerFunc
	GetByID() gin.HandlerFunc
	GetBalanceByID() gin.HandlerFunc
	Debit() gin.HandlerFunc
	Credit() gin.HandlerFunc
}

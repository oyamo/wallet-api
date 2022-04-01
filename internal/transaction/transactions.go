package transaction

import (
	"github.com/gin-gonic/gin"
)

type Handler interface {
	GetByWalletID() gin.HandlerFunc
}

package http

import (
	"github.com/gin-gonic/gin"
	"github.com/oyamo/wallet-api/internal/auth"
	"github.com/oyamo/wallet-api/internal/middleware"
)

// Map auth routes
func MapAuthRoutes(authGroup *gin.RouterGroup, h auth.Handlers, mw *middleware.MiddlewareManager) {
	authGroup.POST("/register", h.Register())
	authGroup.POST("/login", h.Login())
	authGroup.POST("/logout", mw.AuthSessionMiddleware, mw.CSRF, h.Logout())
	authGroup.GET("/find", mw.AuthSessionMiddleware, mw.CSRF, h.FindByName())
	authGroup.GET("/all", mw.AuthSessionMiddleware, mw.CSRF, h.GetUsers())
	authGroup.GET("/:user_id", mw.AuthSessionMiddleware, mw.CSRF, h.GetUserByID())
	// authGroup.Use(middleware.AuthJWTMiddleware(authUC, cfg))
	//authGroup.Use(mw.AuthSessionMiddleware())
	authGroup.GET("/me", mw.AuthSessionMiddleware, mw.CSRF, h.GetMe())
	authGroup.GET("/token", mw.AuthSessionMiddleware, h.GetCSRFToken())
}

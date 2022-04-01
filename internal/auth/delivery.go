package auth

import "github.com/gin-gonic/gin"

// Auth HTTP Handlers interface
type Handlers interface {
	Register() gin.HandlerFunc
	Login() gin.HandlerFunc
	Logout() gin.HandlerFunc
	Update() gin.HandlerFunc
	Delete() gin.HandlerFunc
	GetUserByID() gin.HandlerFunc
	FindByName() gin.HandlerFunc
	GetUsers() gin.HandlerFunc
	GetMe() gin.HandlerFunc
	GetCSRFToken() gin.HandlerFunc
}

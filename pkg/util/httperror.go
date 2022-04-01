package util

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type HTTPError struct {
	error
	code    int
	message string
}

type httpError interface {
	WriteJSONToCtx(c *gin.Context)
}

func (e HTTPError) WriteJSONToCtx(c *gin.Context) {
	c.AbortWithStatusJSON(e.code, gin.H{"status_code": e.code, "data": nil, "message": e.Error()})
}

func NewError(code int, message string) *HTTPError {
	return &HTTPError{error: errors.New(message), code: code, message: message}
}

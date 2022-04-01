package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/oyamo/wallet-api/pkg/util"
	log "github.com/sirupsen/logrus"
	"time"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// RequestLoggerMiddleware Request logger middleware
func (mw *MiddlewareManager) RequestLoggerMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: ctx.Writer}
		ctx.Writer = blw
		ctx.Next()
		start := time.Now()
		req := ctx.Request
		res := blw.body.String()
		status := ctx.Writer.Status()
		size := len(res)
		s := time.Since(start).String()
		requestID := util.GetRequestID(ctx)

		log.Infof("RequestID: %s, Method: %s, URI: %s, Status: %v, Size: %v, Time: %s",
			requestID, req.Method, req.URL, status, size, s,
		)
	}
}

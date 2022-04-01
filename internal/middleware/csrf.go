package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/oyamo/wallet-api/pkg/csrf"
	"github.com/oyamo/wallet-api/pkg/util"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// CSRF Middleware
func (mw *MiddlewareManager) CSRF(ctx *gin.Context) {
	if !mw.cfg.Server.CSRF {
		ctx.Next()
		return
	}

	token := ctx.Request.Header.Get(csrf.CSRFHeader)
	if token == "" {
		log.Errorf("CSRF Middleware get CSRF header, Token: %s, Error: %s, RequestId: %s",
			token,
			"empty CSRF token",
			util.GetRequestID(ctx),
		)
		util.NewError(http.StatusForbidden, "empty CSRF").WriteJSONToCtx(ctx)
		return
	}

	sid, ok := ctx.Get("sid")
	if !csrf.ValidateToken(token, sid.(string)) || !ok {
		log.Errorf("CSRF Middleware csrf.ValidateToken Token: %s, Error: %s, RequestId: %s",
			token,
			"empty token",
			util.GetRequestID(ctx),
		)
		util.NewError(http.StatusForbidden, "empty CSRF").WriteJSONToCtx(ctx)
		return
	}

	ctx.Next()
}

package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/oyamo/wallet-api/pkg/util"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// Auth sessions middleware using redis
func (mw *MiddlewareManager) AuthSessionMiddleware(c *gin.Context) {
	if c.IsAborted() {
		return
	}

	cookie, err := c.Cookie(mw.cfg.Session.Name)
	if err != nil {

		log.Errorf("AuthSessionMiddleware RequestID: %s, Error: %s",
			util.GetRequestID(c),
			err.Error(),
		)

		if err == http.ErrNoCookie {
			util.NewError(http.StatusUnauthorized, "unauthorised client").WriteJSONToCtx(c)
			c.Abort()
			return
		}

		util.NewError(http.StatusUnauthorized, "unauthorized").WriteJSONToCtx(c)
		c.Abort()
		return
	}

	sid := cookie

	sess, err := mw.sessUC.GetSessionByID(c.Request.Context(), cookie)
	if err != nil {
		log.Errorf("GetSessionByID RequestID: %s, CookieValue: %s, Error: %s",
			util.GetRequestID(c),
			cookie,
			err.Error(),
		)
		util.NewError(http.StatusUnauthorized, "unauthorized").WriteJSONToCtx(c)
		c.Abort()

		return
	}

	user, err := mw.authUC.GetByID(c.Request.Context(), sess.UserID)
	if err != nil {
		log.Errorf("GetByWalletID RequestID: %s, Error: %s",
			util.GetRequestID(c),
			err.Error(),
		)
		util.NewError(http.StatusUnauthorized, "unauthorized").WriteJSONToCtx(c)
		c.Abort()

		return
	}

	c.Set("sid", sid)
	c.Set("uid", sess.SessionID)
	c.Set("user", user)
	c.Next()
}

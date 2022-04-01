package http

import (
	"github.com/gin-gonic/gin"
	"github.com/oyamo/wallet-api/config"
	"github.com/oyamo/wallet-api/internal/auth"
	"github.com/oyamo/wallet-api/internal/models"
	"github.com/oyamo/wallet-api/internal/session"
	"github.com/oyamo/wallet-api/pkg/csrf"
	"github.com/oyamo/wallet-api/pkg/util"
	"net/http"

	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
)

// Auth handlers
type authHandlers struct {
	cfg    *config.Config
	authUC auth.UseCase
	sessUC session.UCSession
}

// NewAuthHandlers Auth handlers constructor
func NewAuthHandlers(cfg *config.Config, authUC auth.UseCase, sessUC session.UCSession) auth.Handlers {
	return &authHandlers{cfg: cfg, authUC: authUC, sessUC: sessUC}
}

// Register godoc
// @Summary Register new user
// @Description register new user, returns user and token
// @Tags Auth
// @Accept json
// @Produce json
// @Success 201 {object} models.User
// @Router /auth/register [post]
func (h *authHandlers) Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		span, ctx := opentracing.StartSpanFromContext(util.GetRequestCtx(c), "auth.Register")
		defer span.Finish()

		user := &models.User{}
		if err := util.ReadRequest(c, user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "data": nil, "status_code": http.StatusBadRequest})
			return
		}

		createdUser, err := h.authUC.Register(ctx, user)
		if err != nil {
			err.(*util.HTTPError).WriteJSONToCtx(c)
			return
		}

		sess, err := h.sessUC.CreateSession(ctx, &models.Session{
			UserID: createdUser.User.UserID,
		}, h.cfg.Session.Expire)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "data": nil, "status_code": http.StatusInternalServerError})
			return
		}

		cookie := util.CreateSessionCookie(h.cfg, sess)
		c.SetCookie(cookie.Name, cookie.Value, cookie.MaxAge, cookie.Path, cookie.Domain, cookie.Secure, cookie.HttpOnly)
		c.JSON(http.StatusOK, gin.H{"message": "user created successfully", "status_code": http.StatusOK, "data": createdUser})
	}
}

// Login godoc
// @Summary Login new user
// @Description login user, returns user and set session
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {object} models.User
// @Router /auth/login [post]
func (h *authHandlers) Login() gin.HandlerFunc {
	type Login struct {
		Email    string `json:"email" db:"email" validate:"omitempty,lte=60,email"`
		Password string `json:"password,omitempty" db:"password" validate:"required,gte=6"`
	}
	return func(c *gin.Context) {
		span, ctx := opentracing.StartSpanFromContext(util.GetRequestCtx(c), "auth.Login")
		defer span.Finish()

		login := &Login{}
		if err := util.ReadRequest(c, login); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "data": nil, "status_code": http.StatusBadRequest})
			return
		}

		userWithToken, err := h.authUC.Login(ctx, &models.User{
			Email:    login.Email,
			Password: login.Password,
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorised", "status_code": http.StatusUnauthorized})
			return
		}

		sess, err := h.sessUC.CreateSession(ctx, &models.Session{
			UserID: userWithToken.User.UserID,
		}, h.cfg.Session.Expire)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "status_code": http.StatusInternalServerError})
			return
		}

		cookie := util.CreateSessionCookie(h.cfg, sess)
		c.SetCookie(cookie.Name, cookie.Value, cookie.MaxAge, cookie.Path, cookie.Domain, cookie.Secure, cookie.HttpOnly)
		c.JSON(http.StatusOK, gin.H{"message": "login success", "status_code": http.StatusOK, "data": userWithToken})
	}
}

// Logout godoc
// @Summary Logout user
// @Description logout user removing session
// @Tags Auth
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"ok"
// @Router /auth/logout [post]
func (h *authHandlers) Logout() gin.HandlerFunc {
	return func(c *gin.Context) {
		span, ctx := opentracing.StartSpanFromContext(util.GetRequestCtx(c), "authHandlers.Logout")
		defer span.Finish()

		cookie, err := c.Cookie("session-id")
		if err != nil {
			if errors.Is(err, http.ErrNoCookie) {
				c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error() + " /logout", "data": nil, "status_code": http.StatusUnauthorized})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "data": nil, "status_code": http.StatusInternalServerError})
			return
		}

		if err := h.sessUC.DeleteByID(ctx, cookie); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "data": nil, "status_code": http.StatusInternalServerError})
			return
		}

		util.DeleteSessionCookie(c, h.cfg.Session.Name)

		c.JSON(http.StatusOK, gin.H{"message": "success", "status_code": http.StatusOK})
	}
}

// Update godoc
// @Summary Update user
// @Description update existing user
// @Tags Auth
// @Accept json
// @Param id path int true "user_id"
// @Produce json
// @Success 200 {object} models.User
// @Router /auth/{id} [put]
func (h *authHandlers) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		span, ctx := opentracing.StartSpanFromContext(util.GetRequestCtx(c), "authHandlers.Update")
		defer span.Finish()

		uID, err := uuid.Parse(c.Param("user_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "data": nil, "status_code": http.StatusBadRequest})
			return
		}

		user := &models.User{}
		user.UserID = uID

		if err = util.ReadRequest(c, user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "data": nil, "status_code": http.StatusBadRequest})
			return
		}

		updatedUser, err := h.authUC.Update(ctx, user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "data": nil, "status_code": http.StatusInternalServerError})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "success", "data": updatedUser, "status_code": http.StatusInternalServerError})

	}
}

// GetUserByID godoc
// @Summary get user by id
// @Description get string by ID
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param id path int true "user_id"
// @Success 200 {object} models.User
// @Failure 500 {object} httpErrors.RestError
// @Router /auth/{id} [get]
func (h *authHandlers) GetUserByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		span, ctx := opentracing.StartSpanFromContext(util.GetRequestCtx(c), "authHandlers.GetUserByID")
		defer span.Finish()

		uID, err := uuid.Parse(c.Param("user_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "data": nil, "status_code": http.StatusBadRequest})
			return
		}

		user, err := h.authUC.GetByID(ctx, uID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": err.Error(), "data": nil, "status_code": http.StatusNotFound})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "user found", "data": user, "status_code": http.StatusOK})

	}
}

// Delete
// @Summary Delete user account
// @Description some description
// @Tags Auth
// @Accept json
// @Param id path int true "user_id"
// @Produce json
// @Success 200 {string} string	"ok"
// @Failure 500 {object} httpErrors.RestError
// @Router /auth/{id} [delete]
func (h *authHandlers) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		span, ctx := opentracing.StartSpanFromContext(util.GetRequestCtx(c), "authHandlers.Delete")
		defer span.Finish()

		uID, err := uuid.Parse(c.Param("user_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "data": nil, "status_code": http.StatusBadRequest})
			return
		}

		if err = h.authUC.Delete(ctx, uID); err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": err.Error(), "data": nil, "status_code": http.StatusNotFound})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "deleted successfully", "data": nil, "status_code": http.StatusOK})
	}
}

// FindByName godoc
// @Summary Find by name
// @Description Find user by name
// @Tags Auth
// @Accept json
// @Param name query string false "username" Format(username)
// @Produce json
// @Success 200 {object} models.UsersList
// @Failure 500 {object} httpErrors.RestError
// @Router /auth/find [get]
func (h *authHandlers) FindByName() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "not implemented", "data": nil, "status_code": http.StatusInternalServerError})
	}
}

// GetUsers godoc
// @Summary Get users
// @Description Get the list of all users
// @Tags Auth
// @Accept json
// @Param page query int false "page number" Format(page)
// @Param size query int false "number of elements per page" Format(size)
// @Param orderBy query int false "filter name" Format(orderBy)
// @Produce json
// @Success 200 {object} models.UsersList
// @Failure 500 {object} httpErrors.RestError
// @Router /auth/all [get]
func (h *authHandlers) GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "not implemented", "data": nil, "status_code": http.StatusInternalServerError})
	}
}

// GetMe godoc
// @Summary Get user by id
// @Description Get current user by id
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {object} models.User
// @Failure 500 {object} httpErrors.RestError
// @Router /auth/me [get]
func (h *authHandlers) GetMe() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorised", "status_code": http.StatusUnauthorized})
	}
}

// GetCSRFToken godoc
// @Summary Get CSRF token
// @Description Get CSRF token, required auth session cookie
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {string} string "Ok"
// @Failure 500 {object} httpErrors.RestError
// @Router /auth/token [get]
func (h *authHandlers) GetCSRFToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		span, _ := opentracing.StartSpanFromContext(util.GetRequestCtx(c), "authHandlers.GetCSRFToken")
		defer span.Finish()

		sid, ok := c.Get("sid")
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"message": "sid not found", "status_code": http.StatusBadRequest})
			return
		}
		token := csrf.MakeToken(sid.(string))
		c.Header(csrf.CSRFHeader, token)
		c.Header("Access-Control-Expose-Headers", csrf.CSRFHeader)

		c.JSON(http.StatusOK, gin.H{"message": "", "status_code": http.StatusOK})
	}
}

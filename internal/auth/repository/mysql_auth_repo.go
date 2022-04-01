package repository

import (
	"context"
	"github.com/oyamo/wallet-api/internal/auth"
	"github.com/oyamo/wallet-api/internal/models"
	"gorm.io/gorm"
	"time"

	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
)

// Auth Repository
type authRepo struct {
	db *gorm.DB
}

// Auth Repository constructor
func NewAuthRepository(db *gorm.DB) auth.Repository {
	return &authRepo{db: db}
}

// Create new user
func (r *authRepo) Register(ctx context.Context, user *models.User) (*models.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "authRepo.Register")
	defer span.Finish()

	user.LoginDate = time.Now()

	err := r.db.Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Update existing user
func (r *authRepo) Update(ctx context.Context, user *models.User) (*models.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "authRepo.Update")
	defer span.Finish()
	err := r.db.Where("user_id = ?", user.UserID.String()).Save(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Delete existing user
func (r *authRepo) Delete(ctx context.Context, userID uuid.UUID) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "authRepo.Delete")
	defer span.Finish()

	err := r.db.Where("user_id = ?", userID.String()).Delete(&models.User{}).Error
	if err != nil {
		return err
	}

	return nil
}

// Get user by id
func (r *authRepo) GetByID(ctx context.Context, userID uuid.UUID) (*models.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "authRepo.GetByWalletID")
	defer span.Finish()

	user := models.User{}
	if err := r.db.Where("user_id = ?", userID.String()).First(&user); err.Error != nil {
		return nil, err.Error
	}
	return &user, nil
}

// Find user by email
func (r *authRepo) FindByEmail(ctx context.Context, user *models.User) (*models.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "authRepo.FindByEmail")
	defer span.Finish()

	var foundUser models.User
	err := r.db.Where("email = ?", user.Email).First(&foundUser).Error
	if err != nil {
		return nil, err
	}

	return &foundUser, nil
}

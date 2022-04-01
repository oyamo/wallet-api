package usecase

import (
	"context"
	"github.com/google/uuid"
	"github.com/oyamo/wallet-api/config"
	"github.com/oyamo/wallet-api/internal/auth"
	"github.com/oyamo/wallet-api/internal/models"
	"reflect"
	"testing"
)

func TestNewAuthUseCase(t *testing.T) {
	type args struct {
		cfg       *config.Config
		authRepo  auth.Repository
		redisRepo auth.RedisRepository
	}
	tests := []struct {
		name string
		args args
		want auth.UseCase
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAuthUseCase(tt.args.cfg, tt.args.authRepo, tt.args.redisRepo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAuthUseCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_authUC_Delete(t *testing.T) {
	type fields struct {
		cfg       *config.Config
		authRepo  auth.Repository
		redisRepo auth.RedisRepository
	}
	type args struct {
		ctx    context.Context
		userID uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &authUC{
				cfg:       tt.fields.cfg,
				authRepo:  tt.fields.authRepo,
				redisRepo: tt.fields.redisRepo,
			}
			if err := u.Delete(tt.args.ctx, tt.args.userID); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_authUC_GenerateUserKey(t *testing.T) {
	type fields struct {
		cfg       *config.Config
		authRepo  auth.Repository
		redisRepo auth.RedisRepository
	}
	type args struct {
		userID string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &authUC{
				cfg:       tt.fields.cfg,
				authRepo:  tt.fields.authRepo,
				redisRepo: tt.fields.redisRepo,
			}
			if got := u.GenerateUserKey(tt.args.userID); got != tt.want {
				t.Errorf("GenerateUserKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_authUC_GetByID(t *testing.T) {
	type fields struct {
		cfg       *config.Config
		authRepo  auth.Repository
		redisRepo auth.RedisRepository
	}
	type args struct {
		ctx    context.Context
		userID uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &authUC{
				cfg:       tt.fields.cfg,
				authRepo:  tt.fields.authRepo,
				redisRepo: tt.fields.redisRepo,
			}
			got, err := u.GetByID(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_authUC_Login(t *testing.T) {
	type fields struct {
		cfg       *config.Config
		authRepo  auth.Repository
		redisRepo auth.RedisRepository
	}
	type args struct {
		ctx  context.Context
		user *models.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.UserWithToken
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &authUC{
				cfg:       tt.fields.cfg,
				authRepo:  tt.fields.authRepo,
				redisRepo: tt.fields.redisRepo,
			}
			got, err := u.Login(tt.args.ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Login() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_authUC_Register(t *testing.T) {
	type fields struct {
		cfg       *config.Config
		authRepo  auth.Repository
		redisRepo auth.RedisRepository
	}
	type args struct {
		ctx  context.Context
		user *models.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.UserWithToken
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &authUC{
				cfg:       tt.fields.cfg,
				authRepo:  tt.fields.authRepo,
				redisRepo: tt.fields.redisRepo,
			}
			got, err := u.Register(tt.args.ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Register() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_authUC_Update(t *testing.T) {
	type fields struct {
		cfg       *config.Config
		authRepo  auth.Repository
		redisRepo auth.RedisRepository
	}
	type args struct {
		ctx  context.Context
		user *models.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &authUC{
				cfg:       tt.fields.cfg,
				authRepo:  tt.fields.authRepo,
				redisRepo: tt.fields.redisRepo,
			}
			got, err := u.Update(tt.args.ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Update() got = %v, want %v", got, tt.want)
			}
		})
	}
}

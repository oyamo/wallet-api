package repository

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/oyamo/wallet-api/internal/models"
	"github.com/oyamo/wallet-api/internal/transaction"
	"reflect"
	"testing"
)

func TestNewRedisRepo(t *testing.T) {
	type args struct {
		db *redis.Client
	}
	tests := []struct {
		name string
		args args
		want transaction.RedisRepository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRedisRepo(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRedisRepo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_redisRepo_DeleteTxCtx(t *testing.T) {
	type fields struct {
		db *redis.Client
	}
	type args struct {
		ctx context.Context
		key string
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
			r := redisRepo{
				db: tt.fields.db,
			}
			if err := r.DeleteTxCtx(tt.args.ctx, tt.args.key); (err != nil) != tt.wantErr {
				t.Errorf("DeleteTxCtx() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_redisRepo_GetByIDCtx(t *testing.T) {
	type fields struct {
		db *redis.Client
	}
	type args struct {
		ctx context.Context
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Transaction
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := redisRepo{
				db: tt.fields.db,
			}
			got, err := r.GetByIDCtx(tt.args.ctx, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByIDCtx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetByIDCtx() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_redisRepo_SetTxCtx(t *testing.T) {
	type fields struct {
		db *redis.Client
	}
	type args struct {
		ctx         context.Context
		key         string
		seconds     int
		transaction *models.Transaction
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
			r := redisRepo{
				db: tt.fields.db,
			}
			if err := r.SetTxCtx(tt.args.ctx, tt.args.key, tt.args.seconds, tt.args.transaction); (err != nil) != tt.wantErr {
				t.Errorf("SetTxCtx() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/oyamo/wallet-api/internal/models"
	"github.com/oyamo/wallet-api/internal/transaction"
	"gorm.io/gorm"
	"reflect"
	"testing"
)

func TestNewTXRepo(t *testing.T) {
	type args struct {
		database *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want transaction.Repository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTXRepo(tt.args.database); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTXRepo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mySQLRepo_Create(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx         context.Context
		transaction *models.Transaction
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
			m := mySQLRepo{
				db: tt.fields.db,
			}
			got, err := m.Create(tt.args.ctx, tt.args.transaction)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mySQLRepo_Delete(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx           context.Context
		transactionID uuid.UUID
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
			m := mySQLRepo{
				db: tt.fields.db,
			}
			if err := m.Delete(tt.args.ctx, tt.args.transactionID); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_mySQLRepo_GetByID(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx           context.Context
		transactionID uuid.UUID
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
			m := mySQLRepo{
				db: tt.fields.db,
			}
			got, err := m.GetByID(tt.args.ctx, tt.args.transactionID)
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

func Test_mySQLRepo_GetByWalletId(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx      context.Context
		walletID uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *[]models.Transaction
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := mySQLRepo{
				db: tt.fields.db,
			}
			got, err := m.GetByWalletId(tt.args.ctx, tt.args.walletID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByWalletId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetByWalletId() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mySQLRepo_Update(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx         context.Context
		transaction *models.Transaction
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
			m := mySQLRepo{
				db: tt.fields.db,
			}
			got, err := m.Update(tt.args.ctx, tt.args.transaction)
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

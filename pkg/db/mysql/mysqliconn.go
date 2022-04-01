package mysql

import (
	"fmt"
	"github.com/oyamo/wallet-api/config"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB(config *config.Config) (*gorm.DB, error) {
	gormDSN := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.MySQL.MySQLUser,
		config.MySQL.MySQLPassword,
		config.MySQL.MySQLHost,
		config.MySQL.MySQLPort,
		config.MySQL.MySQLDbname,
	)

	log.Info(gormDSN)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: gormDSN,
	}), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}

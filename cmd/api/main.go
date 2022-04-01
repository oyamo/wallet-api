package main

import (
	"github.com/oyamo/wallet-api/config"
	server "github.com/oyamo/wallet-api/internal/server"
	"github.com/oyamo/wallet-api/migrations"
	"github.com/oyamo/wallet-api/pkg/db/mysql"
	"github.com/oyamo/wallet-api/pkg/db/redis"
	"github.com/oyamo/wallet-api/pkg/util"
	log "github.com/sirupsen/logrus"
	"os"
)

// @title Wallet API
// @version 1.0
// @description An API for managing transactions.go of players
// @contact.name Oyamo Brian
// @contact.url https://github.com/oyamo
// @contact.email oyamo.xyz@gmail.com
// @BasePath /api/v1
func main() {
	log.Infoln("Starting server")

	// Load config
	configPath := util.GetConfigPath(os.Getenv("CONFIG"))

	cfgFile, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}

	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatalf("ParseConfig: %v", err)
	}

	// Connect to db
	mySQLDB, err := mysql.NewDB(cfg)
	if err != nil {
		log.Fatalf("MySQL init: %s \n", err)
	} else {
		log.Infoln("MYSQL connected")
	}

	// Connect to reddis
	redisClient := redis.NewRedisClient(cfg)
	defer redisClient.Close()
	log.Info("Redis connected")

	// Perform automigrations
	err = migrations.AutoMigrate(mySQLDB)
	if err != nil {
		log.Fatalf("Migration: %v", err)
	}

	httpServer := server.NewServer(cfg, redisClient, mySQLDB)
	if err = httpServer.Run(); err != nil {
		log.Fatalf("Server Init Error: %v", err)
	}
}

package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/oyamo/wallet-api/config"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-redis/redis/v8"
)

const (
	certFile       = "ssl/Server.crt"
	keyFile        = "ssl/Server.pem"
	maxHeaderBytes = 1 << 20
	ctxTimeout     = 5
)

// Server struct
type Server struct {
	engine      *gin.Engine
	cfg         *config.Config
	redisClient *redis.Client
	mysqlClient *gorm.DB
}

// NewServer New Server constructor
func NewServer(cfg *config.Config, redisClient *redis.Client, mysqlClient *gorm.DB) *Server {
	return &Server{engine: gin.Default(), cfg: cfg, redisClient: redisClient, mysqlClient: mysqlClient}
}

func (s *Server) Run() error {
	server := &http.Server{
		Addr:           s.cfg.Server.Port,
		ReadTimeout:    time.Second * s.cfg.Server.ReadTimeout,
		WriteTimeout:   time.Second * s.cfg.Server.WriteTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	go func() {
		log.Infof("Server is listening on PORT: %s", s.cfg.Server.Port)
		if err := s.engine.Run(server.Addr); err != nil {
			log.Fatalf("Error starting Server: %v", err)
		}
	}()

	go func() {
		log.Infof("Starting Debug Server on PORT: %s", s.cfg.Server.PprofPort)
		if err := http.ListenAndServe(s.cfg.Server.PprofPort, http.DefaultServeMux); err != nil {
			log.Errorf("Error PPROF ListenAndServe: %s", err)
		}
	}()

	if err := s.MapHandlers(s.engine); err != nil {
		return err
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	_, shutdown := context.WithTimeout(context.Background(), ctxTimeout*time.Second)
	defer shutdown()

	log.Info("Server Exited Properly")
	return nil
}

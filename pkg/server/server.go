package server

import (
	"errors"
	"net"
	"net/http"

	"go.uber.org/zap"

	"github.com/gigovich/netfort/pkg/config"
	"github.com/gin-gonic/gin"
)

// Create GIN engine router.
func Create() *gin.Engine {
	e := gin.Default()
	return e
}

// Run gin-gonic engine to serve requests.
func Run(log *zap.Logger, cfg config.Config, engine *gin.Engine) (*http.Server, error) {
	srv := &http.Server{
		Addr:    cfg.HTTP.Address,
		Handler: engine,
	}

	log.Debug("create net listener", zap.String("address", cfg.HTTP.Address))
	listener, err := net.Listen("tcp", cfg.HTTP.Address)
	if err != nil {
		log.Error("failed create net listener", zap.Error(err))
		return nil, err
	}

	go func() {
		log.Info("start listen server", zap.String("address", cfg.HTTP.Address))
		if err := srv.Serve(listener); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error("serve requests", zap.Error(err))
		}
	}()

	return srv, nil
}

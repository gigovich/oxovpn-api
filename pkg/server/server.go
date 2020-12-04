package server

import (
	"errors"
	"net"
	"net/http"

	"go.uber.org/zap"

	"github.com/gigovich/oxovpn-api/pkg/config"
	"github.com/gigovich/oxovpn-api/pkg/middleware"
	"github.com/gin-gonic/gin"
)

// Create GIN engine router.
func Create(log *zap.Logger, cfg config.Config) (*gin.Engine, error) {
	authMid, err := middleware.Auth(log, cfg)
	if err != nil {
		return nil, err
	}

	e := gin.Default()
	e.Use(authMid.MiddlewareFunc())
	e.Use(gin.Recovery())

	v1 := e.Group("/api/v1")

	auth := v1.Group("/auth")
	auth.POST("/login", authMid.LoginHandler)
	auth.GET("/refresh_token", authMid.RefreshHandler)
	auth.GET("/logout", authMid.LogoutHandler)

	return e, nil
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

package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gigovich/netfort/pkg/config"
	"github.com/gigovich/netfort/pkg/server"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "API server and static content handler",
	Long: `we serving single endpoint for both handlers
but on the different paths.

API server path prefix '/api/v1/',
all static files served from the root '/'.`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.Load(cfgFile)
		if err != nil {
			os.Exit(1)
		}

		// setup logger
		logCfg := zap.Config{
			Development: cfg.Debug,
		}
		if cfg.Debug {
			logCfg.Encoding = "console"
			logCfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
			logCfg.EncoderConfig = zap.NewDevelopmentEncoderConfig()
		} else {
			logCfg.Encoding = "json"
			logCfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
			logCfg.EncoderConfig = zap.NewProductionEncoderConfig()
		}

		logger, err := logCfg.Build()
		if err != nil {
			fmt.Println("create logger:", err)
			os.Exit(1)
		}

		srv, err := server.Run(logger, cfg, server.Create())
		if err != nil {
			os.Exit(1)
		}

		notify := make(chan os.Signal, 1)
		signal.Notify(notify, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
		<-notify

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

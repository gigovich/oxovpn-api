package config

import (
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

// Config object for application settings
type Config struct {
	// Debug configuration for logger and other tracing stuff.
	Debug bool `yaml:"debug" env:"DEBUG"`

	// JWT settings.
	JWT struct {
		Secret string `yaml:"secret" env:"JWT_SECRET"`
	} `yaml:"jwt"`

	// DB settings.
	DB struct {
		// ConnURL
		ConnURL string `yaml:"conn_url" env:"DB_CONN_URL"`
	} `yaml:"db"`

	// // JWT token settings.
	// JWT struct {
	// 	// PublicKey to check signed token.
	// 	PublicKey string `yaml:"public_key" env:"JWT_PUBLIC_KEY"`

	// 	// PrivateKey for sign token.
	// 	PrivateKey string `yaml:"private_key" env:"JWT_PRIVATE_KEY"`
	// } `yaml:"jwt"`

	// HTTP endpoint settings
	HTTP struct {
		// Address for serving HTTP endpoint
		Address string `yaml:"address" env:"HTTP_ADDRESS"`

		// // Timeouts for servering connections in seconds
		// Timeouts struct {
		// 	Read  uint `yaml:"read" env:"HTTP_TIMEOUTS_READ"`
		// 	Write uint `yaml:"write" env:"HTTP_TIMEOUTS_WRITE"`
		// 	Idle  uint `yaml:"idle" env:"HTTP_TIMEOUTS_IDLE"`
		// } `yaml:"timeouts"`

		// // TLS settings for HTTPS server
		// TLS struct {
		// 	CertFile string `yaml:"cert_file" env:"HTTP_TLS_CERT_FILE"`
		// 	KeyFile  string `yaml:"key_file" env:"HTTP_TLS_KEY_FILE"`
		// } `yaml:"tls"`
	} `yaml:"http"`

	// Log settings
	Log struct {
		// Level of logging: fatal,info,warning,error,debug
		Level string `yaml:"level" env:"LOG_LEVEL"`
	} `yaml:"log"`
}

// Default settings config
func Default() Config {
	cfg := Config{}
	cfg.HTTP.Address = ":8080"
	return cfg
}

// Load config object with default values and reload settings from the file.
// Environment variables will change settings at the end of load process.
func Load(cfgPath string) (cfg Config, err error) {
	cfg = Default()
	if err = cleanenv.ReadConfig(cfgPath, &cfg); os.IsNotExist(err) {
		return cfg, nil
	}
	return cfg, err
}

package util

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Environment string `mapstructure:"ENVIRONMENT"`

	// JWT/PASETO Key Configuration
	PublicKeyBase64  string `mapstructure:"PUBLIC_KEY_BASE64"`
	PrivateKeyBase64 string `mapstructure:"PRIVATE_KEY_BASE64"`

	// Token Duration Configuration
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`

	// Email Configuration
	EmailSenderName     string `mapstructure:"EMAIL_SENDER_NAME"`
	EmailSenderAddress  string `mapstructure:"EMAIL_SENDER_ADDRESS"`
	EmailSenderPassword string `mapstructure:"EMAIL_SENDER_PASSWORD"`

	// Data Source Configuration
	DBDriver       string `mapstructure:"DB_DRIVER"`
	DBSourceCoreDB string `mapstructure:"DB_SOURCE_CORE_DB"`
	DBSourceAuthDB string `mapstructure:"DB_SOURCE_AUTH_DB"`
	MigrationURL   string `mapstructure:"MIGRATION_URL"`

	// Internal Address for Internal Connections
	InternalRedisAddress         string `mapstructure:"INTERNAL_REDIS_ADDRESS"`
	InternalManageServiceAddress string `mapstructure:"INTERNAL_MANAGE_SERVICE_ADDRESS"`

	// Bind Address for External Connections
	HTTPManageServiceAddress      string `mapstructure:"HTTP_MANAGE_SERVICE_ADDRESS"`
	GRPCManageServiceAddress      string `mapstructure:"GRPC_MANAGE_SERVICE_ADDRESS"`
	HTTPAuthServiceAddress        string `mapstructure:"HTTP_AUTH_SERVICE_ADDRESS"`
	GRPCAuthServiceAddress        string `mapstructure:"GRPC_AUTH_SERVICE_ADDRESS"`
	HTTPTransactionServiceAddress string `mapstructure:"HTTP_TRANSACTION_SERVICE_ADDRESS"`
	GRPCTransactionServiceAddress string `mapstructure:"GRPC_TRANSACTION_SERVICE_ADDRESS"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app.env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

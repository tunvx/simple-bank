package util

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	// 1. Environment Configuration
	Environment string `mapstructure:"ENVIRONMENT"`

	// 2. JWT/PASETO Key Configuration
	PublicKeyBase64  string `mapstructure:"PUBLIC_KEY_BASE64"`
	PrivateKeyBase64 string `mapstructure:"PRIVATE_KEY_BASE64"`

	// 3. Token Duration Configuration
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`

	// 4. Email Configuration
	EmailSenderName     string `mapstructure:"EMAIL_SENDER_NAME"`
	EmailSenderAddress  string `mapstructure:"EMAIL_SENDER_ADDRESS"`
	EmailSenderPassword string `mapstructure:"EMAIL_SENDER_PASSWORD"`

	// 5. Data Source Configuration
	DBDriver        string `mapstructure:"DB_DRIVER"`
	SourceSchemaURL string `mapstructure:"SOURCE_SCHEMA_URL"`

	DBSourceOriginalDB string `mapstructure:"DB_SOURCE_ORIGINAL_DB"`
	DBSourceAuthDB     string `mapstructure:"DB_SOURCE_AUTH_DB"`

	// Added this for dynamically populated shard URLs
	ListDBSourceCoreDB []string `mapstructure:"-"`

	NumCoreDBShard         int    `mapstructure:"NUM_CORE_DB_SHARD"`
	DBSourceCoreDB_Shard_0 string `mapstructure:"DB_SOURCE_CORE_DB_SHARD_0"`
	DBSourceCoreDB_Shard_1 string `mapstructure:"DB_SOURCE_CORE_DB_SHARD_1"`
	DBSourceCoreDB_Shard_2 string `mapstructure:"DB_SOURCE_CORE_DB_SHARD_2"`
	DBSourceCoreDB_Shard_3 string `mapstructure:"DB_SOURCE_CORE_DB_SHARD_3"`
	DBSourceCoreDB_Shard_4 string `mapstructure:"DB_SOURCE_CORE_DB_SHARD_4"`

	// 6. Internal Address for Internal Connections
	InternalRedisAddress         string `mapstructure:"INTERNAL_REDIS_ADDRESS"`
	InternalManageServiceAddress string `mapstructure:"INTERNAL_MANAGE_SERVICE_ADDRESS"`

	// 7. Bind Address for External Connections
	HTTPShardManServiceAddress    string `mapstructure:"HTTP_SHARDMAN_SERVICE_ADDRESS"`
	GRPCShardManServiceAddress    string `mapstructure:"GRPC_SHARDMAN_SERVICE_ADDRESS"`
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
	viper.SetConfigName("config.env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return
	}

	// Add non-empty shard URLs to ListDBSourceCoreDB
	shardURLs := []string{
		config.DBSourceCoreDB_Shard_0,
		config.DBSourceCoreDB_Shard_1,
		config.DBSourceCoreDB_Shard_2,
		config.DBSourceCoreDB_Shard_3,
		config.DBSourceCoreDB_Shard_4,
	}

	// Loop through shard URLs and add the non-empty ones to ListDBSourceCoreDB
	for _, url := range shardURLs {
		if url != "" {
			config.ListDBSourceCoreDB = append(config.ListDBSourceCoreDB, url)
		}
	}
	return
}

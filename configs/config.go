package configs

import (
	"github.com/spf13/viper"
)

func LoadConfig() (*AppConfig, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./configs")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config AppConfig
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

type AppConfig struct {
	Database DatabaseConfig
	Cache    CacheConfig
}

type DatabaseConfig struct {
	MongoURI    string
	MongoDBName string
}

type CacheConfig struct {
	RedisAddr string
}

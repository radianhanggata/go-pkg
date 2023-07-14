package initializer

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
	DBHost     string `mapstructure:"POSTGRES_HOST"`
	DBUserName string `mapstructure:"POSTGRES_USER"`
	DBSecret   string `mapstructure:"POSTGRES_SECRET"`
	DBName     string `mapstructure:"POSTGRES_DB"`
	DBPort     string `mapstructure:"POSTGRES_PORT"`
	DBSSLMode  string `mapstructure:"POSTGRES_SSLMODE"`
	DBTZ       string `mapstructure:"POSTGRES_TZ"`
	ServerPort string `mapstructure:"PORT"`
}

func LoadConfig(path, filename string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName(filename)

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

func ConnectDB(config *Config) (db *gorm.DB, err error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		config.DBHost,
		config.DBUserName,
		config.DBSecret,
		config.DBName,
		config.DBPort,
		config.DBSSLMode,
		config.DBTZ,
	)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	return
}

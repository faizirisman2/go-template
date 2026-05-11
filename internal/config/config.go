package config

import (
	"fmt"

	"github.com/gofiber/fiber/v2/log"
	"github.com/spf13/viper"
)

type Config struct {
	App AppConfig
	DB  DBConfig
}

type DBConfig struct {
	Postgres Postgres
}

type Postgres struct {
	DBName   string
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

type AppConfig struct {
	Name string
	Port string
}

func LoadConfig() *Config {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Error(fmt.Errorf("Fatal error config file: %s \n", err))
		return nil
	}

	return &Config{
		App: AppConfig{
			Name: viper.GetString("app.name"),
			Port: viper.GetString("app.port"),
		},
		DB: DBConfig{
			Postgres: InitDBPostgres(),
		},
	}
}

func InitDBPostgres() Postgres {
	return Postgres{
		DBName:   viper.GetString("db.postgres.name"),
		Host:     viper.GetString("db.postgres.host"),
		Port:     viper.GetString("db.postgres.port"),
		Username: viper.GetString("db.postgres.username"),
		Password: viper.GetString("db.postgres.password"),
		Database: viper.GetString("db.postgres.database"),
	}
}

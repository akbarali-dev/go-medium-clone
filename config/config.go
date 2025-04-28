package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type config struct {
	Postgres postgres
}

type postgres struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func Load(path string) config {
	godotenv.Load(path + "/.env")

	conf := viper.New()
	conf.AutomaticEnv()

	cfg := config{
		Postgres: postgres{
			Host:     conf.GetString("POSTGRES_HOST"),
			Port:     conf.GetString("POSTGRES_PORT"),
			User:     conf.GetString("POSTGRES_USER"),
			Password: conf.GetString("POSTGRES_PASSWORD"),
			Database: conf.GetString("POSTGRES_DATABASE"),
		},
	}
	return cfg
}

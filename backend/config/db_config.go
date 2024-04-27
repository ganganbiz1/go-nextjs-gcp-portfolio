package config

import (
	"os"
	"strconv"

	"gorm.io/gorm/logger"
)

type PrimaryDBConfig struct {
	User     string
	Password string
	Host     string
	Port     int
	Database string
	TimeZone string
	SSLMode  string
	LogLevel logger.LogLevel
}

type ReplicaDBConfig struct {
	User     string
	Password string
	Host     string
	Port     int
	Database string
	TimeZone string
	SSLMode  string
	LogLevel logger.LogLevel
}

func NewPrimaryDBConfig() (*PrimaryDBConfig, error) {
	p, err := strconv.Atoi(os.Getenv("POSTGRES_PORT"))

	if err != nil {
		return nil, err
	}

	logLevel := logger.Silent
	if os.Getenv("APP_ENV") == "local" || os.Getenv("APP_ENV") == "development" {
		logLevel = logger.Info
	}

	return &PrimaryDBConfig{
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     p,
		Database: os.Getenv("POSTGRES_DB"),
		TimeZone: os.Getenv("POSTGRES_TIME_ZONE"),
		SSLMode:  os.Getenv("POSTGRES_SSL_MODE"),
		LogLevel: logLevel,
	}, nil
}

func NewReplicaDBConfig() (*ReplicaDBConfig, error) {
	p, err := strconv.Atoi(os.Getenv("POSTGRES_PORT_REPLICA"))
	if err != nil {
		return nil, err
	}

	logLevel := logger.Silent
	if os.Getenv("APP_ENV") == "local" || os.Getenv("APP_ENV") == "development" {
		logLevel = logger.Info
	}

	return &ReplicaDBConfig{
		User:     os.Getenv("POSTGRES_USER_REPLICA"),
		Password: os.Getenv("POSTGRES_PASSWORD_REPLICA"),
		Host:     os.Getenv("POSTGRES_HOST_REPLICA"),
		Port:     p,
		Database: os.Getenv("POSTGRES_DB_REPLICA"),
		TimeZone: os.Getenv("POSTGRES_TIME_ZONE"),
		SSLMode:  os.Getenv("POSTGRES_SSL_MODE"),
		LogLevel: logLevel,
	}, nil
}

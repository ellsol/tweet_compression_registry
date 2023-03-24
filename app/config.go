package app

import (
	"fmt"
	"os"
)

const (
	EnvPostgresUser = "POSTGRES_USER"
	EnvPostgresPWD  = "POSTGRES_PWD"
	EnvPostgresHOST = "POSTGRES_HOST"
	EnvPostgresDB   = "POSTGRES_DB"
	EnvPostgresPort = "POSTGRES_PORT"
)

type Config struct {
	Info SqlDBInfo
}

func NewConfig() (*Config, error) {
	info, err := getPostgresData()
	if err != nil {
		return nil, err
	}
	return &Config{info}, nil
}

type SqlDBInfo struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
}

func (pi SqlDBInfo) Info() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", pi.Host, pi.Port, pi.User, pi.Password, pi.DBName)
}

func getPostgresData() (info SqlDBInfo, err error) {
	info = SqlDBInfo{}

	info.Host, err = getEnvOrError(EnvPostgresHOST)
	if err != nil {
		return info, err
	}

	info.Port, err = getEnvOrError(EnvPostgresPort)
	if err != nil {
		return info, err
	}

	info.DBName, err = getEnvOrError(EnvPostgresDB)
	if err != nil {
		return info, err
	}

	info.Password, err = getEnvOrError(EnvPostgresPWD)
	if err != nil {
		return info, err
	}

	info.User, err = getEnvOrError(EnvPostgresUser)
	return info, err
}

func getEnvOr(env string, defaultValue string) string {
	result := os.Getenv(env)
	if result == "" {
		return defaultValue
	}
	return result
}

func getEnvOrError(env string) (string, error) {
	result := os.Getenv(env)
	if result == "" {
		return "nil", fmt.Errorf("env var must be provided: missing %v", env)
	}
	return result, nil
}

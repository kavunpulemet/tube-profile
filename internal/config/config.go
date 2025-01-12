package config

import (
	"fmt"
	"os"
)

const (
	serverPortEnv = "SERVER_PORT"
	dbHostEnv     = "DB_HOST"
	dbPortEnv     = "DB_PORT"
	dbUserEnv     = "DB_USER"
	dbNameEnv     = "DB_NAME"
	dbPasswordEnv = "DB_PASSWORD"
	dbSSLModeEnv  = "DB_SSLMODE"
	jwtSecretEnv  = "JWT_SECRET"
)

type Config struct {
	ServerPort         string
	DBConnectionString string
	JWTSecret          []byte
}

func NewConfig() (Config, error) {
	return Config{
		ServerPort: fmt.Sprintf(":%s", os.Getenv(serverPortEnv)),
		DBConnectionString: fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
			os.Getenv(dbHostEnv), os.Getenv(dbPortEnv), os.Getenv(dbUserEnv),
			os.Getenv(dbNameEnv), os.Getenv(dbPasswordEnv), os.Getenv(dbSSLModeEnv)),
		JWTSecret: []byte(os.Getenv(jwtSecretEnv)),
	}, nil
}

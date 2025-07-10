package env

import (
	custom_error "app/internal/common/error"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type EnvConfig struct {
	// amo
	CLIENT_ID     string
	CLIENT_SECRET string
	REDIRECT_URL  string

	// api
	BASE_RELEASE_URL string
	BASE_DEBUG_URL   string
	BASE_DEV_URL     string

	// db
	POSTGRES_URL      string
	POSTGRES_PASSWORD string
	POSTGRES_USER     string
	POSTGRES_PORT     uint
	DATABASE_NAME     string

	// serv
	READ_TIMEOUT  time.Duration
	WRITE_TIMEOUT time.Duration
	IDLE_TIMEOUT  time.Duration
}

func MustLoadEnv() (*EnvConfig, error) {
	const err_path = "daos_core/internal/bootstrap/env/MustLoadEnv"

	err := godotenv.Load("../.env")
	if err != nil {
		custom_error.GetError(err.Error(), err_path)
		return nil, err
	}

	read := getUintParamFromConfig("READ_TIMEOUT")
	write := getUintParamFromConfig("WRITE_TIMEOUT")
	idle := getUintParamFromConfig("IDLE_TIMEOUT")
	db_port := getUintParamFromConfig("POSTGRES_PORT")

	cfg := &EnvConfig{
		// api
		BASE_RELEASE_URL: os.Getenv("BASE_RELEASE_URL"),
		BASE_DEBUG_URL:   os.Getenv("BASE_DEBUG_URL"),

		// db
		POSTGRES_URL:      os.Getenv("POSTGRES_URL"),
		POSTGRES_PASSWORD: os.Getenv("DATEBASE_PASSWORD"),
		POSTGRES_USER:     os.Getenv("POSTGRES_URL"),
		POSTGRES_PORT:     db_port,
		DATABASE_NAME:     os.Getenv("DATABASE_NAME"),
		BASE_DEV_URL:      os.Getenv("BASE_DEV_URL"),

		// serv
		READ_TIMEOUT:  time.Duration(read),
		WRITE_TIMEOUT: time.Duration(write),
		IDLE_TIMEOUT:  time.Duration(idle),
	}

	return cfg, nil
}

func getUintParamFromConfig(name string) uint {
	// cast to 10 - grade, 8 - uint8
	value, err := strconv.Atoi(os.Getenv(name))
	if err != nil {
		log.Fatalf("par %s is not exist or cant cast to integer", name)
	}

	return uint(value)
}

package bootstrap

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Env struct {
	// server
	BASE_RELEASE_URL string
	BASE_DEBUG_URL   string

	// requests
	READ_TIMEOUT  uint8
	WRITE_TIMEOUT uint8
	IDLE_TIMEOUT  uint8

	// database
	POSGRES_URL       string
	DATEBASE_PASSWORD string
}

func NewEnv() *Env {
	return &Env{}
}

func MustLoadEnv() *Env {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("i cant read env")
	}

	read := getUint8ParamFromConfig("READ_TIMEOUT")
	write := getUint8ParamFromConfig("WRITE_TIMEOUT")
	idle := getUint8ParamFromConfig("IDLE_TIMEOUT")

	cfg := &Env{
		BASE_RELEASE_URL: os.Getenv("BASE_RELEASE_URL"),
		BASE_DEBUG_URL:   os.Getenv("BASE_DEBUG_URL"),

		READ_TIMEOUT:  read,
		WRITE_TIMEOUT: write,
		IDLE_TIMEOUT:  idle,

		POSGRES_URL:       os.Getenv("POSGRES_URL"),
		DATEBASE_PASSWORD: os.Getenv("DATEBASE_PASSWORD"),
	}

	return cfg
}

func getUint8ParamFromConfig(name string) uint8 {
	// cast to 10 - grade, 8 - uint8
	value, err := strconv.ParseInt(os.Getenv(name), 10, 8)
	if err != nil {
		log.Fatalf("par %s is not exist or cant cast to integer", name)
	}

	return uint8(value)
}

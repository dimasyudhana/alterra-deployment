package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var JWTSecret string

type Configuration struct {
	Host      string
	Port      int
	Username  string
	Password  string
	Name      string
	JWTSecret string
}

func InitConfiguration() Configuration {
	err := godotenv.Load()
	if err != nil {
		log.Println("Cannot load environment variables")
	}

	port, err := strconv.Atoi(os.Getenv("Port"))
	if err != nil {
		log.Println("Invalid port number")
	}

	return Configuration{
		Host:      os.Getenv("Host"),
		Port:      port,
		Username:  os.Getenv("Username"),
		Password:  os.Getenv("Password"),
		Name:      os.Getenv("Name"),
		JWTSecret: os.Getenv("JWTSecret"),
	}
}

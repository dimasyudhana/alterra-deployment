package config

import (
	"os"
	"strconv"

	"github.com/labstack/gommon/log"
)

var (
	JWTSecret string
)

type AppConfig struct {
	User     string
	Password string
	Host     string
	Port     int
	Name     string
}

func InitConfig() *AppConfig {
	var cnf = readConfig()
	if cnf == nil {
		return nil
	}

	return cnf
}

func readConfig() *AppConfig {
	var result = new(AppConfig)

	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Fatal("Cannot read config variable")
	// 	return nil
	// }

	result.User = os.Getenv("User")
	result.Password = os.Getenv("Password")
	result.Host = os.Getenv("Host")

	// Cek nilai dari environment variable Port
	if portStr := os.Getenv("Port"); portStr != "" {
		cnvPort, err := strconv.Atoi(portStr)
		if err != nil {
			log.Error("Cannot convert database port", err.Error())
			return nil
		}
		result.Port = int(cnvPort)
	}

	result.Name = os.Getenv("Name")
	JWTSecret = os.Getenv("JWTSecret")
	return result
}

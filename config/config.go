package config

import (
	"os"
	"strconv"

	"github.com/labstack/gommon/log"
)

var JWTSecret string

type Configuration struct {
	Host     string
	Port     int
	Username string
	Password string
	Name     string
}

func InitConfiguration() *Configuration {
	var cnf = readConfig()
	if cnf == nil {
		return nil
	}

	return cnf
}

func readConfig() *Configuration {
	var result = new(Configuration)

	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Fatal("Cannot read config variable")
	// 	return nil
	// }

	result.Username = os.Getenv("Username")
	result.Password = os.Getenv("Password")
	result.Host = os.Getenv("Host")
	cnvPort, err := strconv.Atoi(os.Getenv("Port"))
	if err != nil {
		log.Error("Cannot convert database port: ", err.Error())
		return nil
	}
	result.Port = int(cnvPort)
	result.Name = os.Getenv("Name")
	JWTSecret = os.Getenv("JWTSecret")
	return result
}

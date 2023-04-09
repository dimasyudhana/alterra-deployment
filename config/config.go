package config

import (
	"log"
	"os"
	"strconv"
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

	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Cannot read config variable")
	// 	return nil
	// }

	result.Username = os.Getenv("Username")
	result.Password = os.Getenv("Password")
	result.Host = os.Getenv("Host")
	port, err := strconv.Atoi(os.Getenv("Port"))
	if err != nil {
		log.Println("Invalid port number")
		return nil
	}
	result.Port = port
	result.Name = os.Getenv("Name")
	JWTSecret = os.Getenv("JWTSecret")
	return result
}

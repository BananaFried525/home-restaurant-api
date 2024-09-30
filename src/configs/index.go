package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DatabaseAttribute struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
}

type CacheAttribute struct {
	Host string
	Port string
}

type ConfigsAttribute struct {
	Env      string
	Port     string
	Database *DatabaseAttribute
	Cache    *CacheAttribute
}

var Load *ConfigsAttribute

func loadConfigs() *ConfigsAttribute {
	return &ConfigsAttribute{
		Env:  os.Getenv("NODE_ENV"),
		Port: os.Getenv("PORT"),
		Database: &DatabaseAttribute{
			Host:     os.Getenv("MYSQL_HOST"),
			Port:     os.Getenv("MYSQL_PORT"),
			User:     os.Getenv("MYSQL_USER"),
			Password: os.Getenv("MYSQL_PASSWORD"),
			DbName:   os.Getenv("MYSQL_DATABASE"),
		},
		Cache: &CacheAttribute{
			Host: os.Getenv("REDIS_HOST"),
			Port: os.Getenv("REDIS_PORT"),
		},
	}
}

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Load = loadConfigs()
	log.Println("Configs loaded")
}

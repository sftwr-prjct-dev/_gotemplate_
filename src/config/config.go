package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	AppName string `required:"true" envconfig:"APP_NAME"`
	Version string `required:"true" envconfig:"VERSION"`
	Port    string `required:"true" envconfig:"PORT"`
}

func Init() *Config {
	var cfg Config
	err := godotenv.Load("././.env")
	if err != nil {
		log.Println("Error loading .env file, falling back to cli passed env")
	}
	err = envconfig.Process("", &cfg)
	if err != nil {
		log.Fatalln("Error loading environment variables", err)
	}

	return &cfg
}

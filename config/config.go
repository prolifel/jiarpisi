package config

import (
	"log"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerHost string
	ServerPort string
}

const projectDirName = "grpc"

func loadEnv() {
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	cwd, _ := os.Getwd()
	rootPath := projectName.Find([]byte(cwd))

	err := godotenv.Load(string(rootPath) + "/.env")
	if err != nil {
		log.Fatalf("💀 error loading .env file, msg: %s", err.Error())
	}
}

func GetConfig() Config {
	loadEnv()

	return Config{
		ServerHost: os.Getenv("SERVER_HOST"),
		ServerPort: os.Getenv("SERVER_PORT"),
	}
}

package main

import (
	"github.com/joho/godotenv"
)

func InitEnv() error {
	envErr := godotenv.Load()

	return envErr
}
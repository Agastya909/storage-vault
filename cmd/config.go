package main

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AccessKeyID     string
	SecretAccessKey string
	Region          string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	return &Config{
		AccessKeyID:     os.Getenv("AWS_ACCESS_KEY"),
		SecretAccessKey: os.Getenv("AWS_SECRET_KEY"),
		Region:          os.Getenv("AWS_REGION"),
	}, nil
}

package main

import (
	"os"
	"storage/handlers/aws"

	"github.com/joho/godotenv"
)

type Config struct {
	AccessKeyID     string
	SecretAccessKey string
	Region          string

	Port string
}

type Handlers struct {
	AwsHandler aws.AwsHandler
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
		Port:            os.Getenv("PORT"),
	}, nil
}

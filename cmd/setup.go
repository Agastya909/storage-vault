package main

import (
	"fmt"
	AwsHandler "storage/handlers/aws"
	AwsRepository "storage/repository/aws"
	AwsService "storage/service/aws"
)

func SetupHandlers(cfg *Config) (Handlers, error) {
	awsRepo, err := AwsRepository.NewAwsRepository(cfg.AccessKeyID, cfg.SecretAccessKey, cfg.Region)
	if err != nil {
		fmt.Println("Error creating AWS repository:", err)
		return Handlers{}, fmt.Errorf("failed to create AWS repository: %w", err)
	}

	awsService := AwsService.NewAwsService(awsRepo)

	awsHandler := AwsHandler.NewAwsHandler(awsService)

	return Handlers{
		AwsHandler: *awsHandler,
	}, nil
}

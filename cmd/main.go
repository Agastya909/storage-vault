package main

import (
	"fmt"
	"storage/repository/aws"
)

func main() {
	cfg, err := LoadConfig()
	if err != nil {
		fmt.Println("Error loading configuration:", err)
		return
	}
	awsRepo, err := aws.NewAwsRepository(cfg.AccessKeyID, cfg.SecretAccessKey, cfg.Region)
	if err != nil {
		fmt.Println("Error creating AWS repository:", err)
		return
	}
	bucketName := "vault-storage-01-909"
	result, err := awsRepo.DeleteBucket(bucketName)
	if err != nil {
		fmt.Println("Error deleting bucket:", err)
		return
	}
	fmt.Println(result)
}

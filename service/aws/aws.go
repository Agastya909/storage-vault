package aws

import "fmt"

type AwsService struct {
	repo AwsRepository
}

type AwsRepository interface {
	CreateBucket(bucketName string) (string, error)
	GetBucketInfo(bucketName string) (map[string]any, error)
	CreateSignedPutUrl(bucketName, objectKey string) (string, error)
	GetObjectInfo(bucketName, objectKey string) (map[string]any, error)
}

func NewAwsService(repo AwsRepository) *AwsService {
	return &AwsService{
		repo: repo,
	}
}

func (s *AwsService) CreateBucket(bucketName string) (string, error) {
	return s.repo.CreateBucket(bucketName)
}

func (s *AwsService) GetBucketInfo(bucketName string) (map[string]any, error) {
	return s.repo.GetBucketInfo(bucketName)
}

func (s *AwsService) GetPresignedUrl(bucketName, objectKey string) (string, error) {
	objectInfo, err := s.repo.GetObjectInfo(bucketName, objectKey)
	if err != nil {
		return "", err
	}
	if objectInfo != nil {
		return "", fmt.Errorf("object already exists, use a different key")
	}
	return s.repo.CreateSignedPutUrl(bucketName, objectKey)
}

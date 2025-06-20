package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type AwsRepository struct {
	AccessKeyID     string
	SecretAccessKey string
	Region          string
	AwsCfg          aws.Config
}

func NewAwsRepository(accessKeyID, secretAccessKey, region string) (*AwsRepository, error) {
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion(region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyID, secretAccessKey, "")),
	)
	if err != nil {
		return nil, err
	}
	return &AwsRepository{
		AccessKeyID:     accessKeyID,
		SecretAccessKey: secretAccessKey,
		Region:          region,
		AwsCfg:          cfg,
	}, nil
}

func (r *AwsRepository) CreateBucket(bucketName string) (string, error) {
	s3Client := s3.NewFromConfig(r.AwsCfg)
	input := &s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
		CreateBucketConfiguration: &types.CreateBucketConfiguration{
			LocationConstraint: types.BucketLocationConstraintApSouth1,
		},
	}

	_, err := s3Client.CreateBucket(context.TODO(), input)
	if err != nil {
		return "", fmt.Errorf("failed to create bucket: %w", err)
	}
	return fmt.Sprintf("Bucket created successfully: %s", aws.ToString(&bucketName)), nil
}

func (r *AwsRepository) DeleteBucket(bucketName string) (string, error) {
	s3Client := s3.NewFromConfig(r.AwsCfg)
	_, err := s3Client.DeleteBucket(context.TODO(), &s3.DeleteBucketInput{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		return "", fmt.Errorf("failed to delete bucket: %w", err)
	}
	return fmt.Sprintf("Bucket deleted successfully: %s", bucketName), nil
}

func (r *AwsRepository) CreateSignedPutUrl(bucketName, objectKey string) (string, error) {
	s3Client := s3.NewFromConfig(r.AwsCfg)
	presignClient := s3.NewPresignClient(s3Client)

	input := &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	}

	presignedRequest, err := presignClient.PresignPutObject(context.TODO(), input)
	if err != nil {
		return "", fmt.Errorf("failed to create presigned URL: %w", err)
	}

	return presignedRequest.URL, nil
}

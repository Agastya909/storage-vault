package aws

type AwsService struct {
	repo AwsRepository
}

type AwsRepository interface {
	CreateBucket(bucketName string) (string, error)
	GetBucketInfo(bucketName string) (map[string]any, error)
	CreateSignedPutUrl(bucketName, objectKey string) (string, error)
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
	return s.repo.CreateSignedPutUrl(bucketName, objectKey)
}

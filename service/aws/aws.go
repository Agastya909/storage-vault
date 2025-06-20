package aws

type AwsService struct {
	repo AwsRepository
}

type AwsRepository interface {
	CreateBucket(bucketName string) (string, error)
}

func NewAwsService(repo AwsRepository) *AwsService {
	return &AwsService{
		repo: repo,
	}
}

func (s *AwsService) CreateBucket(bucketName string) (string, error) {
	return s.repo.CreateBucket(bucketName)
}

package aws

import (
	"net/http"
	"storage/utils"

	"github.com/go-chi/chi/v5"
)

type AwsHandler struct {
	AwsService AwsService
}

type AwsService interface {
	CreateBucket(bucketName string) (string, error)
}

func NewAwsHandler(service AwsService) *AwsHandler {
	return &AwsHandler{
		AwsService: service,
	}
}

func (a *AwsHandler) CreateBucket(w http.ResponseWriter, r *http.Request) {
	bucketName := chi.URLParam(r, "bucketName")
	if bucketName == "" {
		http.Error(w, "bucketName is required", http.StatusBadRequest)
		return
	}

	result, err := a.AwsService.CreateBucket(bucketName)
	if err != nil {
		utils.HttpResponseHandler(w, http.StatusInternalServerError, err.Error(), nil, nil)
		return
	}
	utils.HttpResponseHandler(w, http.StatusOK, "Bucket created successfully", map[string]string{"bucketName": result}, nil)
}

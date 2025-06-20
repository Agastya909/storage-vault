package models

type UploadFileRequest struct {
	FileName string `json:"file_name"`
	Bucket   string `json:"bucket"`
}

type UploadFileResponse struct {
	Message string `json:"message"`
}

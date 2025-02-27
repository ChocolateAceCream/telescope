package response

type ClassifyResponse struct {
	ClassName  string  `json:"class_name"`
	Confidence float64 `json:"confidence"`
}

type GetS3UploadPresignedUrlResponse struct {
	PresignedUrl string `json:"presigned_url"`
	ImageUrl     string `json:"image_url"`
}

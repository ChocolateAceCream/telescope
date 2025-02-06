package request

type S3PresignedUrlRequest struct {
	FileName string `form:"file_name"json:"file_name" binding:"required"`
}

type ClassifyRequest struct {
	ImageURL string `form:"image_url"json:"image_url" binding:"required"`
}

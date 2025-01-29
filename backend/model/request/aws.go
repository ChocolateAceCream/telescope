package request

type S3PresignedUrlRequest struct {
	FileName string `form:"file_name"json:"file_name" binding:"required"`
}
